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
	router.POST("/api/create_log_entry/", CreateLogEntry)
	router.POST("/api/create_loan_entry/", CreateLoanEntry)

	router.POST("/api/create_user/", CreateUser)
	router.GET("/api/get_user/:id", GetUser)
	router.GET("/api/get_users/", GetAllUsers)

	router.POST("/api/create_loan/", CreateLoan)
	router.GET("/api/get_loan/:id", GetLoan)
	router.GET("/api/get_loans/", GetAllLoans)
	router.GET("/api/get_loans_by_user/:id", GetAllLoansByUser)

	router.POST("/api/create_log/", CreateLog)
	router.GET("/api/get_log_by_user/:userId/:logId", GetLogByUser)
	router.GET("/api/get_logs_by_user/:id", GetAllLogsByUser)

	router.GET("/api/get_book/:id", GetBook)
	router.GET("/api/get_book_by_title/:title", GetBookByTitle)
	router.GET("/api/get_book_simple_by_title/:title", GetBookSimpleByTitle)
	router.GET("/api/get_book_by_search/:title", GetBooksBySearch)
	router.GET("/api/get_books/", GetAllBooks)
	router.GET("/api/get_book_recs_author/:title", GetBookRecsAuthor)
	// skal dette ikke være year fremfor title?
	router.GET("/api/get_book_recs_year/:title", GetBookRecsYear)

	router.GET("/api/get_vinyl/:id", GetVinyl)
	router.GET("/api/get_vinyl_by_title/:title", GetVinylByTitle)
	router.GET("/api/get_vinyl_simple_by_title/:title", GetVinylSimpleByTitle)
	router.GET("/api/get_vinyl_by_search/:title", GetVinylsBySearch)
	router.GET("/api/get_vinyls/", GetAllVinyls)
	router.GET("/api/get_vinyl_recs_author/:title", GetVinylRecsArtist)
	// skal dette ikke være year fremfor title?
	router.GET("/api/get_vinyl_recs_year/:title", GetVinylRecsYear)

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("0.0.0.0:8081")
}
