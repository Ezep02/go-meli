package models

import "gorm.io/gorm"

type ServiceModel struct {
	gorm.Model
	Title            string  `json:"title" gorm:"size:150; not null"`
	User_id          int     `json:"user_id" gorm:"not null"`
	Description      string  `json:"description" gorm:"default null"`
	Price            float64 `json:"price"`
	Service_Duration int     `json:"service_duration" gorm:"default null"`
}
