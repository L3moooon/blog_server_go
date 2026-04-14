package admin

import "gorm.io/gorm"

type ListDto struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Title    string `gorm:"type:varchar(45);not null" json:"title"`
	Abstract string `gorm:"type:varchar(255)" json:"abstract"`
	View     int64  `gorm:"type:bigint;not null" json:"view"`
	Like     int64  `gorm:"type:bigint;not null" json:"like"`
	Status   bool   `gorm:"type:tinyint;not null" json:"status"`
	Top      bool   `gorm:"type:tinyint;not null" json:"top"`
	Tags     []Tag  `gorm:"many2many:article_tag_map;association_joint_field:tag_id" json:"tags"`
}
type Tag struct {
	gorm.Model
	Name     string
	Articles []ListDto `gorm:"many2many:article_tag_map;"`
}
