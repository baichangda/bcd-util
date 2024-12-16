package immotors

import (
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"
)

func BenchmarkToPacket(b *testing.B) {
	//debug.SetGCPercent(1000)
	hexStr := "000100006466E9B3000300006583EC6E000413C339C87BDE00051DD636E16BEC00060000050000070007F024FFF8FFBC000801CC0001000000094300C4100033000AB70400620000000B000000000000000C0017AE000000000D000992400000000E082C00000000000F5208ABE0000000104E934FE20000001100000000000008008169000000000801000000000000080200000000000008030250AB000000D006003E8D3A8987B4A0365D34F9A820FFFC0080180100060699FC9700008000A0001FF90050010FF915B2001B48006D011F016F26B7E0C95F722400B4F800000000D008004600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D00900120000000000000000000FE000010005648000D00A00554C534A4533363039364D533134303439354C53323143303141353239363030303638393836303932313735303030383935323034333839303836303334323032323030303031303231303030303333353932343437D00B00D96C7DF87DC87DF87DE87DF87DD87DF87DE87E007DE07E007DD87DD07DD87DD87DD87DE87DC87DD87DE07DD87DE87DE07DF07DD07DF87DF07E007DE87DF07DE07E087DF07DE87DD87DF87DF07DC87DF07DE07DF07DE07DC87DC87DF07DE87DF87DE87DF07DE07E007DE07DF87DE87DE87DF87DE07DF07DF07DF07DE87E007DF07DE87DE87DF87DF87E007E007E007DF87E087DF07DE07DF87DF07DF87DE07DF87DE87E087DF07DF07DD87DF87DF07DF87DE87E087DE07DE07DF87DE07E007DF87DF07DE87DF07DF87DF87DE87E087DF87DF87DF07E007DE07E08D00C00190C3E003C003C003C003E003C003C003C003D003C003C003C00D00D00190C3F003D003D003E003E003D003D003E003E003C003D003E00D00E00191830354C50454A334333353234304142434230313032303135D00F000C007D00007831800019007E00D01000380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01100230000000000000000000000000000000000000000000000000000000000000000000000D012003F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D013004D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01400230000000000000000000000000000000000000000000000000000000000000000000000D015000700000000000000D016003100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D017002A000000000000000000000000000000000000000000000000000000000000000000000000000000000000D0180014CF639C87BB0EEB1B720000323374D920FA040300D0190010CE10053DB55580000000000000000000D01A002D000000000000000000000000000000000031323334353637383930313233343536373839303132333435363738D01B003F000000313233343536373839303132333435363738000031323334353637383930313233343536373800003132333435363738393031323334353637380000D01C002E31323334353637383930313233343536373831323334353637383930313233343536370000000000000000000000D01D000C000000000000000000000000D01F001000000000000000000000000000000000D020000E0000000000000000000000000000"
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
	hexStr := "000100006466E9B3000300006583EC6E000413C339C87BDE00051DD636E16BEC00060000050000070007F024FFF8FFBC000801CC0001000000094300C4100033000AB70400620000000B000000000000000C0017AE000000000D000992400000000E082C00000000000F5208ABE0000000104E934FE20000001100000000000008008169000000000801000000000000080200000000000008030250AB000000D006003E8D3A8987B4A0365D34F9A820FFFC0080180100060699FC9700008000A0001FF90050010FF915B2001B48006D011F016F26B7E0C95F722400B4F800000000D008004600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D00900120000000000000000000FE000010005648000D00A00554C534A4533363039364D533134303439354C53323143303141353239363030303638393836303932313735303030383935323034333839303836303334323032323030303031303231303030303333353932343437D00B00D96C7DF87DC87DF87DE87DF87DD87DF87DE87E007DE07E007DD87DD07DD87DD87DD87DE87DC87DD87DE07DD87DE87DE07DF07DD07DF87DF07E007DE87DF07DE07E087DF07DE87DD87DF87DF07DC87DF07DE07DF07DE07DC87DC87DF07DE87DF87DE87DF07DE07E007DE07DF87DE87DE87DF87DE07DF07DF07DF07DE87E007DF07DE87DE87DF87DF87E007E007E007DF87E087DF07DE07DF87DF07DF87DE07DF87DE87E087DF07DF07DD87DF87DF07DF87DE87E087DE07DE07DF87DE07E007DF87DF07DE87DF07DF87DF87DE87E087DF87DF87DF07E007DE07E08D00C00190C3E003C003C003C003E003C003C003C003D003C003C003C00D00D00190C3F003D003D003E003E003D003D003E003E003C003D003E00D00E00191830354C50454A334333353234304142434230313032303135D00F000C007D00007831800019007E00D01000380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01100230000000000000000000000000000000000000000000000000000000000000000000000D012003F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D013004D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01400230000000000000000000000000000000000000000000000000000000000000000000000D015000700000000000000D016003100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D017002A000000000000000000000000000000000000000000000000000000000000000000000000000000000000D0180014CF639C87BB0EEB1B720000323374D920FA040300D0190010CE10053DB55580000000000000000000D01A002D000000000000000000000000000000000031323334353637383930313233343536373839303132333435363738D01B003F000000313233343536373839303132333435363738000031323334353637383930313233343536373800003132333435363738393031323334353637380000D01C002E31323334353637383930313233343536373831323334353637383930313233343536370000000000000000000000D01D000C000000000000000000000000D01F001000000000000000000000000000000000D020000E0000000000000000000000000000"
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
	hexStr := "000100006466E9B3000300006583EC6E000413C339C87BDE00051DD636E16BEC00060000050000070007F024FFF8FFBC000801CC0001000000094300C4100033000AB70400620000000B000000000000000C0017AE000000000D000992400000000E082C00000000000F5208ABE0000000104E934FE20000001100000000000008008169000000000801000000000000080200000000000008030250AB000000D006003E8D3A8987B4A0365D34F9A820FFFC0080180100060699FC9700008000A0001FF90050010FF915B2001B48006D011F016F26B7E0C95F722400B4F800000000D008004600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D00900120000000000000000000FE000010005648000D00A00554C534A4533363039364D533134303439354C53323143303141353239363030303638393836303932313735303030383935323034333839303836303334323032323030303031303231303030303333353932343437D00B00D96C7DF87DC87DF87DE87DF87DD87DF87DE87E007DE07E007DD87DD07DD87DD87DD87DE87DC87DD87DE07DD87DE87DE07DF07DD07DF87DF07E007DE87DF07DE07E087DF07DE87DD87DF87DF07DC87DF07DE07DF07DE07DC87DC87DF07DE87DF87DE87DF07DE07E007DE07DF87DE87DE87DF87DE07DF07DF07DF07DE87E007DF07DE87DE87DF87DF87E007E007E007DF87E087DF07DE07DF87DF07DF87DE07DF87DE87E087DF07DF07DD87DF87DF07DF87DE87E087DE07DE07DF87DE07E007DF87DF07DE87DF07DF87DF87DE87E087DF87DF87DF07E007DE07E08D00C00190C3E003C003C003C003E003C003C003C003D003C003C003C00D00D00190C3F003D003D003E003E003D003D003E003E003C003D003E00D00E00191830354C50454A334333353234304142434230313032303135D00F000C007D00007831800019007E00D01000380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01100230000000000000000000000000000000000000000000000000000000000000000000000D012003F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D013004D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01400230000000000000000000000000000000000000000000000000000000000000000000000D015000700000000000000D016003100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D017002A000000000000000000000000000000000000000000000000000000000000000000000000000000000000D0180014CF639C87BB0EEB1B720000323374D920FA040300D0190010CE10053DB55580000000000000000000D01A002D000000000000000000000000000000000031323334353637383930313233343536373839303132333435363738D01B003F000000313233343536373839303132333435363738000031323334353637383930313233343536373800003132333435363738393031323334353637380000D01C002E31323334353637383930313233343536373831323334353637383930313233343536370000000000000000000000D01D000C000000000000000000000000D01F001000000000000000000000000000000000D020000E0000000000000000000000000000"
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

	for i := range hexStr {
		if hexStr[i] != res_hex[i] {
			t.Logf("%d %s %s", i,
				hexStr[i-3:i+3], res_hex[i-3:i+3])
			break
		}
	}
	if hexStr != res_hex {
		t.Fail()
	}
}

