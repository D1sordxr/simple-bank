package models

import "github.com/google/uuid"

type Post struct {
	Base
	AuthorID uuid.UUID `gorm:"not null"`
	Author   Author    `gorm:"foreignKey:AuthorID"`
	Title    string    `gorm:"not null"`
	Content  string
	Upvote   int
	Tags     []Tag `gorm:"many2many:post_tags;"`
}

type NotValidPost struct { // Previous struct, will be destroyed asap
	Base
	Author  Author `gorm:"many2many:author_posts"`
	Title   string `gorm:"not null"`
	Content string
	Upvote  int
	Tags    []Tag `gorm:"many2many:post_tags;"`
}
