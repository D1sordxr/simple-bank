package models

type Author struct {
	Base
	Posts []Post `gorm:"many2many:author_posts"`
	Name  string `gorm:"unique; not null"`
}
