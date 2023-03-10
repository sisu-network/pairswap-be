package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/pairswap-be/src/config"
	"github.com/sisu-network/pairswap-be/src/store"
)

type App struct {
	Router *mux.Router
	Db     *store.DBStores

	Config config.AppConfig
}

func NewApp() *App {
	cfg := config.NewDefaultAppConfig()
	db, err := store.NewDBStores(cfg.DB)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	app := &App{Router: r, Db: db, Config: cfg}
	app.registerMiddlewares()
	app.registerHandlers()
	return app
}

func (a *App) Start() {
	log.Infof("Server listening at port %d ...", a.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.Config.Port), a.Router); err != nil {
		panic(err)
	}
}

func (a *App) registerHandlers() {
	supportFormHandler := NewSupportFormHandler(a.Db)
	supportFormRoute := &Route{
		Path:    "/support",
		Method:  http.MethodPost,
		Handler: supportFormHandler.HandleSubmitSupportForm,
	}

	gasCostHandler := NewGasCostHandler(a.Config.SisuServerURL, a.Config.SisuGasCostPath)
	gasCostRoute := &Route{
		Path:    "/getGasFeeInToken",
		Method:  http.MethodGet,
		Handler: gasCostHandler.HandleGetGasCost,
	}

	historyHandler := NewHistoryHandler(a.Db)
	submitHistoryRoute := &Route{
		Path:    "/history",
		Method:  http.MethodPost,
		Handler: historyHandler.HandleSubmitHistory,
	}
	getHistoryRoute := &Route{
		Path:    "/histories",
		Method:  http.MethodGet,
		Handler: historyHandler.HandleGetHistory,
	}

	gatewayHandler := NewGatewayHandler(a.Config.SisuServerURL, "/getPubKeys")
	gatewayRoute := &Route{
		Path:    "/gateway",
		Method:  http.MethodGet,
		Handler: gatewayHandler.GetGatewayAddress,
	}

	a.Router.HandleFunc(supportFormRoute.Path, supportFormRoute.Handler).Methods(supportFormRoute.Method)
	a.Router.HandleFunc(gasCostRoute.Path, gasCostRoute.Handler).Methods(gasCostRoute.Method)
	a.Router.HandleFunc(submitHistoryRoute.Path, submitHistoryRoute.Handler).Methods(submitHistoryRoute.Method)
	a.Router.HandleFunc(getHistoryRoute.Path, getHistoryRoute.Handler).Methods(getHistoryRoute.Method)
	a.Router.HandleFunc(gatewayRoute.Path, gatewayRoute.Handler).Methods(gatewayRoute.Method)
}

func (a *App) registerMiddlewares() {
	a.Router.Use(customCORSHeader())
}

func customCORSHeader() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, req)
		})
	}
}
