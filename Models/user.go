package models

import(
	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goec/Repository"
	// "fmt"
)

type User struct{
	// gorm.Model
	Id int `gorm:"primary_key"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(email string,password string) *gorm.DB {
	// passwordEncrypt, _ := crypto.PasswordEncrypt(password)
    db := repository.GormConnect()
    defer db.Close()

	user := User{}
	user.Email = email
	user.Password = password

	// err := db.Create(&user).GetErrors()
    // if err != nil {
    //     return err
    // }
    return db.Create(&user)
}

func GetUser(email string) User {
	db := repository.GormConnect()
	defer db.Close()

	user := User{}
	user.Email = email
	db.Where("email = ?",email).First(&user)
	// fmt.Println(user)

	return user
}