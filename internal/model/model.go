package model

import (
	"blog/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"

	// "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host, databaseSetting.Port, databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)
	//db, err := gorm.Open(databaseSetting.DBType, s,
	//	databaseSetting.UserName,
	//	databaseSetting.Password,
	//	databaseSetting.Host,
	//	databaseSetting.DBName,
	//	databaseSetting.Charset,
	//	databaseSetting.ParseTime,
	//)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseSetting.TablePrefix,
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	// if global.ServerSetting.RunMode == "debug" {
	// db.LogMode(true)
	// }
	// db.SingularTable(true)

	//todo ?
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	// sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}
