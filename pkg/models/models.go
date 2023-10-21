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
	Public    bool       `json:"public"`
	CreateAt  time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type File struct {
	Id        string
	Name      string
	Path      string
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
	Public          bool      `json:"public"`
	CreateAt        time.Time `json:"created_at"`
	FileId          string    `json:"-"`
	ImgFile         File      `json:"imgFile" gorm:"foreignKey:FileId;preload:true"`
	UpdatedAt       time.Time `json:"updated_at"`
	QRCode          string
	DeletedAt       *time.Time `json:"-"`
}

type FilterArticle struct {
	Public bool
}
