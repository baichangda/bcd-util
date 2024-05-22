package gb32960

import (
	"bcd-util/support_parse/parse"
	"encoding/hex"
	"runtime/debug"
	"strings"
	"testing"
)

func BenchmarkToPacket(b *testing.B) {
	//debug.SetGCPercent(400)
	//hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	hexStr := "232302FE4C534A4533363039364D53313430343935010141170608100A10010103010040000003520F2827811C012E2000000002010101594FDB4E2F4A0F3227100500073944E501DD620A0601090E1B01370E14010145010444070300021387000000000801010F282781006C00016C0E180E190E1A0E190E190E180E180E1A0E1B0E180E190E1A0E180E180E190E1A0E1A0E190E180E1A0E180E1A0E1A0E180E170E190E170E190E170E190E1B0E190E190E190E180E180E170E170E180E170E170E170E190E170E180E170E190E170E170E170E180E180E190E190E140E180E180E170E170E150E160E160E180E190E170E180E170E180E170E180E170E160E190E150E180E160E180E170E160E160E170E150E170E170E140E170E160E160E170E170E170E170E160E170E160E170E140E170E170E160E160E170E170E170E160E160E160E16090101000C454545444544444445444544F5"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		b.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	byteBuf.MarkReaderIndex()
	byteBuf.MarkWriterIndex()
	for i := 0; i < b.N; i++ {
		byteBuf.ResetReaderIndex()
		byteBuf.ResetWriterIndex()
		_ = To_Packet(byteBuf)
	}
}

func BenchmarkPacket_Write(b *testing.B) {
	//hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	//hexStr := "232302FE4C534A57543430393150533939313935360101F9170B1D0D1E16010303010000000014B41D2C271525010F2000000002010104384E204E20371D2E27100500073942A601DD5FAC0601030E4E01440E48010239010138070000000000000000000801011D2C271500CC0001C80E4A0E4A0E4E0E4C0E4A0E4C0E4B0E4A0E4A0E490E4B0E4C0E4C0E4B0E4D0E4C0E4A0E4B0E4A0E4B0E4C0E4A0E4A0E4B0E4B0E4B0E4B0E4C0E4B0E4D0E4C0E4C0E4C0E4A0E4A0E4A0E4A0E490E4B0E4A0E4A0E4A0E4A0E4A0E490E4A0E4B0E4E0E4A0E4B0E4C0E4A0E4B0E4A0E4C0E4A0E4A0E490E4B0E490E4C0E490E4C0E4C0E490E490E4C0E480E4A0E4B0E4B0E4B0E4A0E4A0E4A0E4C0E4B0E490E4C0E4A0E4A0E4B0E4A0E4C0E4B0E4B0E4C0E4B0E4A0E4B0E4A0E4E0E4B0E4B0E4B0E4A0E4C0E4A0E4C0E4D0E4A0E4A0E4C0E4A0E4B0E4E0E4B0E4A0E4C0E4C0E4B0E4A0E4A0E4C0E4D0E4D0E4B0E4B0E4C0E4B0E4C0E4B0E4B0E4D0E4D0E4A0E4C0E4C0E4C0E4C0E4C0E4B0E4D0E4A0E4C0E4A0E4B0E4C0E4B0E4B0E4C0E4C0E4C0E4A0E4B0E4B0E4A0E4B0E4C0E4D0E4A0E4A0E490E4A0E4A0E4B0E490E4A0E4A0E4A0E4B0E4B0E4D0E4D0E4A0E4D0E4B0E4C0E4D0E4B0E490E490E490E490E490E480E490E490E4B0E4B0E4B0E4C0E4B0E4C0E4B0E4B0E4C0E4B0E4A0E4A0E4B0E4B0E4A0E4B0E4B0E4B0E4C0E4A0E4D0E4A090101000C3839383938393839383938394E"
	hexStr := "232302FE4C534A5754343039315053393931393536010201170C010E1319010103010000000014B41CB1271D17011F2000000002010104324E204E1B311CAC271005000739429E01DD5FA40601030E1301B00E0B010132010B31070000000000000000000801011CB1271D00CC0001CC0E100E0F0E130E110E110E100E100E100E0F0E0E0E0F0E0F0E100E110E110E110E0F0E120E0E0E0E0E110E0E0E100E100E0F0E100E100E120E100E100E0F0E110E120E120E0E0E110E0E0E0E0E0E0E0E0E0F0E0E0E100E0F0E0D0E0F0E110E130E0D0E0E0E110E0F0E100E0E0E100E0F0E100E0E0E110E0F0E0F0E0E0E110E110E0D0E0E0E120E0D0E0E0E0F0E100E0F0E0E0E100E0E0E110E100E0D0E100E0F0E0E0E100E0E0E120E0F0E0E0E100E0F0E0E0E100E0F0E130E110E0E0E110E100E100E0F0E100E110E0F0E0E0E0F0E0C0E100E130E0F0E0E0E0E0E100E0F0E0E0E0E0E110E110E100E100E0F0E100E0F0E100E0F0E0F0E100E100E0F0E110E100E100E110E100E0F0E130E0E0E100E100E0F0E100E0E0E0E0E110E100E100E0E0E0F0E0D0E0F0E0F0E120E100E0E0E0E0E0D0E0E0E0F0E0F0E0C0E0F0E0E0E0F0E100E0D0E0D0E100E0E0E110E0D0E110E100E0E0E0C0E0E0E0E0E0D0E0D0E0B0E0C0E0E0E100E0F0E0F0E0F0E0E0E0F0E0E0E100E110E0F0E0F0E0E0E100E100E0E0E0F0E0E0E0E0E0F0E0D0E100E0E0E0C0E110E0E0E0F090101000C323232323232323232323132E9"

	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		b.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := To_Packet(byteBuf)
	res := parse.ToByteBuf_capacity(1024)
	res.MarkReaderIndex()
	res.MarkWriterIndex()
	for i := 0; i < b.N; i++ {
		res.ResetReaderIndex()
		res.ResetWriterIndex()
		packet.Write(res)
	}
}

