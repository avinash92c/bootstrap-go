package database

import (
	"regexp"
	"strings"
	"time"

	"github.com/avinash92c/bootstrap-go/foundation"
	"github.com/avinash92c/bootstrap-go/security"
	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/v2/logger"
)

type DB struct {
	DB *sqlx.DB
}

type DatabaseService interface{}

//GetConnectionPool creates a connection pool and returns accessor to it
func GetConnectionPool(config foundation.ConfigStore) *DB {
	enabled := config.GetConfig("boostrapdb.enable").(string)
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

		//DECRYPT AND REGEN DBURL
		secret := config.GetConfig("boostrapdb.secret").(string)
		url = processdburl(url, secret)

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

// ShutdownPool Shuts down Connection Pool
func ShutdownPool(db *DB) error {
	return db.DB.Close()
}

func processdburl(dburl, secret string) string {
	if strings.Contains(dburl, "ENC(") {
		regex := regexp.MustCompile(`ENC\((.*?)\)`)
		matches := regex.FindStringSubmatch(dburl)
		decstr := security.DecryptAES(matches[1], secret)
		return regex.ReplaceAllString(dburl, decstr)
	}
	return dburl
}
