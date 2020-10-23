package routes

import (
	"fmt"
	"go_authentication/database"
	"go_authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLogin(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.IsAuthenticated(database.Mysql)
	if err != nil {
		fmt.Println("Login gagal")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "Error Authetication"})
}

func AuthRegister(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Register(database.Mysql)
	if err != nil {
		fmt.Println("Registrasi gagal")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id_user": user.ID})
}
func GetUserData(c *gin.Context) {
	userDetail := models.UserDetail{}
	data, err := userDetail.GetUser(database.Mysql, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"users": data})
	return
}
