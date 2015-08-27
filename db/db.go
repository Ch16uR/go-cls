package db

import (
	. "../models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
)

type DbConfig struct {
	DbType  string
	DbHost  string
	DbName  string
	DbLogin string
	DbPass  string
}

var DB gorm.DB

func init() {

	var err error
	var dbString string

	dbConfig := openConfig()
	if dbConfig.DbType == "mysql" {
		dbString = dbConfig.DbName + ":" + dbConfig.DbPass + "@tcp(" + dbConfig.DbHost + ":3306)/" + dbConfig.DbName + "?charset=utf8&parseTime=True"
	} else if dbConfig.DbType == "postgres" {
		dbString = dbConfig.DbType + "://" + dbConfig.DbLogin + ":" + dbConfig.DbPass + "@" + dbConfig.DbHost + "/" + dbConfig.DbName + "?sslmode=disable"
	} else {
		panic("Не удается определить SQL драйвер " + dbConfig.DbType)
	}
	DB, err = gorm.Open(dbConfig.DbType, dbString)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Соединение с базой данных установленно!\n")
	}
	DB.DB()
	DB.AutoMigrate(&Workplace{}, &Promocode{})

}

func openConfig() DbConfig {

	dbConfig := DbConfig{}
	data, err := ioutil.ReadFile("./db/config.txt")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &dbConfig)
	if err != nil {
		fmt.Println(err)
	}

	return dbConfig
}
