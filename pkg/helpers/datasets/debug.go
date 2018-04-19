package datasets

import (
	"runtime"

	DEBUG "github.com/computes/go-debug"
)

var debug = DEBUG.Debug("go-sdk:helpers:datasets")
var _debugStack = DEBUG.Debug("go-sdk:helpers:datasets:stack")

var debugStack = func() {
	stack := make([]byte, 1024)
	runtime.Stack(stack, false)

	_debugStack(string(stack))
}
