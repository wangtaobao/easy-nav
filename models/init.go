package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewGormDB() {
	db, err := gorm.Open(sqlite.Open("index.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动建表
	err = db.AutoMigrate(&IndexReply{})

	DB = db

}
