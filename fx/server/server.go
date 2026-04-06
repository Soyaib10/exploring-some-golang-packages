package server

import (
	"context"
	"fmt"

	"github.com/Soyaib10/exploring-some-golang-packages/fx/config"
	"github.com/Soyaib10/exploring-some-golang-packages/fx/database"
	"go.uber.org/fx"
)

type Server struct {
	db   *database.Database
	port string
}

func NewServer(lc fx.Lifecycle, db *database.Database, cfg *config.Config) *Server {
	s := &Server{db: db, port: cfg.Port}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("server starting on port: ", s.port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("server shutting down gracefully")
			return nil
		},
	})

	return s
}

func (s *Server) Start() {
	fmt.Println("Server চালু হলো port:", s.port)
}
