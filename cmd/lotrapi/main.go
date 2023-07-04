package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/yosa12978/lotrapi/internal/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	app.Run()
}
