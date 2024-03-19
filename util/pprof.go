package util

import (
	"net/http"
	"net/http/pprof"
)

func StartWeb_pprof(addr string) {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		Log.Infof("start pprof web ListenAndServe[%s]", addr)
		err := http.ListenAndServe(addr, mux)
		if err != nil {
			Log.Errorf("%+v", err)
			return
		}
	}()
}
