package config

import "fmt"

type Config struct {
    DBHost string
    Port   string
}

func NewConfig() *Config {
    fmt.Println("Config তৈরি হচ্ছে...")
    return &Config{
        DBHost: "localhost",
        Port:   "8080",
    }
}