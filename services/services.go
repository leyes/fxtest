package services

import (
	"fmt"

	"example.com/fxtest/config"
	"go.uber.org/fx"
)

type SomeServiceType struct {
	fx.In
	Config *config.Config
}

var SomeService SomeServiceType

func (s *SomeServiceType) DoSomething() {
	fmt.Printf("I am some service doing something with config: %v\n", s.Config)
}

var ServicesModule = fx.Module("services",
	fx.Populate(&SomeService),
)
