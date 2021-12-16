package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	eh "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/errorhandler"
	pb "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Book struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Author      string `json:"author,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Year        int64  `json:"year,omitempty"`
}

type BookSimple struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Year   int64  `json:"year,omitempty"`
}

func GetBook(c *gin.Context) {
	bookId := c.Param("id")
	id, err := strconv.ParseInt(bookId, 10, 64)
	eh.PanicOnError(err, "failed to parse bookId to int64")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book, err := con.GetBook(ctx, &wrapperspb.Int64Value{Value: id})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func GetBookByTitle(c *gin.Context) {
	bookTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book, err := con.GetBookByTitle(ctx, &pb.BookTitle{Title: bookTitle})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func GetBookSimpleByTitle(c *gin.Context) {
	bookTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book, err := con.GetBookSimpleByTitle(ctx, &pb.BookTitle{Title: bookTitle})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func GetBooksBySearch(c *gin.Context) {
	bookTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	books, err := con.GetBooksBySearch(ctx, &pb.BookTitle{Title: bookTitle})
	booklist := []Book{}
	for _, b := range books.Books {
		booklist = append(booklist, Book{Id: b.Id, Name: b.Name, Description: b.Description, Author: b.Author, Amount: b.Amount, Year: b.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, booklist)
	}
}

func GetAllBooks(c *gin.Context) {
	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	books, err := con.GetAllBooks(ctx, &emptypb.Empty{})
	bookList := []Book{}
	for _, b := range books.Books {
		bookList = append(bookList, Book{Id: b.Id, Name: b.Name, Description: b.Description, Author: b.Author, Amount: b.Amount, Year: b.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, bookList)
	}
}

func GetBookRecsAuthor(c *gin.Context) {
	bookTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	books, err := con.GetBookRecsAuthor(ctx, &pb.BookTitle{Title: bookTitle})
	booklist := []BookSimple{}
	for _, b := range books.Books {
		booklist = append(booklist, BookSimple{Id: b.Id, Name: b.Name, Author: b.Author, Year: b.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, booklist)
	}
}

func GetBookRecsYear(c *gin.Context) {
	bookTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	books, err := con.GetBookRecsYear(ctx, &pb.BookTitle{Title: bookTitle})
	booklist := []BookSimple{}
	for _, b := range books.Books {
		booklist = append(booklist, BookSimple{Id: b.Id, Name: b.Name, Author: b.Author, Year: b.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, booklist)
	}
}
