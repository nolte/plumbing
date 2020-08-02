package cmd

import (
	"context"
	"log"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// GitHubWorkflow Mage Command Namespace.
type GitHubWorkflow mg.Namespace

func startActBuild(ctx context.Context) error {
	jobName := ctx.Value("jobName").(string)
	return sh.Run("act", "-j", jobName, "-P", "ubuntu-latest=nektos/act-environments-ubuntu:18.04")
}

// StartJob the cluster.
func (GitHubWorkflow) StartJob(ctx context.Context) error {
	log.Printf("Start Github Workflow")
	return startActBuild(ctx)
}
