package app

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/yosa12978/lotrapi/internal/api"
)

func Run() {
	listener, err := net.Listen("tcp4", os.Getenv("ADDR"))
	if err != nil {
		panic(err)
	}
	go api.Run(listener)

	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	log.Info().Str("cause", (<-out).String()).Msg("server stopped")
}
