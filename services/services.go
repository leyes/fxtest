package services

import (
	"fmt"

	"example.com/fxtest/config"
	"example.com/fxtest/dependencies"
	"go.uber.org/fx"
)

type SomeServiceType struct {
	fx.In
	Config     *config.Config
	MaybeValue *string `optional:"true"`
	Deps       *dependencies.DependenciesType
}

var SomeService SomeServiceType

func (s *SomeServiceType) DoSomething() {
	maybeValue := "NOT ASSIGNED"
	if s.MaybeValue != nil {
		maybeValue = *s.MaybeValue
	}
	fmt.Printf("I am some service (with MaybeValue %s) doing something with config: %v and dependencies %v\n", maybeValue, s.Config, s.Deps)
}

func NewValue() *string {
	s := "SomeValue"
	return &s
}

// Testing conditionals in provide
var provideNewValue = false

var ServicesModule = fx.Module("services",
	func() fx.Option {
		if provideNewValue {
			return fx.Options(
				fx.Provide(NewValue),
			)
		}
		return fx.Options()
	}(),
	dependencies.DependenciesModule,
	fx.Populate(&SomeService),
)
