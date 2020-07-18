package controller

import "github.com/jinzhu/gorm"

// Article defining the schema
type Article struct {
	gorm.Model
	Ide     int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
	Author  string `json:"Author"`
}

// Articles {} Structure reflecting the DB Schema
var Articles = []Article{
	{Ide: 1, Title: "First Article", Desc: "Description 1", Content: "Content of Article 1", Author: "Ada Lovelace"},
}
var db *gorm.DB
var err error
