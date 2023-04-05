package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model

	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Post struct {
	gorm.Model

	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func dbInit() *gorm.DB {
	dsn := "root:pass1111@tcp(127.0.0.1:3306)/toit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func insertInitialUser(db *gorm.DB) {
	user := User{
		Username: "太郎",
		Age:      20,
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func insertInitialPost(db *gorm.DB) {
	post := Post{
		Title:   "テストtitle",
		Content: "テストcontent",
		Author:  "テストauthor",
	}
	result := db.Create(&post)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func readAllUsers(db *gorm.DB) []User {
	users := []User{}
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return users
}

func readAllPosts(db *gorm.DB) []Post {
	posts := []Post{}
	result := db.Find(&posts)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return posts
}

func main() {

	db := dbInit()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	insertInitialUser(db)
	insertInitialPost(db)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/posts", func(c *gin.Context) {
		posts := readAllPosts(db)
		c.JSON(200, posts)
	})

	router.Run()
}
