package models

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
