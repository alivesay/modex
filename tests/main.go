package main

//import "github.com/veandco/go-sdl2/sdl"

import "github.com/alivesay/modex"
import "github.com/alivesay/modex/core"
import "time"
import "runtime"

func init() {
	runtime.LockOSThread()
}

func main() {
	m := modex.NewModex()

	core.Log(core.LOG_NOTICE, "Booting...")
	m.Boot()
	for i := 0; i < 10; i++ {
		m.Update()
		m.Render()
		time.Sleep(time.Second * 1)
	}
	m.Shutdown()

	/*
	   sdl.Init(sdl.INIT_EVERYTHING)

	   window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	       800, 600, sdl.WINDOW_SHOWN)
	   if err != nil {
	       panic(err)
	   }
	   defer window.Destroy()

	   surface, err := window.GetSurface()
	   if err != nil {
	       panic(err)
	   }

	   rect := sdl.Rect{0, 0, 200, 200}
	   surface.FillRect(&rect, 0xffff0000)
	   window.UpdateSurface()

	   sdl.Delay(1000)
	   sdl.Quit()
	*/
}
