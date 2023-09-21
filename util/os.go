package util

import (
	"os"
	"os/signal"
)

func ExitOnKill(fnOnExit func() error) {
	go func() {
		c := make(chan os.Signal)
		defer close(c)
		signal.Notify(c, os.Interrupt, os.Kill)
		s := <-c
		Log.Infof("exit signal %+v", s)
		if fnOnExit != nil {
			err := fnOnExit()
			if err != nil {
				Log.Errorf("%+v", err)
			}
		}
	}()
}
