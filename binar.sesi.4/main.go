package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

type Person struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Age      int    `json:"age"`
}

type LoginSuccess struct {
	Person Person `json:"person"`
	Token  string `json:"token"`
}

type Errorku struct {
	Code    string `json:"error_code"`
	Message string `json:"message"`
}

var db *gorm.DB

func main() {
	var err error
	//for heroku deploy
	host := os.Getenv("DATABASE_URL")
	if host != "" {
		db, err = gorm.Open("postgres", host)
	} else {
		//local database
		db, err = gorm.Open("postgres", "host=localhost user=prima dbname=binar sslmode=disable password=root")
	}
	db.AutoMigrate(Person{})
	defer db.Close()
	if err != nil {
		log.Fatal("database error")
	}

	r := gin.Default()
	r.GET("/person", GetPerson)
	r.POST("/person", SetPerson)
	r.POST("/login", Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func GetPerson(c *gin.Context) {
	error := Errorku{}
	token := c.Request.Header.Get("token")
	if CheckToken(token) {
		defer log.Print("defer")
		log.Print("bukan defer")
		persons := []Person{}
		db.Find(&persons)
		log.Print(persons)
		c.JSON(200, persons)
	} else {
		error.Code = "12343"
		error.Message = "maaf salah token"
		c.JSON(400, error)
	}
}

func CheckToken(token string) bool {
	if token == "hgashjgd878yad9w" {
		return true
	}
	return false
}

func Login(c *gin.Context) {
	person := Person{}
	error := Errorku{}
	c.BindJSON(&person)
	err := db.Where("name = ? AND password = ?", person.Name, person.Password).Find(&person).Error
	log.Print(person)
	if err == nil {
		s := LoginSuccess{}
		s.Person = person
		s.Token = "hgashjgd878yad9w"
		c.JSON(200, s)
	} else {
		error.Code = "12343"
		error.Message = "login salah"
		c.JSON(400, error)
	}
}

func SetPerson(c *gin.Context) {
	error := Errorku{}
	token := c.Request.Header.Get("token")
	if CheckToken(token) {
		log.Print(c.Request)
		person := Person{}
		c.BindJSON(&person)

		if person.Age < 20 {
			error.Code = "12342"
			error.Message = "maaf umur anda belum cukup"
			c.JSON(400, error)
		} else {
			db.Save(&person)
			log.Print(person)
			c.JSON(200, person)
		}
	} else {
		error.Code = "12343"
		error.Message = "maaf salah token"
		c.JSON(400, error)
	}
}
