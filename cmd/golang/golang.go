// Package golang for gh tasks
package golang

import (
	"context"
	"log"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Golang Mage Command Namespace.
type Golang mg.Namespace

// Fmtl the go code
func (Golang) Fmtl(ctx context.Context) error {

	result, err := sh.Output("gofmt", "-l", "$(`find . -name '*.go' | grep -v vendor`)")
	check(err)
	log.Print(result)
	return sh.Run("env")
}
