package cmd_gb32960

import (
	"bcd-util/support_parse/gb32960"
	"encoding/hex"
	"net"
	"sync"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	const hexStr = "232302fe5445535430303030303030303030303030010141170608100a10010103010040000003520f2827811c012e2000000002010101594fdb4e2f4a0f3227100500073944e501dd620a0601090e1b01370e14010145010444070300021387000000000801010f282781006c00016c0e180e190e1a0e190e190e180e180e1a0e1b0e180e190e1a0e180e180e190e1a0e1a0e190e180e1a0e180e1a0e1a0e180e170e190e170e190e170e190e1b0e190e190e190e180e180e170e170e180e170e170e170e190e170e180e170e190e170e170e170e180e180e190e190e140e180e180e170e170e150e160e160e180e190e170e180e170e180e170e180e170e160e190e150e180e160e180e170e160e160e170e150e170e170e140e170e160e160e170e170e170e170e160e170e160e170e140e170e170e160e160e170e170e170e160e160e160e16090101000c454545444544444445444544f5"
	dial, err := net.Dial("tcp", "gb32960.evsmc.cn:19007")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		defer group.Done()
		buf := make([]byte, 1024)
		for {
			n, err := dial.Read(buf)
			if err != nil {
				t.Errorf("%+v", err)
				return
			}
			t.Logf("<--- %s", hex.EncodeToString(buf[0:n]))
		}
	}()

	//先平台登入
	bs := gb32960.ToPacketBytes(0x05, 0xFE, "TEST0000000000000", &gb32960.PlatformLoginData{
		F_collectTime: time.Now(),
		F_sn:          1,
		F_username:    "RisingAuto_7",
		F_password:    "FeiFanEVCar_20220815",
		F_encode:      1,
	})
	_, err = dial.Write(bs)
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("---> %s", hex.EncodeToString(bs))

	time.Sleep(1 * time.Second)

	//车辆登入
	bs = gb32960.ToPacketBytes(0x01, 0xFE, "TEST0000000000000", &gb32960.VehicleLoginData{
		F_collectTime:   time.Now(),
		F_sn:            1,
		F_iccid:         "00000000000000000000",
		F_subSystemNum:  0,
		F_systemCodeLen: 0,
		F_systemCode:    "",
	})
	_, err = dial.Write(bs)
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("---> %s", hex.EncodeToString(bs))

	time.Sleep(1 * time.Second)

	decodeString, err := hex.DecodeString(hexStr)
	_, err = dial.Write(decodeString)
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("---> %s", hexStr)

	group.Wait()

}
