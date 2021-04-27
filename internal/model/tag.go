package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name       string `json:"name"`
	CreatedBy  string `gorm:"column:created_by;default:null" json:"created_by"`
	ModifiedBy string `gorm:"column:modified_by;default:null" json:"modified_by"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	err := db.Model(&t).Where("state=?", t.State).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	if err = db.Select("id,created_on,modified_on,deleted_on,state,name,created_by,modified_by").Where("state=?", t.State).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	//对0值判断
	err := db.Model(t).Where("id=? and is_del=?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?", t.Model.ID, 0).Delete(&t).Error
}

func (t Tag) GetOne(db *gorm.DB) (int64, error) {
	var data *Tag = &Tag{}
	res := db.Where("name=?", t.Name).Where("created_by=?", t.CreatedBy).Find(&data)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
