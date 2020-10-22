package main

import "github.com/gin-gonic/gin"

func Routes() *gin.Engine {
	router := gin.Default()
	return router
}
func main() {
	router := Routes()
	router.Run(":3000")
}
