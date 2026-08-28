package main

import (
	"os"
	"runtime/debug"

	"github.com/Folger-Shakespeare-Library/durb/internal/cli"
)

// version is set at build time via ldflags.
var version = "dev"

func main() {
	if version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			version = info.Main.Version
		}
	}
	cli.Version = version
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
