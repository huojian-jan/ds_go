package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	gorm.Model
	Title        string `gorm:"type:varchar(50);not null"`
	Content      string `gorm:"type:text;not null"`
	CategoryName string `gorm:"type:varchar(50);not null"`
	HeadImg      string `gorm:"type varchar(255)"`
	UserId       uint   `gorm:"not null"`
	CategoryId   uint   `gorm:"not null"`
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.NewV4()
	return scope.SetColumn("ID", uid)
}
