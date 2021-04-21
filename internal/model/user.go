package model

import "github.com/jinzhu/gorm"

type User struct {
	*Model
	Name   string `json:"name"`
	Passwd string `json:"password"`
	Mobile string `json:"mobile"`
	State  uint8  `json:"state"`
}

func (t User) TableName() string {
	return "blog_users"
}

func (t User) GetOne(db *gorm.DB) (*User, error) {
	// var data *User = &User{}
	var data *User
	err := db.Select("id", "password").Where("mobile=?", t.Mobile).Where("is_del=?", 0).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t User) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}
