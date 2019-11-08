package models

import (
	"github.com/alexcesaro/log/stdlog"
	"github.com/jinzhu/gorm"
)

var logger = stdlog.GetFromFlags()

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) BeforeSave() {
	logger.Info("BeforeSave")
}

func (p *Product) BeforeCreate() {
	logger.Info("BeforeCreate")
}

func (p *Product) AfterCreate() {
	logger.Info("AfterCreate")
}

func (p *Product) AfterSave() {
	logger.Info("AfterSave")
}
