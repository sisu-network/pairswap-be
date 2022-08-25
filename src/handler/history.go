package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/pairswap-be/src/model"
	"github.com/sisu-network/pairswap-be/src/store"
)

type HistoryRequest struct {
	Address     string    `json:"address,omitempty"`
	Recipient   string    `json:"recipient,omitempty"`
	SrcChain    string    `json:"src_chain,omitempty"`
	DestChain   string    `json:"dest_chain,omitempty"`
	TokenSymbol string    `json:"token_symbol,omitempty"`
	Amount      string    `json:"amount,omitempty"`
	SrcHash     string    `json:"src_hash,omitempty"`
	DestHash    string    `json:"dest_hash,omitempty"`
	SrcLink     string    `json:"src_link,omitempty"`
	DestLink    string    `json:"dest_link,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HistoryHandler struct {
	db *store.DBStores
}

func NewHistoryHandler(db *store.DBStores) *HistoryHandler {
	return &HistoryHandler{db: db}
}

func (h *HistoryHandler) HandleSubmitHistory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		_ = responseError(w, err)
		return
	}

	request := &HistoryRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		fmt.Println(err)
		_ = responseError(w, err)
		return
	}

	history := &model.History{
		Address:     request.Address,
		Recipient:   request.Recipient,
		SrcChain:    request.SrcChain,
		DestChain:   request.DestChain,
		TokenSymbol: request.TokenSymbol,
		Amount:      request.Amount,
		SrcHash:     request.SrcHash,
		DestHash:    request.DestHash,
		SrcLink:     request.SrcLink,
		DestLink:    request.DestLink,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}

	if err := h.db.HistoryStore.Create(context.Background(), history); err != nil {
		_ = responseError(w, err)
		return
	}

	_ = responseSuccess(w, map[string]interface{}{"msg": "create history successfully"})
}

func (h *HistoryHandler) HandleGetHistory(w http.ResponseWriter, r *http.Request) {
	queryStr := r.URL.Query()
	address := queryStr.Get("address")

	if len(address) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("missing address")); err != nil {
			log.Warn(err)
		}
		return
	}
	histories, err := h.db.HistoryStore.GetAllByAddress(context.Background(), address)

	if err != nil {
		_ = responseError(w, err)
		return
	}

	_ = responseSuccess(w, map[string]interface{}{"msg": "get histories successfully", "histories": histories})
}
