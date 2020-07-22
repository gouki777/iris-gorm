package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"iris-gorm/model"
)

func Mydb() *gorm.DB {
	//open gorm mysql
	db, err := gorm.Open("mysql", "mydb:654321@(172.30.50.15:3306)/mydb?charset=utf8&parseTime=True&loc=Local&timeout=20s")
	//error open error log
	if err != nil {
		fmt.Printf("[Error] open db error:%v", err)
	}
	//If mysql-table doesn't exist. => Create sql table
	if !db.HasTable(&model.Movie{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Movie{}).Error; err != nil {
			panic(err)
		}
	}
	//defer db.Close()
	return db
}
