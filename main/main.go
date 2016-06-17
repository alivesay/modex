package main

import (
	"github.com/alivesay/modex"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

const profEnabled = true

func init() {
	runtime.LockOSThread()
}

func main() {
	if profEnabled {
		f, err := os.Create("modex.prof")
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	m := modex.NewModex()
	m.Boot()

	if profEnabled && false {
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	}
	m.Run()
	m.Shutdown()
}
