package main

import (
	"go_authentication/database"
	"go_authentication/models"
	"go_authentication/routes"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
		authGroup.POST("users", TokenAuthMiddleware(), routes.GetUserData)
	}
}

// Middleware Authentication
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
//Cek Token valid
func TokenValid(r *http.Request) error {
	token, err := models.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
