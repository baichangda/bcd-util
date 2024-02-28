package immotors_bin

import "testing"

func TestGetVins(t *testing.T) {
	vins := GetVins(15000, 50000)
	t.Log(vins[0])
	t.Log(vins[len(vins)-1])
}