func TestPacket_Write_base64(t *testing.T) {
	base64Str := "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"
	bytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	unGzip, err := util.UnGzip(bytes)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	hexStr := strings.ToUpper(hex.EncodeToString(unGzip))
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packets := To_Packets(byteBuf)
	res := parse.ToByteBuf_empty()
	Write_Packets(packets, res)
	res_bytes := res.ToBytes()
	res_hex := strings.ToUpper(hex.EncodeToString(res_bytes))
	t.Logf("%s", hexStr)
	t.Logf("%s", res_hex)

	for i := range hexStr {
		if hexStr[i] != res_hex[i] {
			t.Logf("%d %s %s", i,
				hexStr[i:i+1], res_hex[i:i+1])
			break
		}
	}
	if hexStr != res_hex {
		t.Fail()
	}
}

func TestPerformance_parse(t *testing.T) {
	//debug.SetGCPercent(500)
	hexStr := "000100006466E9B3000300006583EC6E000413C339C87BDE00051DD636E16BEC00060000050000070007F024FFF8FFBC000801CC0001000000094300C4100033000AB70400620000000B000000000000000C0017AE000000000D000992400000000E082C00000000000F5208ABE0000000104E934FE20000001100000000000008008169000000000801000000000000080200000000000008030250AB000000D006003E8D3A8987B4A0365D34F9A820FFFC0080180100060699FC9700008000A0001FF90050010FF915B2001B48006D011F016F26B7E0C95F722400B4F800000000D008004600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D00900120000000000000000000FE000010005648000D00A00554C534A4533363039364D533134303439354C53323143303141353239363030303638393836303932313735303030383935323034333839303836303334323032323030303031303231303030303333353932343437D00B00D96C7DF87DC87DF87DE87DF87DD87DF87DE87E007DE07E007DD87DD07DD87DD87DD87DE87DC87DD87DE07DD87DE87DE07DF07DD07DF87DF07E007DE87DF07DE07E087DF07DE87DD87DF87DF07DC87DF07DE07DF07DE07DC87DC87DF07DE87DF87DE87DF07DE07E007DE07DF87DE87DE87DF87DE07DF07DF07DF07DE87E007DF07DE87DE87DF87DF87E007E007E007DF87E087DF07DE07DF87DF07DF87DE07DF87DE87E087DF07DF07DD87DF87DF07DF87DE87E087DE07DE07DF87DE07E007DF87DF07DE87DF07DF87DF87DE87E087DF87DF87DF07E007DE07E08D00C00190C3E003C003C003C003E003C003C003C003D003C003C003C00D00D00190C3F003D003D003E003E003D003D003E003E003C003D003E00D00E00191830354C50454A334333353234304142434230313032303135D00F000C007D00007831800019007E00D01000380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01100230000000000000000000000000000000000000000000000000000000000000000000000D012003F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D013004D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01400230000000000000000000000000000000000000000000000000000000000000000000000D015000700000000000000D016003100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D017002A000000000000000000000000000000000000000000000000000000000000000000000000000000000000D0180014CF639C87BB0EEB1B720000323374D920FA040300D0190010CE10053DB55580000000000000000000D01A002D000000000000000000000000000000000031323334353637383930313233343536373839303132333435363738D01B003F000000313233343536373839303132333435363738000031323334353637383930313233343536373800003132333435363738393031323334353637380000D01C002E31323334353637383930313233343536373831323334353637383930313233343536370000000000000000000000D01D000C000000000000000000000000D01F001000000000000000000000000000000000D020000E0000000000000000000000000000"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	parse.TestMultiThreadPerformance_parse(decodeString, 1, 1000000000, func(byteBuf *parse.ByteBuf) {
		To_Packets(byteBuf)
		//util.Log.Infof("%d", byteBuf.ReaderIndex())
	})
}

