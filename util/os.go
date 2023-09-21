package util

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var fns []func() error
var fnsLock = sync.Mutex{}

func ExitOnKill(fnOnExit func() error) {
	fnsLock.Lock()
	fns = append(fns, fnOnExit)
	if len(fns) == 1 {
		go func() {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			s := <-c
			Log.Infof("exit signal %+v", s)
			fnsLock.Lock()
			for _, f := range fns {
				if f != nil {
					err := f()
					if err != nil {
						Log.Errorf("%+v", err)
					}
				}
			}
			fnsLock.Unlock()
			os.Exit(0)
		}()
	}
	defer fnsLock.Unlock()
}
