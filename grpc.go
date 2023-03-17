package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"payment_full/db"
	"payment_full/pb"
	"payment_full/rpc"
)

func main() {

	db.ConnectDb()
	// ############## grpc server ###########################
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterTransactionServer(srv, &rpc.PaymentServer{})
	reflection.Register(srv)
	log.Printf("server listening at %v", lis.Addr())
	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}
