package util

import (
	"os"
	"os/signal"
	"sync"
)

var funcs []func() error
var funcsLock = sync.Mutex{}

func ExitOnKill(fnOnExit func() error) {
	funcsLock.Lock()
	defer funcsLock.Unlock()
	funcs = append(funcs, fnOnExit)
	if len(funcs) == 1 {
		go func() {
			c := make(chan os.Signal)
			defer close(c)
			signal.Notify(c, os.Interrupt, os.Kill)
			for v := range c {
				Log.Infof("exit signal %+v", v)
				funcsLock.Lock()
				for _, f := range funcs {
					if f != nil {
						err := f()
						if err != nil {
							Log.Errorf("%+v", err)
						}
					}
				}
				funcsLock.Unlock()
				os.Exit(0)
			}
		}()
	}
}
