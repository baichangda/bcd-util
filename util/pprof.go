package util

import (
	"net/http"
	"net/http/pprof"
)

// StartWeb_pprof 开启pprof性能采集web服务、监听在54321端口
// 采集cpu信息到文件 curl -o cpu.out http://127.0.0.1:54321/debug/pprof/profile?seconds=30
// 采集内存信息到文件 curl -o mem.out http://127.0.0.1:54321/debug/pprof/allocs?seconds=30
func StartWeb_pprof() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		Log.Infof("start pprof web ListenAndServe[:54321]")
		err := http.ListenAndServe(":54321", mux)
		if err != nil {
			Log.Errorf("%+v", err)
			return
		}
	}()
}
