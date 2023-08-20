package dependencies

import (
	"example.com/fxtest/utils"
	"go.uber.org/fx"
)

type Foo *string

type Bar *string

type Baz *string

type DepFields struct {
	Foo Foo
	Bar Bar
	Baz Baz
}

type DependenciesType struct {
	DepFields
}

type DependenciesParams struct {
	fx.In

	DepFields
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

func NewDependencies(p DependenciesParams) *DependenciesType {
	retval := utils.Construct[DependenciesParams, DependenciesType](p)
	return retval
}

var DependenciesModule = fx.Module("dependencies",
	fx.Provide(NewFoo),
	fx.Provide(NewBar),
	fx.Provide(NewBaz),
	// This is so we can reuse the fields from DepFields struct in both parameter object
	// (DependenciesParams) and return object (DependenciesType).
	fx.Provide(func() DepFields {
		return DepFields{}
	}),
	fx.Provide(NewDependencies),
)
