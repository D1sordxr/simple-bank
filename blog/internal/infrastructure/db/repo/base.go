package repo

import "gorm.io/gorm"

type BaseGormRepo struct {
	Session *gorm.DB
}
