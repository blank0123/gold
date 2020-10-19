package models

import (
	"github.com/kainhuck/gold/pkg/config"
	"github.com/kainhuck/gold/pkg/log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(config.Collection.Mysql.Driver, config.Collection.Mysql.DSN)
	db.LogMode(true)

	if err != nil {
		log.SugarLogger.Fatalf("数据库连接失败, err: %v", err)
	}

	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
