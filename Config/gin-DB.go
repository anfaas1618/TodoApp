package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//Setup Database

// DB create GORM pointer
var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig BuildConfig create dB config and return a  pointer to struct of DBConfig
func BuildDBConfig() *DBConfig {
	dBConfig := DBConfig{"host.docker.internal", 3306, "root", "todo", ""}
	return &dBConfig
}

// DbURL creates a url for connecting to database
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
