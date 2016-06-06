package modex

import "runtime"

func init() {
	runtime.LockOSThread()
}

func main() {
	m := NewModex()
	m.Boot()
	m.Run()
	m.Shutdown()
}
