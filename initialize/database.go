package initialize

import (
	"show-calendar/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbOnce sync.Once
var dbErr error

func NewDB(config config.Pg) *gorm.DB {
	if db == nil {
		dbOnce.Do(func() {
			logger := NewLogger()
			db, dbErr = gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{})
			if dbErr != nil || db == nil {
				logger.Error("db init error ", dbErr)
			}
		})
	}
	return db
}
