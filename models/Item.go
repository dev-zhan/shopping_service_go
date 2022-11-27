package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Country     string  `json:"country"`
	Description string  `json:"description"`
}
