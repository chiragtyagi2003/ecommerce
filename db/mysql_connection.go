package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var MysqlConnection *gorm.DB

func ConnectMysqlDb() {

	err := godotenv.Load()

	if err != nil {
		print("Error loading .env file")
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDbName := os.Getenv("MYSQL_NAME")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlDbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Error connection")
	}

	MysqlConnection = db

}
