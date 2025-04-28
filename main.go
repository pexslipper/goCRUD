package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
)

type Users struct {
	gorm.Model
	Name  string
	Score uint
}

func main() {

	dsn := "root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&Users{})

	// Create sample users
	// users := []*Users{
	// 	{Name: "Jinzhu", Score: 18},
	// 	{Name: "Jackson", Score: 19},
	// }

	// db.Create(users) // pass a slice to insert multiple row

	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		var users []Users
		db.Find(&users)
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		var user Users
		name := c.Param("name")
		result := db.Where("name = ?", name).First(&user)
		if result.RowsAffected == 0 {
			c.JSON(200, gin.H{
				"Message": "No user name " + name,
			})
		} else {
			c.JSON(200, gin.H{
				"users": user,
			})

		}
	})

	router.Run("localhost:3000") // listen and serve on 0.0.0.0:8080
}
