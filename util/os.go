package util

import (
	"os"
	"os/signal"
)

func ExitOnKill(fnOnExit func() error) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for v := range c {
			Log.Infof("exit signal %+v", v)
			if fnOnExit != nil {
				err := fnOnExit()
				if err != nil {
					Log.Errorf("%+v", err)
				}
			}
			os.Exit(0)
		}
	}()
}
