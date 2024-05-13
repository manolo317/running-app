package main

import (
	"github.com/labstack/echo"
	"running-app/src/domain"
	"running-app/src/infrastructure"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
	dsn = "root:test@tcp(127.0.0.1:3306)/running_app?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	dbinit()
	infrastructure.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}

func dbinit() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.Migrator().CreateTable(domain.User{})
}
