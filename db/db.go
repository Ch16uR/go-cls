package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	DBHOST  = "192.168.0.103"
	DBTYPE  = "postgres"
	DBNAME  = "js_lo"
	DBLOGIN = "js_lo"
	DBPASS  = "bsgjjzp"
)

var DB gorm.DB

func init() {

	var err error
	DB, err = gorm.Open(DBTYPE, DBTYPE+"://"+DBLOGIN+":"+DBPASS+"@192.168.0.103/"+DBNAME+"?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	DB.DB()

}
