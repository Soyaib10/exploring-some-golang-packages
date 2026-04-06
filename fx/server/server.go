package server

import (
	"fmt"

	"github.com/Soyaib10/exploring-some-golang-packages/fx/config"
	"github.com/Soyaib10/exploring-some-golang-packages/fx/database"
)

type Server struct {
	db   *database.Database
	port string
}

func NewServer(db *database.Database, cfg *config.Config) *Server {
	fmt.Println("Server তৈরি হচ্ছে, port:", cfg.Port)
	return &Server{db: db, port: cfg.Port}
}

func (s *Server) Start() {
	fmt.Println("Server চালু হলো port:", s.port)
}
