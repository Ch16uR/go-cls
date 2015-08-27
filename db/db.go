package db

import (
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
	dbConfig := openConfig()
	DB, err = gorm.Open(dbConfig.DbType, dbConfig.DbType+"://"+dbConfig.DbLogin+":"+dbConfig.DbPass+"@"+dbConfig.DbHost+"/"+dbConfig.DbName+"?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Соединение с базой данных установленно!\n")
	}
	DB.DB()

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
