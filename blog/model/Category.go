package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(50);not null;unique"`
	Instruction string `json:"instruction" gorm:"type:text"`
}
