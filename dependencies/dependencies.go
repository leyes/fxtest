package dependencies

import "go.uber.org/fx"

type Foo string

type Bar string

type Baz string

type DependenciesType struct {
	Foo Foo
	Bar Bar
	Baz Baz
}

type DependenciesParams struct {
	fx.In
	DependenciesType
}

func NewFoo() Foo {
	return "foo"
}

func NewBar() Bar {
	return "bar"
}

func NewBaz() Baz {
	return "baz"
}

var DependenciesModule = fx.Options(
	fx.Provide(NewFoo),
	fx.Provide(NewBar),
	fx.Provide(NewBaz),
)
