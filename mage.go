//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Info() error {
	return sh.Run(
		"env")
}
