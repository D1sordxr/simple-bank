package models

type Tag struct {
	Base
	Name  string `gorm:"not null;unique"`
	Posts []Post `gorm:"many2many:post_tags;"`
}
