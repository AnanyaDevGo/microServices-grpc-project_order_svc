package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AnanyaDevGo/microServices-grpc-project_order_svc/pkg/client"
	"github.com/AnanyaDevGo/microServices-grpc-project_order_svc/pkg/config"
	"github.com/AnanyaDevGo/microServices-grpc-project_order_svc/pkg/db"
	"github.com/AnanyaDevGo/microServices-grpc-project_order_svc/pkg/pb"
	service "github.com/AnanyaDevGo/microServices-grpc-project_order_svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }
    h := db.Init(c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

    lis, err := net.Listen("tcp", c.Port)

    if err != nil {
        log.Fatalln("Failed to listing:", err)
    }

    productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

    if err != nil {
        log.Fatalln("Failed to listing:", err)
    }

    fmt.Println("Order Svc on", c.Port)

    s := service.Server{
        H:          h,
        ProductSvc: productSvc,
    }

    grpcServer := grpc.NewServer()

    pb.RegisterOrderServiceServer(grpcServer, &s)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalln("Failed to serve:", err)
    }
}