package dependencies

import "go.uber.org/fx"

type Foo string

type Bar string

type Baz string

type DependenciesType struct {
	fx.In
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

func NewDependencies(p DependenciesParams) *DependenciesType {
	d := Construct(p, 
	return *DependenciesType.(d)
}

var DependenciesModule = fx.Module("dependencies",
	fx.Provide(NewFoo),
	fx.Provide(NewBar),
	fx.Provide(NewBaz),
	fx.Populate(&d),
)
