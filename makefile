GOOS_LINUX = linux
GOOS_WINDOWS = windows
GOOS_MAC = darwin
GOARCH_AMD64 = amd64
GOARCH_ARM64 = arm64
CGO_ENABLED = 0

xray:
	set GOOS=$(GOOS_LINUX)
	set GOARCH=$(GOARCH_AMD64)
	set CGO_ENABLED=$(CGO_ENABLED)
	go build -o yez_xray -trimpath -ldflags "-s -w -buildid=" ./main
	set GOOS=$(GOOS_WINDOWS)
	go build -o yez_xray_windows_amd64.exe -trimpath -ldflags "-s -w -buildid=" ./main
	set GOOS=$(GOOS_MAC)
	set GOARCH=$(GOARCH_ARM64)
	go build -o yez_xray_mac_arm64 -trimpath -ldflags "-s -w -buildid=" ./main