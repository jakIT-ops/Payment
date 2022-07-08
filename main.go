package main

import (
	"log"
	"os"
	"payment_full/api"
	"payment_full/db"
	"payment_full/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// config
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// INITIAL DATABASE
	db.ConnectDb()

	// ######grpc server###############
	// lis, err := net.Listen("tcp", ":4040")
	// if err != nil {
	// 	panic(err)
	// }
	// srv := grpc.NewServer()
	// pb.RegisterTransactionServer(srv, &rpc.PaymentServer{})
	// reflection.Register(srv)
	// log.Printf("server 	listening at %v", lis.Addr())
	// if e := srv.Serve(lis); e != nil {
	// 	panic(err)
	// }

	file, err := os.OpenFile("./log/123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app := fiber.New()

	// Middlewares
	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon_io/favicon.ico",
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output: file,
	}))

	// INITIAL ROUTE
	app.Static("/", "./public")
	api.NewServer(config, app)

	app.Listen(":8080")
}
