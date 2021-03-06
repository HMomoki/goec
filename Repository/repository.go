package repository

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	//DBとの接続
	DBMS     := "mysql"
	USER     := "workuser"
	PASS     := "password"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME   := "SampleBlog"
  
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)
  
	if err != nil {
	  panic(err.Error())
	}

	return db
}