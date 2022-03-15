package models

import "gorm.io/gorm"

type Coffe struct {
	gorm.Model
	name  string `json:"name"`
	price int    `json:"price"`
}
