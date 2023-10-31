package dao

import (
	"fmt"
	"log/slog"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sirius2001/Cat-Community/conf"
)

type DB struct {
	db *gorm.DB
}

func NewDao(dial string) *DB {
	conf := conf.GetConfig()
	db, err := gorm.Open(conf.DBConfig.DB, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.DBConfig.Host, conf.DBConfig.Port, conf.DBConfig.User, conf.DBConfig.DB, conf.DBConfig.Password))
	if err != nil {
		slog.Error("dial database erro:", err)
		return nil
	}
	return &DB{
		db: db,
	}
}


