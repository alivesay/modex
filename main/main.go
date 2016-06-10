package main

import (
	"github.com/alivesay/modex"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	m := modex.NewModex()
	m.Boot()
	m.Run()
	m.Shutdown()
}
