package db_model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Question struct {
	ID         int    `gorm:"primary_key"`
	Year       int    `gorm:"not null"`
	Genre      string `gorm:"not null"`
	Question   string `gorm:"not null"`
	Answer     string `gorm:"not null"`
	Commentary string `gorm:"not null"`
}

func GetAll() []Question {
	db, err := gorm.Open("sqlite3", "data/questions.db")
	if err != nil {
		panic("failed to connect database")
	}

	questions := []Question{}
	db.Find(&questions)

	return questions
}
