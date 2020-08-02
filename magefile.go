//+build mage

package main

import (
	"github.com/magefile/mage/sh"
	// mage:import
	cmd "github.com/nolte/plumbing/cmd"
)

func Info() error {
	return sh.Run(
		"env")
}

var Aliases = map[string]interface{}{
	"all": cmd.Kind.Recreate,
}
