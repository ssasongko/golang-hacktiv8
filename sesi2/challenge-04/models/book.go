package models

type Book struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"not null;type: varchar(50)"`
	Author      string `gorm:"not null;type: varchar(50)"`
	Description string `gorm:"not null;type: varchar(50)"`
}
