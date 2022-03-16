package models

import "gorm.io/gorm"

type Coffe struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
}
