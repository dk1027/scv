package shared

import (
	"plugin"
)

type IRule interface {
	Fire(runtimeParam string)
}
type IAction interface {
	Run(param string) // Run the command and wait for its completion
}

type ITrigger interface {
	Run()   // Expected to be run as a goroutine
	Close() // Called during tear down to clean up resources
}

type CreateTriggerFnT = func(IRule, string) ITrigger

func TriggerLoader(libpath string) CreateTriggerFnT {
	lib, err := plugin.Open(libpath)
	if err != nil {
		panic(err)
	}
	sym, err := lib.Lookup("CreateTriggerImpl")
	CHECK_ERR(err)

	fn, ok := sym.(CreateTriggerFnT)
	CHECK_OK(ok, "Incorrect Type")

	return fn
}
