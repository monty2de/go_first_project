package main

import (
	"example/ticket/models"
	"net/http"

	"example/ticket/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TokenMiddleware(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing data"})
		c.Abort()
		return
	}

	var user models.User
	Dsn := "root:@tcp(127.0.0.1:3306)/ticket_sys_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	db.Where("Token = ?", token).Find(&user)
	// edit the condition
	if user.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

func main() {

	Dsn := "root:@tcp(127.0.0.1:3306)/ticket_sys_go?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	Db.AutoMigrate(&models.User{}, &models.Ticket{})

	router := gin.Default()
	router.Run("localhost:8080")

	router.POST("/user/login", controllers.Login)

	// Define a group of routes
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(TokenMiddleware)
	{
		// user
		protectedRoutes.GET("/user/index", controllers.UserIndex)
		protectedRoutes.GET("/user/index/withteckts", controllers.GetAll)
		protectedRoutes.GET("/user/:id", controllers.UserShow)
		protectedRoutes.POST("/user/create", controllers.UserCreate)
		protectedRoutes.PATCH("/user/edit/:id", controllers.UserUpdate)
		protectedRoutes.DELETE("/user/delete/:id", controllers.UserDelete)

		//ticket
		protectedRoutes.GET("/ticket/index", controllers.TicketIndex)
		protectedRoutes.GET("/ticket/:id", controllers.TicketShow)
		protectedRoutes.POST("/ticket/create", controllers.TicketCreate)
		protectedRoutes.PATCH("/ticket/edit/:id", controllers.TicketUpdate)
		protectedRoutes.DELETE("/ticket/delete/:id", controllers.TicketDelete)

	}

}
