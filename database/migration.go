package database

import(
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goec/Repository"
	"goec/Models"
)

func InitDB(){
	db := repository.GormConnect()
	//Migrationの実行
	db.AutoMigrate(&models.User{},&models.Profile{})
	defer db.Close()
}