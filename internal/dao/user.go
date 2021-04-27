package dao

import (
	"blog/internal/model"
)

func (d *Dao) GetUser(mobile string) (*model.User, error) {
	user := model.User{Mobile: mobile, Model: &model.Model{State: 1}}
	return user.GetOne(d.engine, mobile)
}

func (d *Dao) GetUserName(name string) error {
	user := model.User{Name: name}
	return user.CheckName(d.engine)
}

func (d *Dao) CreateUser(name, mobile, passwd string) error {
	user := model.User{
		Name:   name,
		Mobile: mobile,
		Passwd: passwd,
	}
	return user.Create(d.engine)
}

func (d *Dao) LastId() (uint32, error) {
	user := model.User{}
	return user.LastId(d.engine)
}
