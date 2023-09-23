package controllers

import (
	"example/ticket/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var users []models.User

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	db.Find(&users)
	c.IndentedJSON(http.StatusOK, users)

}

// func UserCreate(c *gin.Context) {
// 	var newUser models.User

// 	if err := c.BindJSON(&newUser); err != nil {
// 		return
// 	}

// 	log.Fatal(newUser)

// 	// insert into database
// 	Dsn := "root:@tcp(127.0.0.1:3306)/ticket_sys_go?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, _ := gorm.Open(mysql.Open(Dsn), &gorm.Config{})

// 	db.Create(&models.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password, UserType: newUser.UserType})

// 	c.IndentedJSON(http.StatusCreated, newUser)
// }

func UserCreate(c *gin.Context) {
	// var newUser models.User
	// if err := c.BindJSON(&newUser); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
	// 	return
	// }
	token, errfortoken := models.GenerateRandomToken()
	if errfortoken != nil {
		return
	}

	Name, _ := c.GetQuery("Name")
	Password, _ := c.GetQuery("Password")
	Email, _ := c.GetQuery("Email")
	UserType, _ := c.GetQuery("UserType")
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	// Create a new record in the database
	result := db.Create(&models.User{Name: Name, Email: Email, Password: Password, UserType: UserType, Token: token})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error3": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, "user created")
}

func UserShow(c *gin.Context) {
	id := c.Param("id")
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	var user models.User

	db.First(&user, id)

	c.IndentedJSON(http.StatusOK, user)
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	Name, _ := c.GetQuery("Name")
	Password, _ := c.GetQuery("Password")
	Email, _ := c.GetQuery("Email")
	UserType, _ := c.GetQuery("UserType")

	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	var user models.User

	db.First(&user, id)

	// if Name != "" {
	// 	user.Name = Name
	// }
	// if Password != "" {
	// 	user.Password = Password
	// }
	// if Email != "" {
	// 	user.Email = Email
	// }
	// if UserType != "" {
	// 	user.UserType = UserType
	// }

	// update the user table
	db.Model(&user).Updates(models.User{Name: Name, Email: Email, Password: Password, UserType: UserType})
	c.IndentedJSON(http.StatusOK, user)
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	var user models.User

	db.Delete(&user, id)

	c.IndentedJSON(http.StatusOK, "user deleted")
}

func GetAll(c *gin.Context) {
	var users []models.User
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	db.Model(&models.User{}).Preload("Tickets").Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
	Password, _ := c.GetQuery("Password")
	Email, _ := c.GetQuery("Email")
	//check database coonection
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	var user models.User

	db.Where("Email = ? AND Password = ?", Email, Password).Find(&user)

	c.IndentedJSON(http.StatusOK, user.Token)
}
