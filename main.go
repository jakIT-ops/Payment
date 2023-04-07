package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"payment_full/api"
	db "payment_full/db/sqlc"
	"payment_full/gapi"
	"payment_full/pb"
	"payment_full/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	//runGinServer(config, store)
	//
	//redisOpt := asynq.RedisClientOpt{
	//	Addr: config.RedisAddress,
	//}

	//taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	//go runTaskProcessor(config, redisOpt, store)
	//go runGatewayServer(config, store, taskDistributor)
	runGrpcServer(config, store)
	runGinServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

//func runTaskProcessor(config utils.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
//	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
//	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
//	log.Info().Msg("start task processor")
//	err := taskProcessor.Start()
//	if err != nil {
//		log.Fatal().Err(err).Msg("failed to start task processor")
//	}
//}

func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

//func runGatewayServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
//	server, err := gapi.NewServer(config, store, taskDistributor)
//	if err != nil {
//		log.Fatal().Err(err).Msg("cannot create server")
//	}
//
//	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
//		MarshalOptions: protojson.MarshalOptions{
//			UseProtoNames: true,
//		},
//		UnmarshalOptions: protojson.UnmarshalOptions{
//			DiscardUnknown: true,
//		},
//	})
//
//	grpcMux := runtime.NewServeMux(jsonOption)
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
//	if err != nil {
//		log.Fatal().Err(err).Msg("cannot register handler server")
//	}
//
//	mux := http.NewServeMux()
//	mux.Handle("/", grpcMux)
//
//	statikFS, err := fs.New()
//	if err != nil {
//		log.Fatal().Err(err).Msg("cannot create statik fs")
//	}
//
//	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
//	mux.Handle("/swagger/", swaggerHandler)
//
//	listener, err := net.Listen("tcp", config.HTTPServerAddress)
//	if err != nil {
//		log.Fatal().Err(err).Msg("cannot create listener")
//	}
//
//	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
//	handler := gapi.HttpLogger(mux)
//	err = http.Serve(listener, handler)
//	if err != nil {
//		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
//	}
//}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

//
//func main() {
//	// config
//	config, err := utils.LoadConfig(".")
//	if err != nil {
//		log.Fatal("cannot load config:", err)
//	}
//	//// INITIAL DATABASE
//	//db.ConnectDb()
//	//
//	//// ######grpc server###############
//	////lis, err := net.Listen("tcp", ":4040")
//	////if err != nil {
//	////	panic(err)
//	////}
//	////srv := grpc.NewServer()
//	////pb.RegisterTransactionServer(srv, &rpc.PaymentServer{})
//	////reflection.Register(srv)
//	////log.Printf("server 	listening at %v", lis.Addr())
//	////if e := srv.Serve(lis); e != nil {
//	////	panic(err)
//	////}
//
//	file, err := os.OpenFile("./log/123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatalf("error opening file: %v", err)
//	}
//	defer file.Close()
//
//	app := fiber.New()
//
//	//
//
//	// Middlewares
//	app.Use(cors.New())
//	app.Use(favicon.New(favicon.Config{
//		File: "./public/favicon_io/favicon.ico",
//	}))
//	app.Use(logger.New(logger.Config{
//		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
//		Output: file,
//	}))
//
//	// INITIAL ROUTE
//	app.Static("/", "./public")
//	//app.Post("/create", api.CreateUsers)
//	//app.Post("/login", api.LoginUser)
//
//	api.NewServer(config, app)
//
//	app.Listen(":8080")
//}
