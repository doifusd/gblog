package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Descs         string `json:"descs"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (t Article) Count(db *gorm.DB) (int64, error) {
	var count int64
	db = db.Where("created_by=?", t.CreatedBy)
	db = db.Where("state=?", t.State)
	err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func (t Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var tags []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Title != "" {
		db = db.Where("name=?", t.Title)
	}
	db = db.Where("create_id=?", t.CreatedBy).Where("state=?", t.State)
	if err = db.Where("is_del=?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Article) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Article) Update(db *gorm.DB, values interface{}) error {
	//对0值判断
	err := db.Model(t).Where("id=? and is_del=?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Article) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?", t.Model.ID, 0).Delete(&t).Error
}

func (t Article) GetOne(db *gorm.DB) (int64, error) {
	var data *Article = &Article{}
	res := db.Where("title=?", t.Title).Where("created_by=?", t.CreatedBy).Find(&data)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
