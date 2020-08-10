package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func StartRPC() {

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		fmt.Println("listen error!")
		return
	}

	gRPCServer := grpc.NewServer()
	RegisterIUserRPCServiceServer(gRPCServer, NewUserRPCService())
	grpc.WithInsecure()

	fmt.Println("gRPCServer.UserRPCService...." + lis.Addr().String() )

	gRPCServer.Serve(lis)

}
