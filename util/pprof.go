package util

import (
	"net/http"
	"net/http/pprof"
)

// StartWeb_pprof 开启pprof性能采集web服务、监听在54321端口
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
