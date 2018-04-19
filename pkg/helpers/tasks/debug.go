package tasks

import (
	"runtime"

	DEBUG "github.com/computes/go-debug"
)

var debug = DEBUG.Debug("go-sdk:helpers:tasks")
var _debugStack = DEBUG.Debug("go-sdk:helpers:tasks:stack")

var debugStack = func() {
	stack := make([]byte, 1024)
	runtime.Stack(stack, false)

	_debugStack(string(stack))
}
