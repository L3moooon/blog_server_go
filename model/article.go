package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID         uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string  `gorm:"type:varchar(45);not null" json:"title"`
	CoverImg   string  `gorm:"type:varchar(255)" json:"cover_img"`
	Abstract   string  `gorm:"type:varchar(255)" json:"abstract"`
	Content    string  `gorm:"type:text;not null" json:"content"`
	AuthorName string  `gorm:"type:varchar(45);not null" json:"author_name"`
	View       int64   `gorm:"type:bigint;not null" json:"view"`
	Like       int64   `gorm:"type:bigint;not null" json:"like"`
	Tag        []int64 `gorm:"type:json;not null" json:"tag"`
	Status     bool    `gorm:"type:tinyint;not null" json:"status"`
	Top        bool    `gorm:"type:tinyint;not null" json:"top"`
}

func (Article) TableName() string {
	return "article"
}
