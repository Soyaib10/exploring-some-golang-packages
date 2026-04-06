package database

import (
	"fmt"

	"github.com/Soyaib10/exploring-some-golang-packages/fx/config"
)

type Database struct {
	Host string
}

func NewDatabase(cfg *config.Config) *Database {
	fmt.Println("Database connect হচ্ছে:", cfg.DBHost)
	return &Database{Host: cfg.DBHost}
}
