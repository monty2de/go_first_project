package controllers

import (
	"example/ticket/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Dsn = "root:@tcp(127.0.0.1:3306)/ticket_sys_go?charset=utf8mb4&parseTime=True&loc=Local"
var db, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})

func TicketIndex(c *gin.Context) {
	var tickets []models.Ticket
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Find(&tickets)
	c.IndentedJSON(http.StatusOK, tickets)

}

func TicketCreate(c *gin.Context) {

	Title, _ := c.GetQuery("Title")
	Content, _ := c.GetQuery("Content")
	Status, _ := c.GetQuery("Status")
	UserID, _ := c.GetQuery("UserID")

	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	intUserID, _ := strconv.Atoi(UserID)
	// Create a new record in the database
	result := db.Create(&models.Ticket{Title: Title, Content: Content, Status: Status, UserID: intUserID})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error3": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, "ticket created")
}

func TicketShow(c *gin.Context) {
	id := c.Param("id")
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	var ticket models.Ticket

	db.First(&ticket, id)

	c.IndentedJSON(http.StatusOK, ticket)
}

func TicketUpdate(c *gin.Context) {
	id := c.Param("id")
	Title, _ := c.GetQuery("Title")
	Content, _ := c.GetQuery("Content")
	Status, _ := c.GetQuery("Status")
	UserID, _ := c.GetQuery("UserID")

	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	var ticket models.Ticket

	db.First(&ticket, id)
	intUserID, _ := strconv.Atoi(UserID)
	// update the user table
	db.Model(&ticket).Updates(models.Ticket{Title: Title, Content: Content, Status: Status, UserID: intUserID})
	c.IndentedJSON(http.StatusOK, ticket)
}

func TicketDelete(c *gin.Context) {
	id := c.Param("id")
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	var ticket models.Ticket

	db.Delete(&ticket, id)

	c.IndentedJSON(http.StatusOK, "ticket deleted")
}
