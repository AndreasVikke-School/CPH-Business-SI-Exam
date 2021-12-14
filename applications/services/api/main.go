package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	_ "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/docs" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	configuration Configuration
)

// @BasePath /api/

// Hello World
// @Schemes
// @Description Says Hello World
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /hello_world [get]
func HelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

// @title Book & Venyl Loan Service
// @version 1.0
// @description API for school project
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/
// @schemes http
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/hello_world/", HelloWorld)
	router.POST("/api/create_log/", CreateLog)
	router.POST("/api/create_loan_entry/", CreateLoanEntry)

	router.POST("/api/create_user/", CreateUser)
	router.GET("/api/get_user/:id", GetUser)
	router.GET("/api/get_users/", GetAllUsers)

	router.POST("/api/create_loan/", CreateLoan)
	router.GET("/api/get_loan/:id", GetLoan)
	router.GET("/api/get_loans/", GetAllLoans)
	router.GET("/api/get_loans_by_user/:id", GetAllLoansByUser)

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("0.0.0.0:8081")
}
