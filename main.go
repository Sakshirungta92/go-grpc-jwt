package main

import (
	"fmt"
	"log"
	"net"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"authservice/api"
	"authservice/pkg/auth"
	"authservice/pkg/models"
	"authservice/pkg/protobuff/service"
	"authservice/pkg/utils"
)

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
	log.Fatalf("failed to listen: %v", err)
	}

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	utils.Host, utils.Port, utils.User, utils.Dbname, utils.Password)
	
	// create a server instance
	server := api.Server{}
	server.Auth = &auth.JWT{}

	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %s", err)
	}

	Load(server.DB)
	fmt.Println("Successfully connected!")

	// create a gRPC server object
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			auth.UnaryServerInterceptor(),
		),
	)

	// attach the Auth service to the server
	service.RegisterAuthServer(grpcServer, &server)
	
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}