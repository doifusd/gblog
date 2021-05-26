package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title      string `json:"title"`
	Descs      string `json:"descs"`
	Content    string `json:"content"`
	Cover      string `json:"cover"`
	CreatedBy  uint32 `gorm:"column:created_by;" json:"created_by"`
	ModifiedBy uint32 `gorm:"column:modified_by;default:0" json:"modified_by"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (t Article) Count(db *gorm.DB) (int64, error) {
	var count int64
	db = db.Where("created_by=?", t.CreatedBy)
	db = db.Where("state=?", t.State)
	err := db.Model(&t).Where("state=?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

//todo 关联tag表
func (t Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var tags []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Title != "" {
		db = db.Where("name=?", t.Title)
	}
	db = db.Where("created_by=?", t.CreatedBy).Where("state=?", t.State)
	if err = db.Where("state=?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Article) Create(db *gorm.DB, tags []uint32, createdBy uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Create(&t).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var buffer strings.Builder
	sql := "insert into `blog_article_tag` (`article_id`,`tag_id`,`created_by`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		tx.Rollback()
		return err
	}
	for i, tagId := range tags {
		if i == len(tags)-1 {
			buffer.WriteString(fmt.Sprintf("('%d','%d',%d);", t.ID, tagId, createdBy))
		} else {
			buffer.WriteString(fmt.Sprintf("('%d','%d',%d),", t.ID, tagId, createdBy))
		}
	}
	err = tx.Exec(buffer.String()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (t Article) Update(db *gorm.DB, values map[string]interface{}, tags []uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	err := tx.Model(t).Updates(values).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//删除
	now := time.Now().Format("2006-01-02 15:04:05")
	old_tags := map[string]interface{}{
		"state":      0,
		"deleted_on": now,
	}
	err = tx.Model(&ArticleTag{}).Where("article_id=?", t.ID).Updates(old_tags).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var buffer strings.Builder
	sql := "insert into `blog_article_tag` (`article_id`,`tag_id`,`created_by`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		tx.Rollback()
		return err
	}

	for i, tagId := range tags {
		if i == len(tags)-1 {
			buffer.WriteString(fmt.Sprintf("('%d','%d',%d);", t.ID, tagId, values["modified_by"]))
		} else {
			buffer.WriteString(fmt.Sprintf("('%d','%d',%d),", t.ID, tagId, values["modified_by"]))
		}
	}
	err = tx.Exec(buffer.String()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (t Article) Delete(db *gorm.DB, values map[string]interface{}) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Model(t).Updates(values).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	old_tags := map[string]interface{}{
		"state":      0,
		"deleted_on": now,
	}
	err = tx.Model(&ArticleTag{}).Where("article_id=?", t.ID).Updates(old_tags).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (t Article) GetOne(db *gorm.DB) (*Article, error) {
	var data *Article = &Article{}
	err := db.Model(t).Find(&data).Error
	// if res.Error != nil {
	// return 0, res.Error
	// }
	// return res.RowsAffected, nil
	return data, err
}
