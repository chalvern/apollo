package model

import "time"

// Model 通用的 Model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"" json:"-"`
	UpdatedAt time.Time `gorm:"" json:"-"`
}
