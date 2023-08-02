This version gives the following output:
```
[Fx] PROVIDE	*config.Config <= example.com/fxtest/config.NewConfig()
[Fx] PROVIDE	*main.Server <= main.NewServer()
[Fx] PROVIDE	*services.SomeServiceType <= example.com/fxtest/services.NewSomeService() from module "services"
[Fx] PROVIDE	fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE	fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE	fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE		example.com/fxtest/services.glob..func1() from module "services"
[Fx] ERROR		fx.Invoke(example.com/fxtest/services.glob..func1()) called from:
example.com/fxtest/services.init
	/Users/gregory/g/git/fxtest/services/services.go:31
runtime.doInit
	/usr/local/go/src/runtime/proc.go:6506
runtime.doInit
	/usr/local/go/src/runtime/proc.go:6483
runtime.main
	/usr/local/go/src/runtime/proc.go:233
Failed: missing dependencies for function "example.com/fxtest/services".glob..func1
	/Users/gregory/g/git/fxtest/services/services.go:31:
missing type:
	- services.SomeServiceType (did you mean to Provide it?)
[Fx] ERROR		Failed to start: missing dependencies for function "example.com/fxtest/services".glob..func1
	/Users/gregory/g/git/fxtest/services/services.go:31:
missing type:
	- services.SomeServiceType (did you mean to Provide it?)
```
