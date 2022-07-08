package api

import (
	"fmt"
	"payment_full/token"
	"payment_full/utils"

	"github.com/gofiber/fiber/v2"
)

var A = new(AccountServer)

// tokenMaker :=
type Server struct {
	config     utils.Config
	tokenMaker token.Maker
}

func NewServer(config utils.Config, router *fiber.App) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}
	server.InitRouter(router)
	return server, nil
}

func (server *Server) InitRouter(router *fiber.App) {

	//user
	router.Post("/users", server.createUser)
	router.Post("/users/login", server.loginUser)
	// router.Get("/test", authMiddleware(server.tokenMaker))
	authRoutes := router.Group("/", authMiddleware(server.tokenMaker))
	// account
	authRoutes.Post("/account", A.createAccount)
	authRoutes.Get("/account/:id", A.getAccount)
	authRoutes.Get("/accounts", A.listAccounts)
	authRoutes.Put("/account/:id", A.updateAccount)
	authRoutes.Delete("/account/:id", A.deleteAccount)
	// router.Post("/account", A.createAccount)
	// router.Get("/account/:id", A.getAccount)
	// router.Get("/accounts", A.listAccounts)
	// router.Put("/account/:id", A.updateAccount)
	// router.Delete("/account/:id", A.deleteAccount)
}
