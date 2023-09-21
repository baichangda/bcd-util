package util

import (
	"github.com/pkg/errors"
	"testing"
)

func TestExitOnKill(t *testing.T) {
	ExitOnKill(func() error {
		Log.Info("fn1")
		return nil
	})
	ExitOnKill(func() error {
		return errors.New("fn2")
	})
	ExitOnKill(func() error {
		Log.Info("fn3")
		return nil
	})

}
