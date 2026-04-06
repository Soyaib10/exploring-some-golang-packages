package handler

import (
	"fmt"

	"github.com/Soyaib10/exploring-some-golang-packages/fx/database"
)

type Handler struct {
	db *database.Database
}

func NewHandler(db *database.Database) *Handler {
	fmt.Println("Handler created")
	return &Handler{db: db}
}