func TestPacket_Write(t *testing.T) {
	//hexStr := "232301FE4C534A574C34303933505330333838353801001E17060F091211000138393836303932323738303030333936343832350000B6"
	//hexStr := "232302fe4c534a4533363039364d53313430343935010141170a1b0f1927010103010040000003520f2827811c012e2000000002010101594fdb4e2f4a0f3227100500073944e501dd620a0601090e1b01370e14010145010444070300021387000000000801010f282781006c00016c0e180e190e1a0e190e190e180e180e1a0e1b0e180e190e1a0e180e180e190e1a0e1a0e190e180e1a0e180e1a0e1a0e180e170e190e170e190e170e190e1b0e190e190e190e180e180e170e170e180e170e170e170e190e170e180e170e190e170e170e170e180e180e190e190e140e180e180e170e170e150e160e160e180e190e170e180e170e180e170e180e170e160e190e150e180e160e180e170e160e160e170e150e170e170e140e170e160e160e170e170e170e170e160e170e160e170e140e170e170e160e160e170e170e170e160e160e160e16090101000c454545444544444445444544f5"
	//hexStr := "232302fe4c534a4533363039364d5331343034393501014d170608103933010103010000001726ae10fc275154011f20000000020201044c4e204e1b431108271002044b4e204e204511082710050007390f7b01dd636e0601200fc101020fb901013e01023c0703000217870000000008010110fc2751006c00016c0fbf0fb90fbf0fbd0fbf0fbb0fbf0fbd0fc00fbc0fc00fbb0fba0fbb0fbb0fbb0fbd0fb90fbb0fbc0fbb0fbd0fbc0fbe0fba0fbf0fbe0fc00fbd0fbe0fbc0fc10fbe0fbd0fbb0fbf0fbe0fb90fbe0fbc0fbe0fbc0fb90fb90fbe0fbd0fbf0fbd0fbe0fbc0fc00fbc0fbf0fbd0fbd0fbf0fbc0fbe0fbe0fbe0fbd0fc00fbe0fbd0fbd0fbf0fbf0fc00fc00fc00fbf0fc10fbe0fbc0fbf0fbe0fbf0fbc0fbf0fbd0fc10fbe0fbe0fbb0fbf0fbe0fbf0fbd0fc10fbc0fbc0fbf0fbc0fc00fbf0fbe0fbd0fbe0fbf0fbf0fbd0fc10fbf0fbf0fbe0fc00fbc0fc1090101000c3e3c3c3c3e3c3c3c3d3c3c3c1f"
	//hexStr := "232301FE4C534A574C34303933505330333838353801001E17060A0A0301000138393836303932323738303030333936343832350000B1"
	//hexStr := "232301FE4C534A4533363039584D5331343035303201001E17080B0B122A00013839383630393231373530303038393532333831000095"
	//hexStr := "232302FE4C534A575434303958505339393232353001013F18010410170601030301000000007D000ED6235733011F2000000005000739428101DD613806010B0D4401260D3901023C01013B070000000000000000000801010ED6235700700001700D3F0D3F0D420D3F0D410D400D400D3E0D420D3F0D440D400D3F0D3E0D3C0D3F0D410D430D3D0D3F0D3D0D3F0D3E0D3E0D3D0D3F0D3D0D3D0D3C0D3B0D3E0D3B0D3C0D3C0D3D0D3D0D3C0D390D3C0D3A0D3C0D3A0D3A0D3B0D390D3A0D3A0D390D3A0D3A0D390D390D3B0D390D3B0D3A0D3C0D390D3C0D3A0D3D0D3B0D3D0D3C0D3D0D3C0D410D3B0D3E0D3B0D3B0D3A0D3C0D3F0D3A0D3D0D3B0D3B0D3C0D3C0D3C0D3B0D3B0D3A0D3C0D3C0D3C0D3A0D3F0D3A0D3E0D3C0D3D0D390D3B0D3A0D3E0D3D0D390D3A0D3B0D3A0D3B0D3B0D3A0D3A0D3A0D3B0D3B0D3B0D3D0D3B09010100103B3C3C3C3C3C3C3C3B3C3C3C3C3C3C3CA5"
	//hexStr := "232302024c534a574d3430393850533037373131380100061803140b10386a"
	//hexStr := "2323050154455354303030303030303030303030300100061803140f271d1f"
	//hexStr := "2323010254455354303030303030303030303030300100061803140f271e1b"
	hexStr := "232303fe4c534a574b343039314e5331313933383701014d160a13003121010203010000ffffffff0fa3272735010f20000000020201033b4e204e203e0fb403e702fffefffffffffe0000ffff050100000000000000000601250e8701080e8401023d01013c070000000000000000000801010fa22727006c01006c0e850e860e860e850e850e850e850e840e850e850e860e850e850e850e850e850e850e850e840e840e850e840e840e840e850e840e840e850e840e850e850e840e850e840e840e840e870e860e850e860e860e860e860e860e860e860e860e850e860e860e860e850e850e850e840e840e840e840e840e840e840e850e850e850e850e850e840e850e850e850e840e850e850e850e850e850e850e850e850e860e850e850e850e850e860e850e850e850e860e850e860e850e850e840e850e860e850e850e850e860e840e860e850e850e850e860e860e86090101000c3c3d3c3d3c3d3c3d3c3c3c3c68"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := To_Packet(byteBuf)
	res := parse.ToByteBuf_empty()
	packet.Write(res)
	res_bytes := res.ToBytes()
	res_hex := strings.ToUpper(hex.EncodeToString(res_bytes))
	t.Logf("%s", hexStr)
	t.Logf("%s", res_hex)
	if hexStr != res_hex {
		t.Fail()
	}
}

func TestPerformance_parse(t *testing.T) {
	//util.StartWeb_pprof()
	debug.SetGCPercent(400)
	hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	parse.TestMultiThreadPerformance_parse(decodeString, 1, 1000000000, func(byteBuf *parse.ByteBuf) {
		To_Packet(byteBuf)
		//util.Log.Infof("%d", byteBuf.ReaderIndex())
	})
}

func TestPerformance_deParse(t *testing.T) {
	hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	packet := To_Packet(parse.ToByteBuf(decodeString))
	buf := parse.ToByteBuf_capacity(1024)
	parse.TestMultiThreadPerformance_deParse(buf, 1, 1000000000, func(byteBuf *parse.ByteBuf) {
		packet.Write(byteBuf)
		//util.Log.Infof("%d", byteBuf.WriterIndex())
	})
}

func Test123(t *testing.T) {
	var n int8 = -100
	t.Log(int16(n))
}
