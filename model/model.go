package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Shorten struct {
	ID       uint64 `json: "id" gorm:"primaryKey"`
	Redirect string `json: "redirect" gorm:"not null"`
	Shorten  string `json: "shorten" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json: "random"`
}

func Setup() {
	dns := "host=172.168.1.2 user=admin password=superSecret dbname=shortening post=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Shorten{})

	if err != nil {
		fmt.Println(err)
	}
}
