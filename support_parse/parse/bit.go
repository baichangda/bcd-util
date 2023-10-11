package parse

import (
	"math/bits"
)

type BitBuf_reader struct {
	bitOffset int
	b         byte
	byteBuf   *ByteBuf
}

func ToBitBuf_reader(byteBuf *ByteBuf) *BitBuf_reader {
	return &BitBuf_reader{
		bitOffset: 0,
		b:         0,
		byteBuf:   byteBuf,
	}
}

func (e *BitBuf_reader) Read(bit int, bigEndian bool, unsigned bool) int64 {
	byteBuf := e.byteBuf
	bitOffset := e.bitOffset
	var b uint8
	if bitOffset == 0 {
		b = byteBuf.Read_uint8()
	} else {
		b = e.b
	}

	temp := bit + bitOffset
	finalBitOffset := temp & 7
	var byteLen int
	var l int64
	if (temp & 7) == 0 {
		byteLen = (temp >> 3)
		byteLen = temp >> 3
		l = int64(b) << (temp - 8)
	} else {
		byteLen = (temp >> 3) + 1
		l = int64(b) << (temp - finalBitOffset)
	}

	for i := 1; i < byteLen; i++ {
		b = byteBuf.Read_uint8()
		l |= (int64(b) << ((byteLen - 1 - i) << 3))
	}

	e.bitOffset = finalBitOffset
	e.b = b

	//如果是小端模式、则翻转bit
	var cRight int64
	if bigEndian {
		cRight = l >> ((byteLen << 3) - bitOffset - bit)
	} else {
		cRight = int64(bits.Reverse64(uint64(l))) >> (64 - (byteLen << 3) + bitOffset)
	}

	if !unsigned && ((cRight>>(bit-1))&0x01) == 1 {
		return cRight | (int64(-1) << bit)
	} else {
		return cRight & ((0x01 << bit) - 1)
	}
}

func (e *BitBuf_reader) Skip(bit int) {
	byteBuf := e.byteBuf
	bitOffset := e.bitOffset
	b := e.b

	temp := bit + bitOffset
	finalBitOffset := temp & 7
	newBitOffsetZero := finalBitOffset == 0
	var byteLen int
	if newBitOffsetZero {
		byteLen = (temp >> 3)
	} else {
		byteLen = (temp >> 3) + 1
	}
	if byteLen == 1 {
		if bitOffset == 0 {
			b = byteBuf.Read_uint8()
		}
	} else {
		if bitOffset == 0 {
			if newBitOffsetZero {
				byteBuf.Skip(byteLen)
			} else {
				byteBuf.Skip(byteLen - 1)
				b = byteBuf.Read_uint8()
			}
		} else {
			if newBitOffsetZero {
				byteBuf.Skip(byteLen - 1)
			} else {
				byteBuf.Skip(byteLen - 2)
				b = byteBuf.Read_uint8()
			}
		}
	}
	e.bitOffset = finalBitOffset
	e.b = b
}

func (e *BitBuf_reader) Finish() {
	e.b = 0
	e.bitOffset = 0
}

type BitBuf_writer struct {
	bitOffset int
	b         byte
	byteBuf   *ByteBuf
}

func ToBitBuf_writer(byteBuf *ByteBuf) *BitBuf_writer {
	return &BitBuf_writer{
		bitOffset: 0,
		b:         0,
		byteBuf:   byteBuf,
	}
}

func (e *BitBuf_writer) Write(v int64, bit int, bigEndian bool, unsigned bool) {
	byteBuf := e.byteBuf
	bitOffset := e.bitOffset
	b := e.b
	if !unsigned && v < 0 {
		v = v & ((0x01 << bit) - 1)
	}
	if !bigEndian {
		v = int64(bits.Reverse64(uint64(v))) >> (64 - bit)
	}
	temp := bit + bitOffset
	finalBitOffset := temp & 7
	var byteLen int
	var newV int64
	if finalBitOffset == 0 {
		byteLen = temp >> 3
		newV = v
	} else {
		byteLen = (temp >> 3) + 1
		newV = v << (8 - finalBitOffset)
	}
	if bitOffset == 0 {
		b = (byte)(newV >> ((byteLen - 1) << 3))
	} else {
		b |= (byte)(newV >> ((byteLen - 1) << 3))
	}
	for i := 1; i < byteLen; i++ {
		byteBuf.Write_uint8(b)
		b = (byte)(newV >> ((byteLen - i - 1) << 3))
	}
	bitOffset = finalBitOffset
	if bitOffset == 0 {
		byteBuf.Write_uint8(b)
	}

	e.bitOffset = bitOffset
	e.b = b
}

func (e *BitBuf_writer) Skip(bit int) {
	byteBuf := e.byteBuf
	bitOffset := e.bitOffset
	b := e.b
	temp := bit + bitOffset
	newBitOffsetZero := (temp & 7) == 0
	var byteLen int
	if newBitOffsetZero {
		byteLen = temp >> 3
	} else {
		byteLen = (temp >> 3) + 1
	}
	if byteLen == 1 {
		if newBitOffsetZero {
			byteBuf.Write_uint8(b)
			b = 0
		}
	} else {
		if bitOffset == 0 {
			if newBitOffsetZero {
				byteBuf.Write_zero(byteLen)
			} else {
				byteBuf.Write_zero(byteLen - 1)
			}

		} else {
			byteBuf.Write_uint8(b)
			if newBitOffsetZero {
				byteBuf.Write_zero(byteLen - 1)
			} else {
				byteBuf.Write_zero(byteLen - 2)
			}
		}
		b = 0
	}
	e.bitOffset = temp & 7
	e.b = b
}

func (e *BitBuf_writer) Finish() {
	if e.bitOffset > 0 {
		e.byteBuf.Write_uint8(e.b)
	}
	e.b = 0
	e.bitOffset = 0
}
