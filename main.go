package main

import (
	"go_authentication/database"
	"go_authentication/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//connect and migrate database
	database.Mysql = database.Connect()
	defer database.Mysql.Close()
	// inisialisasi router
	router := gin.Default()
	setupRouter(router)
	router.Run(":3000")
}

// inisialisai route
func setupRouter(r *gin.Engine) {
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("register", routes.AuthRegister)
		authGroup.POST("login", routes.AuthLogin)
		authGroup.POST("users", routes.GetUserData)
	}
}
