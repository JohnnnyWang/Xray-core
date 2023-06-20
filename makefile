GOOS_LINUX = linux
GOOS_WINDOWS = windows
GOARCH = amd64
CGO_ENABLED = 0

xray:
	go clean -i -x
	set GOOS=$(GOOS_LINUX)
	set GOARCH=$(GOARCH)
	set CGO_ENABLED=$(CGO_ENABLED)
	go build -o yez_xray -trimpath -ldflags "-s -w -buildid=" ./main
	set GOOS=$(GOOS_WINDOWS)
	go build -o yez_xray_windows_amd64.exe -trimpath -ldflags "-s -w -buildid=" ./main