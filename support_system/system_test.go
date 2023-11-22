package support_system

import (
	"bcd-util/util"
	"encoding/json"
	"github.com/robfig/cron/v3"
	"testing"
)

func TestCollect(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("*/1 * * * * ?", func() {
		collect, err := Collect()
		if err != nil {
			t.Fatal(err)
		}
		marshal, err := json.Marshal(collect)
		if err != nil {
			t.Fatal(err)
		}
		util.Log.Infof("%s", string(marshal))
	})
	if err != nil {
		t.Fatal(err)
	}
	c.Run()
}
