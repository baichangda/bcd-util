package parse

import (
	"bcd-util/util"
	"sync/atomic"
	"time"
)

func TestMultiThreadPerformance_parse(bytes []byte, threadNum int, num int, fn func(*ByteBuf)) {
	var count uint64 = 0
	var runNum int32 = 0
	for i := 0; i < threadNum; i++ {
		atomic.AddInt32(&runNum, 1)
		go func() {
			byteBuf := ToByteBuf(bytes)
			byteBuf.MarkReaderIndex()
			byteBuf.MarkWriterIndex()
			for j := 0; j < num; j++ {
				byteBuf.ResetReaderIndex()
				byteBuf.ResetWriterIndex()
				fn(byteBuf)
				atomic.AddUint64(&count, 1)
			}
			atomic.AddInt32(&runNum, -1)
		}()
	}

	util.ExitOnKill(nil)

	for {
		time.Sleep(3 * time.Second)
		if atomic.LoadInt32(&runNum) == 0 {
			break
		}
		val := atomic.SwapUint64(&count, 0) / 3
		util.Log.Infof("parse threadNum:%d num:%d totalSpeed/s:%d perThreadSpeed/s:%d", threadNum, num, val, val/uint64(threadNum))
	}
}

func TestMultiThreadPerformance_deParse(byteBuf *ByteBuf, threadNum int, num int, fn func(*ByteBuf)) {
	var count uint64 = 0
	var runNum int32 = 0
	for i := 0; i < threadNum; i++ {
		atomic.AddInt32(&runNum, 1)
		go func() {
			byteBuf.MarkReaderIndex()
			byteBuf.MarkWriterIndex()
			for j := 0; j < num; j++ {
				byteBuf.ResetReaderIndex()
				byteBuf.ResetWriterIndex()
				fn(byteBuf)
				atomic.AddUint64(&count, 1)
			}
			atomic.AddInt32(&runNum, -1)
		}()
	}

	util.ExitOnKill(nil)

	for {
		time.Sleep(3 * time.Second)
		if atomic.LoadInt32(&runNum) == 0 {
			break
		}
		val := atomic.SwapUint64(&count, 0) / 3
		util.Log.Infof("deParse threadNum:%d num:%d totalSpeed/s:%d perThreadSpeed/s:%d", threadNum, num, val, val/uint64(threadNum))
	}
}
