package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	ID       uint32 `gorm:"primarykey"`
	UserId   uint
	SenderId uint
	Content  string
}
