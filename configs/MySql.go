package configs

import (
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDns string = ""
var mysqlTestDns string = ""
var GormDB *gorm.DB

func ConnectMySqlDatabase() (db *gorm.DB, err error) {
	dsn := mysqlDns

	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//_ = MySqlDb // 일단 컴파일이 되도록 설정
	color.Green("Database Migrated... Success!")
	return GormDB, err
}
