package models

import (
	"fmt"
	"ginapi/pkg/setting"
	"ginapi/pkg/utils"
	"log"

	"github.com/jinzhu/gorm"
	// grom
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model this is Model
type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt utils.JSONTime `json:"created_at"`
	UpdatedAt utils.JSONTime `json:"updated_at"`
}

func init() {
	var err error
	db, err = gorm.Open(setting.TYPE, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.USER,
		setting.PASSWORD,
		setting.HOST,
		setting.DBNAME))
	if err != nil {
		log.Fatalf("Fail to start mysql server: %v", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB this is CloseDB
func CloseDB() {
	defer db.Close()
}
