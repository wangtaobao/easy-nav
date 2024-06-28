package models

import "gorm.io/gorm"

type IndexReply struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"column:name;type:varchar(100);"`
	Url  string `json:"url" form:"url" gorm:"column:url;type:varchar(100);"`
}
