package database

import (
	"time"

	"github.com/avinash92c/bootstrap-go/foundation"
	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/v2/logger"
)

type DB struct {
	DB *sqlx.DB
}

type DatabaseService interface{}

//GetConnectionPool creates a connection pool and returns accessor to it
func GetConnectionPool(config foundation.ConfigStore) *DB {
	driver := config.GetConfig("db.driver").(string)
	url := config.GetConfig("db.url").(string)
	maxopen := config.GetConfig("db.max-open").(int)
	maxidle := config.GetConfig("db.max-idle").(int)
	durationval := config.GetConfig("db.max-timeout").(string)
	timeout, err := time.ParseDuration(durationval)
	if err != nil {
		logger.Warn("Error Parsing Max-Connection-Timeout Configuration. Using Default 100ms")
		timeout, _ = time.ParseDuration(durationval)
	}
	db, err := sqlx.Connect(driver, url)
	if err != nil {
		logger.Error("Error Connecting to Database")
		panic(err)
	}
	db.SetMaxIdleConns(maxidle)
	db.SetMaxOpenConns(maxopen)
	db.SetConnMaxLifetime(timeout)
	logger.Info("DB Connection Pool Initialized SuccessFully")
	return &DB{DB: db}
}
