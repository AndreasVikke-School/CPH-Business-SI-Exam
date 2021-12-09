package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	eh "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/errorhandler"
	pb "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/rpc"

	"google.golang.org/grpc"
)

var (
	configuration Configuration
	port          = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedLoanServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	eh.PanicOnError(err, "failed to listen")

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	pb.RegisterLoanServiceServer(s, &server{})

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	MigratePostgres(configuration)

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	eh.PanicOnError(err, "failed to serve")
}
