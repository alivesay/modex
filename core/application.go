package core

import (
	"sync"
)

type Application struct {
	Running           bool
	ShutdownRequested bool
}

var applicationInstance *Application
var applicationOnce sync.Once

func GetInstanceApplication() *Application {
	applicationOnce.Do(func() {
		applicationInstance = &Application{}
	})
	return applicationInstance
}
