package server

import (
	"context"
	"fmt"

	"github.com/Soyaib10/exploring-some-golang-packages/fx/config"
	"github.com/Soyaib10/exploring-some-golang-packages/fx/database"
	"go.uber.org/fx"
)

type Server struct {
    primary *database.Database
    replica *database.Database
    port    string
}

func NewServer(
    lc fx.Lifecycle,
    cfg *config.Config,
    primary *database.Database,
    replica *database.Database,
) *Server {
    s := &Server{primary: primary, replica: replica, port: cfg.Port}

    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            fmt.Println("Server starting on port:", s.port)
            fmt.Println("Primary DB:", s.primary.Host)
            fmt.Println("Replica DB:", s.replica.Host)
            return nil
        },
        OnStop: func(ctx context.Context) error {
            fmt.Println("Server shutting down gracefully")
            return nil
        },
    })

    return s
}