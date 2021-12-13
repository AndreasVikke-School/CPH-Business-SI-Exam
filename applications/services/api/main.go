package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	configuration Configuration
)

func HelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/hello_world/", HelloWorld)
	router.POST("/api/create_log/", CreateLog)
	router.POST("/api/create_borrower/", CreateBorrower)
	//route getUser -> metodekald som laver grpc client som kalder grpc serveren hos postgres

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	router.Run("0.0.0.0:8081")
}
