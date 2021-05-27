package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title      string   `json:"title"`
	Descs      string   `json:"descs"`
	Content    string   `json:"content"`
	Cover      string   `json:"cover"`
	CreatedBy  uint32   `gorm:"column:created_by;" json:"created_by"`
	ModifiedBy uint32   `gorm:"column:modified_by;default:0" json:"modified_by"`
	Tags       []string `json:"tags"`
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

func (t Article) Info(db *gorm.DB, articleId uint32) (*Article, error) {
	var article *Article = &Article{}
	if err := db.Model(t).Where("id=?", articleId).First(&article).Error; err != nil {
		return nil, err
	}

	dtb := db.Select("bt.name")
	dtb = dtb.Joins("left join blog_tag bt on blog_article_tag.tag_id=bt.id ")
	dtb = dtb.Where("blog_article_tag.state=?", 1).Where("bt.state=?", 1)
	dtb = dtb.Where("blog_article_tag.article_id=?", article.ID)
	var atag = []*struct {
		Name string
	}{}
	err := dtb.Table("blog_article_tag").Find(&atag).Error
	if err != nil {
		return nil, err
	}
	tags := []string{}
	for _, v := range atag {
		tags = append(tags, v.Name)
	}
	article.Tags = tags
	return article, nil
}

func (t Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var article []*Article
	var err error
	select_str := "id,title,descs,cover,content,state,created_by,modified_by,created_on,modified_on,deleted_on"
	article_query := db.Select(select_str)
	if pageOffset >= 0 && pageSize > 0 {
		article_query = article_query.Offset(pageOffset).Limit(pageSize)
	}
	if t.Title != "" {
		article_query = article_query.Where("title=?", t.Title)
	}
	article_query = article_query.Where("created_by=?", t.CreatedBy).Where("state=?", t.State)
	if err = article_query.Find(&article).Error; err != nil {
		return nil, err
	}

	var aid = []string{}
	// aid := make([]string, len(article))
	for _, val := range article {
		if val.ID > 1 {
			aid = append(aid, fmt.Sprintf("%d", val.ID))
		}
	}
	dtb := db.Select("bt.name,blog_article_tag.article_id")
	dtb = dtb.Joins("left join blog_tag bt on blog_article_tag.tag_id=bt.id ")
	dtb = dtb.Where("blog_article_tag.state=?", 1).Where("bt.state=?", 1)
	dtb = dtb.Where("blog_article_tag.article_id in (?)", aid)
	var atag = []*struct {
		ArticleID uint32
		Name      string
	}{}
	if err = dtb.Table("blog_article_tag").Find(&atag).Error; err != nil {
		return nil, err
	}
	for _, v := range article {
		for _, vt := range atag {
			if v.ID == vt.ArticleID {
				v.Tags = append(v.Tags, vt.Name)
			}
		}
	}
	return article, nil
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
