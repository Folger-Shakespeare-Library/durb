package main

import (
	"os"

	"github.com/Folger-Shakespeare-Library/durb/internal/cli"
)

// version is set at build time via ldflags.
var version = "dev"

func main() {
	cli.Version = version
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
