package main

import (
	"github.com/Soyaib10/exploring-some-golang-packages/fx/config"
	"github.com/Soyaib10/exploring-some-golang-packages/fx/database"
	"github.com/Soyaib10/exploring-some-golang-packages/fx/server"
	"go.uber.org/fx"
	// "go.uber.org/fx"
)

// manual way
// func main() {
// 	cfg := config.NewConfig()
// 	db := database.NewDatabase(cfg)
// 	srv := server.NewServer(db, cfg)
// 	srv.Start()
// }

// fx way
func main() {
	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(database.NewDatabase),
		fx.Provide(server.NewServer),
		fx.Invoke(func(s *server.Server) {
			s.Start()
		}),
	).Run()
}
