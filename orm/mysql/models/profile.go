package models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Name      string
	User      User `gorm:"association_foreignkey:UserRefer"`
	UserRefer uint
}
