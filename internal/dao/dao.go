package dao

import "gorm.io/gorm"

//Dao 数据库类
type Dao struct {
	engine *gorm.DB
}

//New 初始化
func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
