package pkg

import "github.com/magefile/mage/sh"

type ActJob struct {
	Name string
}

func (j ActJob) Execute() error {
	return sh.Run("act", "-j", j.Name, "-P", "ubuntu-latest=nektos/act-environments-ubuntu:18.04")
}
