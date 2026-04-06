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
		// fx.Provide(
		// 	fx.Annotate(
		// 		database.NewPrimaryDB,
		// 		fx.ResultTags(`name:"primary"`),
		// 	),
		// ),
		fx.Provide(config.NewConfig),
		fx.Provide(fx.Annotate(database.NewPrimaryDB, fx.ResultTags(`name:"primary"`))),
		fx.Provide(fx.Annotate(database.NewReplicaDB, fx.ResultTags(`name:"replica"`))),
		fx.Provide(fx.Annotate(server.NewServer, fx.ParamTags(``, ``, `name:"primary"`, `name:"replica"`))),
		fx.Invoke(func(s *server.Server) {
		}),
	).Run()
}
