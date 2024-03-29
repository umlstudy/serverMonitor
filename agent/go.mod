module github.com/umlstudy/serverMonitor/agent

go 1.12

replace github.com/umlstudy/serverMonitor/common => ../common

require (
	github.com/go-errors/errors v1.0.1
	github.com/k0kubun/go-ansi v0.0.0-20180517002512-3bf9e2903213
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
	github.com/pkg/errors v0.8.1
	github.com/shirou/gopsutil v2.18.12+incompatible
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4
	github.com/umlstudy/serverMonitor/common v0.0.0-00010101000000-000000000000
)
