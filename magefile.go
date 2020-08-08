//+build mage

package main

import (
	"context"

	"github.com/magefile/mage/mg"

	cmd "github.com/nolte/plumbing/cmd/github"

	// mage:import
	_ "github.com/nolte/plumbing/cmd/golang"
)

// GH will be run the GitHub Actions Locally.
func GH(ctx context.Context) {
	GHLint(ctx)
	GHBuild(ctx)
	GHAcc(ctx)
}
func GHLint(ctx context.Context) {
	ctx = context.WithValue(ctx, "jobName", "lint")
	mg.CtxDeps(ctx, cmd.GitHubWorkflow.StartJob)
}

func GHBuild(ctx context.Context) {
	ctx = context.WithValue(ctx, "jobName", "build")
	mg.CtxDeps(ctx, cmd.GitHubWorkflow.StartJob)
}
func GHAcc(ctx context.Context) {
	ctx = context.WithValue(ctx, "jobName", "acc")
	mg.CtxDeps(ctx, cmd.GitHubWorkflow.StartJob)
}
