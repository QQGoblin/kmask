package version

import (
	"fmt"
	"runtime"
)

var (
	MainVersion = "0.0.0"
	GitCommit   = "xxxxxxxxxxxxxxxxxxxxxxxxxx"
	BuildDate   = "1970-01-01T00:00:00Z"
)

type info struct {
	Version   string
	GitCommit string
	BuildDate string
	Platform  string
}

func Get() info {

	return info{
		Version:   MainVersion,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
