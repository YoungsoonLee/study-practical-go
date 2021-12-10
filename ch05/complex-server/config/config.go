package config

import (
	"io"
	"log"
)

// AppConfig ...
type AppConfig struct {
	Logger *log.Logger
}

// InitConfig ...
func InitConfig(w io.Writer) AppConfig {
	return AppConfig{
		Logger: log.New(
			w, "", log.Ldate|log.Ltime|log.Lshortfile,
		),
	}
}
