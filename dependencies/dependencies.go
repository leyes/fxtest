package dependencies

import (
	"example.com/fxtest/utils"
	"go.uber.org/fx"
)

type Foo *string

type Bar *string

type Baz *string

type DependenciesType struct {
	Foo Foo
	Bar Bar
	Baz Baz
}

type DependenciesParams struct {
	fx.In
	Foo Foo
	Bar Bar
	Baz Baz
}

func NewFoo() Foo {
	s := "foo"
	return &s
}

func NewBar() Bar {
	s := "bar"
	return &s
}

func NewBaz() Baz {
	s := "foo"
	return &s
}

func NewDependencies(params DependenciesParams) *DependenciesType {
	retval := utils.Construct[DependenciesParams, DependenciesType](params)
	return retval
}

var DependenciesModule = fx.Module("dependencies",
	fx.Provide(NewFoo),
	fx.Provide(NewBar),
	fx.Provide(NewBaz),

	fx.Provide(NewDependencies),
)
