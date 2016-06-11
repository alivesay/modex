package main

import (
	"github.com/alivesay/modex"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

const profEnabled = false

func init() {
	runtime.LockOSThread()
}

func main() {
	m := modex.NewModex()
	m.Boot()

	if profEnabled {
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	}
	m.Run()
	m.Shutdown()
}
