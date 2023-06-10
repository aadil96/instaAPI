package router

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/aadil96/instaAPI/Models"
	"github.com/gin-gonic/gin"
)

var users []models.User
  
func InitRoutes() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.GET("/users/:user", getUsersById)
	r.POST("/users", addUser)
	// r.PUT("/users/:user", getUsers)
	r.DELETE("/users/:user", deleteUser)

	r.Run()
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(200, users)
}

func getUsersById(c *gin.Context) {
	var id string = c.Param("user")
	intId, _ := strconv.Atoi(id)
	
	for _, user := range users {
		fmt.Println(user.ID == intId)
		if user.ID == intId {
			c.IndentedJSON(200, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
	
}

func addUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "error": err.Error()})
		return
	}

	for _, u := range users {
		if u.ID == user.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
			return
		}
		
		if (u.Email == user.Email) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
			return
		}

		if (u.Mobile == user.Mobile) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User exists with this mobile number"})
			return
		}
	}

	users = append(users, user)
	c.IndentedJSON(200, users)
}


func deleteUser(c *gin.Context) {
	var id string = c.Param("user")
	intId, _ := strconv.Atoi(id)

	for i, u := range users {
		if u.ID == intId {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(200, users)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
 