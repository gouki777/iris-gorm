package model

import "time"

// Movie db struct
type Movie struct {
	Id         int    `gorm:"column:Id;AUTO_INCREMENT;PRIMARY_KEY;"` //table struct：id
	Name       string `gorm:"not null;column:Name;"`                 //table struct：name
	Download   string `gorm:"column:Download;"`
	CreateTime time.Time
	Sha256     string `gorm:"not null;column Sha256;"`
}
