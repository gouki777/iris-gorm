package model

import "time"

// Order is the database representation of an order
type Movie struct {
	Id         int    `gorm:"column:Id;AUTO_INCREMENT;PRIMARY_KEY;"` //表字段名为：id
	Name       string `gorm:"not null;column:Name;"`                 //表字段名为：name
	Download   string `gorm:"column:Download;"`
	CreateTime time.Time
	Sha256     string `gorm:"not null;column Sha256;"`
}
