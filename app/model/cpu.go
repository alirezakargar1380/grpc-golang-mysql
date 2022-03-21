package model

import (
	"github.com/alirezakargar1380/grpc_mysql/app/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Cpu struct {
	gorm.Model
	Name  string `gorm:""json:"name"`
	Brand string `json:"brand"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Cpu{})
}

func CreateBook(b *Cpu) *Cpu {
	db.Create(&b)
	return b
}
