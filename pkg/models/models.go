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
	Theme    string `json:"Theme"`
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
	ID              string `gorm:"primary_key" json:"id"`
	Title           string `json:"title"`
	Subtitle        string `json:"subtitle"`
	Content         string `json:"content"`
	PublicationDate string `json:"publication_date"`
	ThemeId         string
	Theme           Theme     `json:"theme" gorm:"foreignKey:ThemeId;preload:true"`
	Public          bool      `json:"public"`
	CreateAt        time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	FileId          string    `json:"-"`
	ImgFile         File      `json:"imgFile" gorm:"foreignKey:FileId;preload:true"`
	QRCode          string
	DeletedAt       *time.Time `json:"-"`
}

type FilterArticle struct {
	Public bool
}

type Theme struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Image string `json:"image"`
}
