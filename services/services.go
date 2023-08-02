package services

import (
	"fmt"

	"example.com/fxtest/config"
	"go.uber.org/fx"
)

var SomeService *SomeServiceType

type SomeServiceParams struct {
	fx.In
	SomeServiceType
}

type SomeServiceType struct {
	Config *config.Config
}

func (s *SomeServiceType) DoSomething() {
	fmt.Println("I am some service doing something with config: ", s.Config)
}

func NewSomeService(params SomeServiceParams) *SomeServiceType {
	return &SomeServiceType{}
}

var Module = fx.Module("services",
	fx.Provide(NewSomeService),
	fx.Invoke(func(params SomeServiceParams) {
		SomeService = NewSomeService(params)
	}))
