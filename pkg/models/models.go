package models

import (
	"time"
)

type Author struct {
	ID        string     `gorm:"primary_key" json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	CreateAt  time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// Category модель категории статей
type Category struct {
	ID        string     `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	CreateAt  time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// Article модель статьи
type Article struct {
	ID              string     `gorm:"primary_key" json:"id"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	PublicationDate string     `json:"publication_date"`
	AuthorID        string     `json:"-"`
	Author          Author     `json:"author" gorm:"foreignKey:AuthorID;preload:true"`
	Categories      []Category `gorm:"many2many:article_categories;" json:"categories"`
	CreateAt        time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Image           string
	URL             string
	QRCode          string
	DeletedAt       *time.Time `json:"-"`
}
