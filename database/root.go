package database

import (
	"strings"
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
	enabled := config.GetConfig("db.enable").(string)
	if strings.EqualFold(enabled, "Y") {
		driver := config.GetConfig("db.driver").(string)
		url := config.GetConfig("db.url").(string)
		maxopen := config.GetConfigX("db.max-open", 5).(int)
		maxidle := config.GetConfigX("db.max-idle", 5).(int)
		durationval := config.GetConfigX("db.max-timeout", 2000).(string)
		timeout, err := time.ParseDuration(durationval)
		if err != nil {
			logger.Warn("Error Parsing Max-Connection-Timeout Configuration. Using Default 2000ms")
			timeout = 2 * time.Second
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
	logger.Info("Bootstrap DB Connection Pool Disabled SuccessFully")
	return &DB{DB: nil}
}
