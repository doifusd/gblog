package global

// import "gorm.io/gorm"
import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	DBEngine *gorm.DB
	Redis    *redis.Client
)
