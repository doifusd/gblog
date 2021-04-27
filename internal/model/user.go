package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	*Model
	Name   string `json:"name"`
	Passwd string `gorm:"column:password" json:"password"`
	Mobile string `json:"mobile"`
}

func (t User) TableName() string {
	return "blog_users"
}

func (t User) GetOne(db *gorm.DB, mobile string) (*User, error) {
	// var data *LoginRes = &LoginRes{}
	var data *User = &User{}
	err := db.Select("id,password,name").Where("mobile=?", mobile).Where("state=?", 1).First(data).Error
	// err := db.Select("id", "password").First(&t).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t User) CheckName(db *gorm.DB) error {
	err := db.Select("id, name").Where(&t).First(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t User) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t User) LastId(db *gorm.DB) (uint32, error) {
	res := db.Select("id").Last(&t)
	if res.Error != nil {
		return 0, res.Error
	}
	return t.ID, nil
}
