package model

import (
	"time"
)

type TodoModel struct {
	ID        int64  `gorm:"primary_key;auto_increment;not_null"`
	Title     string `gorm:"type:varchar(255);unique;not null"`
	Notes     string `gorm:"type:varchar(255);not null"`
	Done      bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
}

func (m TodoModel) BeforeCreate() error {

	m.CreatedAt = time.Now()

	return nil
}
