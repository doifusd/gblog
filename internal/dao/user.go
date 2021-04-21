package dao

import (
	"blog/internal/model"
)

func (d *Dao) GetUser(mobile string) (*model.User, error) {
	user := model.User{Mobile: mobile}
	return user.GetOne(d.engine)
}

func (d *Dao) CreateUser(name, mobile, passwd string) error {
	user := model.User{
		Name:   name,
		Mobile: mobile,
		Passwd: passwd,
	}
	return user.Create(d.engine)
}
