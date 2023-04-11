package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Post struct {
	gorm.Model

	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Genre   string `json:"genre"`
}

func dbInit() *gorm.DB {
	dsn := "file:./backend/pkg/db/sqlite/db.sqlite3?cache=shared&mode=rwc&_fk=1"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func insertDummyUsers(db *gorm.DB) {
	users := []User{
		{Username: "太郎", Age: 20},
		{Username: "次郎", Age: 21},
		{Username: "三郎", Age: 22},
		{Username: "四郎", Age: 23},
		{Username: "五郎", Age: 24},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
}

func insertDummyPosts(db *gorm.DB) {
	posts := []Post{
		{Title: "テストtitle1", Content: "テストcontent1", Author: "テストauthor1", Genre: "テストgenre1"},
		{Title: "テストtitle2", Content: "テストcontent2", Author: "テストauthor2", Genre: "テストgenre2"},
		{Title: "テストtitle3", Content: "テストcontent3", Author: "テストauthor3", Genre: "テストgenre3"},
		{Title: "テストtitle4", Content: "テストcontent4", Author: "テストauthor4", Genre: "テストgenre4"},
		{Title: "テストtitle5", Content: "テストcontent5", Author: "テストauthor5", Genre: "テストgenre5"},
	}

	for _, post := range posts {
		result := db.Create(&post)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
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

func isTableEmpty(db *gorm.DB, tableName string) bool {
	var count int64
	db.Table(tableName).Count(&count)
	return count == 0
}

func main() {
	db := dbInit()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})

	// Insert dummy users only if the User table is empty
	if isTableEmpty(db, "users") {
		insertDummyUsers(db)
	}

	// Insert dummy posts only if the Post table is empty
	if isTableEmpty(db, "posts") {
		insertDummyPosts(db)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/home", func(c *gin.Context) {
		posts := readAllPosts(db)
		c.JSON(200, posts)
	})

	router.POST("/home", func(c *gin.Context) {
		var post Post
		c.BindJSON(&post)
		result := db.Create(&post)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		posts := readAllPosts(db)
		c.JSON(200, posts)
	})

	router.Run()
}

// package main

// import (
// 	"database/sql"
// 	"log"
// 	"time"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model

// 	Username  string `json:"username"`
// 	Password  string `json:"password"`
// 	Email     string `json:"email"`
// 	Age       int    `json:"age"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// }

// type Post struct {
// 	gorm.Model

// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// 	Author  string `json:"author"`
// 	Genre   string `json:"genre"`
// }

// func dbInit() *gorm.DB {
// 	dsn := "root:freedomfox@tcp(127.0.0.1:3306)/minisns?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	return db
// }

// func createDatabaseIfNotExists() {
// 	db, err := sql.Open("mysql", "root:freedomfox@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("failed to connect to MySQL")
// 	}
// 	defer db.Close()

// 	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS minisns")
// 	if err != nil {
// 		panic("failed to create database minisns")
// 	}
// }

// func insertDummyUsers(db *gorm.DB) {
// 	users := []User{
// 		{Username: "太郎", Age: 20},
// 		{Username: "次郎", Age: 21},
// 		{Username: "三郎", Age: 22},
// 		{Username: "四郎", Age: 23},
// 		{Username: "五郎", Age: 24},
// 	}

// 	for _, user := range users {
// 		result := db.Create(&user)
// 		if result.Error != nil {
// 			log.Fatal(result.Error)
// 		}
// 	}
// }

// func insertDummyPosts(db *gorm.DB) {
// 	posts := []Post{
// 		{Title: "テストtitle1", Content: "テストcontent1", Author: "テストauthor1", Genre: "テストgenre1"},
// 		{Title: "テストtitle2", Content: "テストcontent2", Author: "テストauthor2", Genre: "テストgenre2"},
// 		{Title: "テストtitle3", Content: "テストcontent3", Author: "テストauthor3", Genre: "テストgenre3"},
// 		{Title: "テストtitle4", Content: "テストcontent4", Author: "テストauthor4", Genre: "テストgenre4"},
// 		{Title: "テストtitle5", Content: "テストcontent5", Author: "テストauthor5", Genre: "テストgenre5"},
// 	}

// 	for _, post := range posts {
// 		result := db.Create(&post)
// 		if result.Error != nil {
// 			log.Fatal(result.Error)
// 		}
// 	}
// }

// func readAllUsers(db *gorm.DB) []User {
// 	users := []User{}
// 	result := db.Find(&users)
// 	if result.Error != nil {
// 		log.Fatal(result.Error)
// 	}
// 	return users
// }

// func readAllPosts(db *gorm.DB) []Post {
// 	posts := []Post{}
// 	result := db.Find(&posts)
// 	if result.Error != nil {
// 		log.Fatal(result.Error)
// 	}
// 	return posts
// }

// func isTableEmpty(db *gorm.DB, tableName string) bool {
// 	var count int64
// 	db.Table(tableName).Count(&count)
// 	return count == 0
// }

// func main() {

// 	createDatabaseIfNotExists()
// 	db := dbInit()
// 	db.AutoMigrate(&User{})
// 	db.AutoMigrate(&Post{})

// 	// Insert dummy users only if the User table is empty
// 	if isTableEmpty(db, "users") {
// 		insertDummyUsers(db)
// 	}

// 	// Insert dummy posts only if the Post table is empty
// 	if isTableEmpty(db, "posts") {
// 		insertDummyPosts(db)
// 	}

// 	router := gin.Default()
// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:3000"},
// 		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Type"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	}))

// 	router.GET("/home", func(c *gin.Context) {
// 		posts := readAllPosts(db)
// 		c.JSON(200, posts)
// 	})

// 	router.POST("/home", func(c *gin.Context) {
// 		var post Post
// 		c.BindJSON(&post)
// 		result := db.Create(&post)
// 		if result.Error != nil {
// 			log.Fatal(result.Error)
// 		}
// 		posts := readAllPosts(db)
// 		c.JSON(200, posts)
// 	})

// 	router.Run()
// }
