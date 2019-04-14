package errorx

import "github.com/yakaa/cuter/lib/logx"

func Rescue(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logx.Severe(p)
	}
}