func TestPerformance_deParse(t *testing.T) {
	hexStr := "000100006466E9B3000300006583EC6E000413C339C87BDE00051DD636E16BEC00060000050000070007F024FFF8FFBC000801CC0001000000094300C4100033000AB70400620000000B000000000000000C0017AE000000000D000992400000000E082C00000000000F5208ABE0000000104E934FE20000001100000000000008008169000000000801000000000000080200000000000008030250AB000000D006003E8D3A8987B4A0365D34F9A820FFFC0080180100060699FC9700008000A0001FF90050010FF915B2001B48006D011F016F26B7E0C95F722400B4F800000000D008004600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D00900120000000000000000000FE000010005648000D00A00554C534A4533363039364D533134303439354C53323143303141353239363030303638393836303932313735303030383935323034333839303836303334323032323030303031303231303030303333353932343437D00B00D96C7DF87DC87DF87DE87DF87DD87DF87DE87E007DE07E007DD87DD07DD87DD87DD87DE87DC87DD87DE07DD87DE87DE07DF07DD07DF87DF07E007DE87DF07DE07E087DF07DE87DD87DF87DF07DC87DF07DE07DF07DE07DC87DC87DF07DE87DF87DE87DF07DE07E007DE07DF87DE87DE87DF87DE07DF07DF07DF07DE87E007DF07DE87DE87DF87DF87E007E007E007DF87E087DF07DE07DF87DF07DF87DE07DF87DE87E087DF07DF07DD87DF87DF07DF87DE87E087DE07DE07DF87DE07E007DF87DF07DE87DF07DF87DF87DE87E087DF87DF87DF07E007DE07E08D00C00190C3E003C003C003C003E003C003C003C003D003C003C003C00D00D00190C3F003D003D003E003E003D003D003E003E003C003D003E00D00E00191830354C50454A334333353234304142434230313032303135D00F000C007D00007831800019007E00D01000380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01100230000000000000000000000000000000000000000000000000000000000000000000000D012003F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D013004D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01400230000000000000000000000000000000000000000000000000000000000000000000000D015000700000000000000D016003100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D017002A000000000000000000000000000000000000000000000000000000000000000000000000000000000000D0180014CF639C87BB0EEB1B720000323374D920FA040300D0190010CE10053DB55580000000000000000000D01A002D000000000000000000000000000000000031323334353637383930313233343536373839303132333435363738D01B003F000000313233343536373839303132333435363738000031323334353637383930313233343536373800003132333435363738393031323334353637380000D01C002E31323334353637383930313233343536373831323334353637383930313233343536370000000000000000000000D01D000C000000000000000000000000D01F001000000000000000000000000000000000D020000E0000000000000000000000000000"
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
