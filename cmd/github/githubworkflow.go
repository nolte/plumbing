// Package github for gh tasks
package github

import (
	"context"

	"github.com/magefile/mage/mg"
	"github.com/nolte/plumbing/pkg"
)

// GitHubWorkflow Mage Command Namespace.
type GitHubWorkflow mg.Namespace

// StartJob Github Workflow
func (GitHubWorkflow) StartJob(ctx context.Context) error {
	jobName, ok := ctx.Value("jobName").(string)
	if !ok {
		jobName = "build"
	}
	job := pkg.ActJob{
		Name: jobName,
	}
	return job.Execute()
}
