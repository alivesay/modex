package main

import (
	"github.com/alivesay/modex"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	m := modex.NewModex()
	m.Boot()
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	m.Run()
	m.Shutdown()
}
