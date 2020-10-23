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
func GetUsers(c *gin.Context) {
	userDetail := models.UserDetail{}
	userId, err := models.ExtractToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	sql := "SELECT id,name,email FROM users WHERE id = ?"
	rows, err := database.Mysql.Query(sql, userId)
	if err != nil {
		fmt.Println("Error Query Data Users")
		return
	}
	for rows.Next() {
		if err := rows.Scan(&userDetail.ID, &userDetail.Name, &userDetail.Email); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"users": userDetail})
	return
}
