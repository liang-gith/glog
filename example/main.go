package main

import (
	"github.com/x-mod/glog"
)

func main() {
	glog.Open(
		glog.LogDir("logs"),
		// glog.LogToStderr(true),
		glog.AlsoLogToStderr(true),
		glog.Verbosity(2),
	)
	defer glog.Flush()

	go func() {
		glog.Info("go routine f1.... ok")
	}()

	go func() {
		glog.Info("go routine f2.... ok")
	}()
	glog.Error("error .... ok")
	glog.Warning("warning .... ok")
	glog.Info("info .... ok")
	glog.V(1).Info("V(1) info .... ok")
	glog.V(2).Info("V(2) info .... ok")
	glog.V(3).Info("V(3) info .... ok")
	glog.V(4).Info("V(4) info .... ok")
	glog.V(5).Info("V(5) info .... ok")
}
