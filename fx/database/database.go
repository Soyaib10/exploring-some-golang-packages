package database

import (
	"fmt"

	"github.com/Soyaib10/exploring-some-golang-packages/fx/config"
)

type Database struct {
	Host string
}

func NewPrimaryDB(cfg *config.Config) *Database {
    fmt.Println("Primary DB connecting:", cfg.DBHost)
    return &Database{Host: cfg.DBHost}
}

func NewReplicaDB(cfg *config.Config) *Database {
    fmt.Println("Replica DB connecting:", cfg.DBHost)
    return &Database{Host: "replica-" + cfg.DBHost}
}