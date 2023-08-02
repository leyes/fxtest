package main

import (
	"fmt"

	"example.com/fxtest/config"
	"example.com/fxtest/services"
	"go.uber.org/fx"
)

type Server struct {
	addr string
}

func NewServer(cfg *config.Config) *Server {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return &Server{addr: addr}
}

func (s *Server) Start() {
	fmt.Printf("Listening on %s\n", s.addr)
	fmt.Printf("Invoking some service: \n")
	services.SomeService.DoSomething()
}

func main() {
	fx.New(
		fx.Provide(config.NewConfig, NewServer),
		services.Module,
		fx.Invoke(func(svr *Server) {
			svr.Start()
		}),
	).Run()
}
