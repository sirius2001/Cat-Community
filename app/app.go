package main

import "github.com/jinzhu/gorm"

type App struct {
	MoudeleManager *Manager
	DB             *gorm.DB
	
}
