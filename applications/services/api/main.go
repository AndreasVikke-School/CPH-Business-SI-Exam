package main

import (
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

// @title           Book & Venyl Loan Service
// @version         1.0
// @description     API for school project
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
// @schemes   http
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/create/", CreateUser)
			user.GET("/get/:id", GetUser)
			user.GET("/all/", GetAllUsers)
		}

		loan := api.Group("/loan")
		{
			loan.POST("/create/", CreateLoanEntry)
			loan.GET("/get/:id", GetLoan)
			loan.GET("/all/", GetAllLoans)
			loan.GET("/all-by-user/:id", GetAllLoansByUser)
		}

		log := api.Group("/log")
		{
			log.POST("/create/", CreateLogEntry)
			log.GET("/get-by-user/:userId/:logId", GetLogByUser)
			log.GET("/all-by-user/:id", GetAllLogsByUser)
		}

		book := api.Group("/book")
		{
			book.GET("/write-csv-to-db/", WriteCsvToDb)
			book.GET("/get/:title", GetBookByTitle)
			book.GET("/get-simple/:title", GetBookSimpleByTitle)
			book.GET("/search/:title", GetBooksBySearch)
			book.GET("/all/", GetAllBooks)
			book.GET("/get-recs-author/:title", GetBookRecsAuthor)
			book.GET("/get-recs-year/:title", GetBookRecsYear)
			book.GET("/checkout/:title", CheckoutBook)
			book.GET("/return/:title", ReturnBook)
		}

		vinyl := api.Group("/vinyl")
		{
			vinyl.GET("/get/:id", GetVinyl)
			vinyl.GET("/get-by-title/:title", GetVinylByTitle)
			vinyl.GET("/get-simple/:title", GetVinylSimpleByTitle)
			vinyl.GET("/search/:title", GetVinylsBySearch)
			vinyl.GET("/all/", GetAllVinyls)
			vinyl.GET("/get-recs-artist/:title", GetVinylRecsArtist)
			vinyl.GET("/get-recs-year/:title", GetVinylRecsYear)
		}

	}

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("0.0.0.0:8081")
}
