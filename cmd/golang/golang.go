// Package golang for gh tasks
package golang

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	// using golint from path call
	_ "github.com/golangci/golangci-lint/pkg/commands"

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

func findGoFilesInFolder(startDirectory string) ([]string, error) {
	var filearray []string

	err := filepath.Walk(startDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".go") {
				out, err := sh.Output("gofmt", "-l", path)
				if err != nil {
					return err
				}
				if out != "" {
					log.Printf("Some Error at %s", out)
					filearray = append(filearray, out)
					return nil
				}
				return nil
			}
			return nil
		})

	return filearray, err
}

func (Golang) Lint() error {
	return sh.Run("golangci-lint", "run")
}

// Fmt will be autoformat the miss formatted files.
func (Golang) Fmt(ctx context.Context) {
	filearray, err := findGoFilesInFolder(".")
	check(err)

	for _, file := range filearray {
		err := sh.Run("gofmt", "-w", "-s", file)
		check(err)
	}
}

// CheckFmt checking the sources with go gofmt.
func (Golang) CheckFmt(ctx context.Context) error {
	filearray, err := findGoFilesInFolder(".")
	check(err)

	if len(filearray) == 0 {
		return nil
	}

	errorMessage := fmt.Sprintf(
		`Invalid Formatted Files: %v
	use:
		mage golang:fmt
	for formatting, the sources.`,
		filearray,
	)

	return errors.New(errorMessage)
}
