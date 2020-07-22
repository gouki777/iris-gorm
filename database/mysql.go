package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mydb() *gorm.DB {
	db, err := gorm.Open("mysql", "mydb:654321@(172.30.50.15:3306)/mydb?charset=utf8&parseTime=True&loc=Local&timeout=20s")
	if err != nil {
		fmt.Printf("[Error] open db error:%v", err)
	}
	//defer db.Close()
	return db
}
