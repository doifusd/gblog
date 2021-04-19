package dao

// import "gorm.io/gorm"
import "github.com/jinzhu/gorm"

//Dao 数据库类
type Dao struct {
	engine *gorm.DB
}

//New 初始化
func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
