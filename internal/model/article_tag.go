package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagID      uint32 `json:"tag_id"`
	ArticleID  uint32 `json:"article_id"`
	CreatedBy  uint32 `gorm:"column:created_by;" json:"created_by"`
	ModifiedBy uint32 `gorm:"column:modified_by;default:0" json:"modified_by"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (t ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}
