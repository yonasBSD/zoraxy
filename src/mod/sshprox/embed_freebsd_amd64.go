//go:build freebsd && amd64
// +build freebsd,amd64

package sshprox

import "embed"

/*
	Bianry embedding for AMD64 builds

	Make sure when compile, gotty binary exists in static.gotty
*/
var (
	//go:embed gotty/gotty_freebsd_amd64
	//go:embed gotty/.gotty
	//go:embed gotty/LICENSE
	gotty embed.FS
)
