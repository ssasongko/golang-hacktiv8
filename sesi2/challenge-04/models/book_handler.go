package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name_Book string `gorm:"not null;type: varchar(255)"`
	Author    string `gorm:"not null;type: varchar(255)"`
}
