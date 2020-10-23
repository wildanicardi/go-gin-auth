package main

import (
	"go_authentication/database"
	"go_authentication/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//connect and migrate database
	database.Mysql = database.Connect()
	defer database.Mysql.Close()
	// inisialisasi router
	router := gin.Default()
	router.GET("/api/data", welcome)
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("register", routes.AuthRegister)
		authGroup.POST("login", routes.AuthLogin)
		authGroup.POST("users", routes.GetUsers)
	}
	router.Run(":3000")
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Welcome"})
	return
}
