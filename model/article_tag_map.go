package model

type ArticleTagMap struct {
	ArticleID uint64 `gorm:"primaryKey;type:bigint;not null" json:"article_id"`
	TagID     uint64 `gorm:"primaryKey;type:bigint;not null" json:"tag_id"`
}

func (ArticleTagMap) TableName() string {
	return "article_tag_map"
}
