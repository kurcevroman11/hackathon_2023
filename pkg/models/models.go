package models

import (
	"time"
)

type ArticleData struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	CreateAt string `json:"createAt"`
	Content  string `json:"content"`
	Image    []byte `json:"image"`
}

type Author struct {
	ID        string     `gorm:"primary_key" json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	CreateAt  time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// Article модель статьи
type Article struct {
	ID              string    `gorm:"primary_key" json:"id"`
	Title           string    `json:"title"`
	Subtitle        string    `json:"subtitle"`
	Content         string    `json:"content"`
	PublicationDate string    `json:"publication_date"`
	AuthorID        string    `json:"-"`
	Author          Author    `json:"author" gorm:"foreignKey:AuthorID;preload:true"`
	CreateAt        time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Image           []byte
	QRCode          string
	DeletedAt       *time.Time `json:"-"`
}
