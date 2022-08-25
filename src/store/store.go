package store

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/sisu-network/lib/log"
	gormsql "gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type DBStores struct {
	TokenStore       *TokenStore
	SupportFormStore *SupportFormStore
	HistoryStore     *HistoryStore
	db               *gorm.DB
	Config           DbConfig
}

func NewDBStores(cfg DbConfig) (*DBStores, error) {
	store := &DBStores{

		Config: cfg,
	}

	db, err := store.connectORM(cfg)
	if err != nil {
		return nil, err
	}

	store.TokenStore = NewTokenStore(db)
	store.SupportFormStore = NewSupportFormStore(db)
	store.HistoryStore = NewHistoryStore(db)
	store.db = db

	return store, nil
}

func (d *DBStores) connect() *sql.DB {
	host := d.Config.Host
	username := d.Config.User
	password := d.Config.Password
	schema := d.Config.Schema

	log.Info("Schema = ", schema)

	var err error
	var database *sql.DB
	// TODO: This is a temporary fix to run local docker. The correct fix is to redo the entire
	// life cycle of Sisu and dheart.
	for i := 0; i < 5; i++ {
		// Connect to the db with retry
		log.Verbose("Attempt number ", i+1)
		database, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, d.Config.Port))
		if err == nil {
			break
		}
		time.Sleep(time.Second * 3)
	}

	if err != nil {
		log.Error("All DB connection retry failed")
		panic(err)
	}

	_, err = database.Exec("CREATE DATABASE IF NOT EXISTS " + schema)
	if err != nil {
		panic(err)
	}
	database.Close()

	database, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, d.Config.Port, schema))
	if err != nil {
		panic(err)
	}

	log.Info("Db is connected successfully")

	return database
}

func (d *DBStores) connectORM(config DbConfig) (*gorm.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Schema)

	db, err := gorm.Open(gormsql.Open(source))
	if err != nil {
		log.Error("cannot connect to database: ", config.Schema)
		return nil, err
	}

	dbNative, err := db.DB()
	if err != nil {
		log.Error("cannot obtain sql database object: ", config.Schema)
		return nil, err
	}

	dbNative.Exec("CREATE DATABASE IF NOT EXISTS " + config.Schema)

	return db, nil
}
