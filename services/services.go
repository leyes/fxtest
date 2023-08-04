package services

import (
	"fmt"

	"example.com/fxtest/config"
	"go.uber.org/fx"
)

var SomeService *SomeServiceType

type SomeServiceConfig struct {
	Config *config.Config
}

type SomeServiceParams struct {
	fx.In
	SomeServiceConfig
}

type SomeServiceType struct {
	SomeServiceConfig
}

func (s *SomeServiceType) DoSomething() {
	fmt.Printf("I am some service doing something with config: %v\n", s.Config)
}

func NewSomeService(params SomeServiceParams) *SomeServiceType {
	return &SomeServiceType{SomeServiceConfig: params.SomeServiceConfig}
}

var ServicesModule = fx.Module("services",
	fx.Provide(NewSomeService),
	fx.Invoke(func() {
		NewSomeService(SomeServiceParams{})
	}),
)
