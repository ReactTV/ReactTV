package version

import (
	"fmt"
	"runtime"
)

var (
	BinVersion = "unknown"
	GitCommit  = "unknown"
	BuildTime  = "unknown"
)

func Runtime() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
