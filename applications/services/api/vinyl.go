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

type Vinyl struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Artist      string `json:"artist,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Year        int64  `json:"year,omitempty"`
}

type VinylSimple struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Artist string `json:"artist,omitempty"`
	Year   int64  `json:"year,omitempty"`
}

func GetVinyl(c *gin.Context) {
	vinylId := c.Param("id")
	id, err := strconv.ParseInt(vinylId, 10, 64)
	eh.PanicOnError(err, "failed to parse bookId to int64")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewVinylServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book, err := con.GetVinyl(ctx, &wrapperspb.Int64Value{Value: id})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func GetVinylByTitle(c *gin.Context) {
	vinylTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewVinylServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book, err := con.GetVinylByTitle(ctx, &pb.VinylTitle{Title: vinylTitle})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func GetVinylSimpleByTitle(c *gin.Context) {
	vinylTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewVinylServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book, err := con.GetVinylSimpleByTitle(ctx, &pb.VinylTitle{Title: vinylTitle})
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func GetVinylsBySearch(c *gin.Context) {
	vinylTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewVinylServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	vinyls, err := con.GetVinylsBySearch(ctx, &pb.VinylTitle{Title: vinylTitle})
	vinylList := []Vinyl{}
	for _, v := range vinyls.Vinyls {
		vinylList = append(vinylList, Vinyl{Id: v.Id, Name: v.Name, Description: v.Description, Artist: v.Artist, Amount: v.Amount, Year: v.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, vinylList)
	}
}

func GetAllVinyls(c *gin.Context) {
	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewVinylServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	vinyls, err := con.GetAllVinyls(ctx, &emptypb.Empty{})
	vinylList := []Vinyl{}
	for _, v := range vinyls.Vinyls {
		vinylList = append(vinylList, Vinyl{Id: v.Id, Name: v.Name, Description: v.Description, Artist: v.Artist, Amount: v.Amount, Year: v.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, vinylList)
	}
}

func GetVinylRecsArtist(c *gin.Context) {
	vinylTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewVinylServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	vinyls, err := con.GetVinylRecsArtist(ctx, &pb.VinylTitle{Title: vinylTitle})
	vinylList := []VinylSimple{}
	for _, v := range vinyls.Vinyls {
		vinylList = append(vinylList, VinylSimple{Id: v.Id, Name: v.Name, Artist: v.Artist, Year: v.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, vinylList)
	}
}

func GetVinylRecsYear(c *gin.Context) {
	vinylTitle := c.Param("title")

	conn, err := grpc.Dial(configuration.Neo4j.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "failed to connect to grpc")
	defer conn.Close()

	con := pb.NewVinylServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	vinyls, err := con.GetVinylRecsYear(ctx, &pb.VinylTitle{Title: vinylTitle})
	vinylList := []VinylSimple{}
	for _, v := range vinyls.Vinyls {
		vinylList = append(vinylList, VinylSimple{Id: v.Id, Name: v.Name, Artist: v.Artist, Year: v.Year})
	}

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, vinylList)
	}
}
