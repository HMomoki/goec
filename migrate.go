package main

import(
	"sample_blog/database"
)

func main() {
	//Migrationの実行
	database.InitDB()
}