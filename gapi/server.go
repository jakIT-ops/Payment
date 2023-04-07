package gapi

import (
	"fmt"
	db "payment_full/db/sqlc"
	"payment_full/pb"
	"payment_full/token"
	"payment_full/utils"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	//taskDistributor worker.TaskDistributor
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("токен үүсгэлт амжилтгүй: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
