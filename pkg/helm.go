package pkg

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

type HelmRepository struct {
	Name string
	URL  string
}

type HelmChart struct {
	Repository HelmRepository
	Name       string
	Version    string
}

type HelmDeployment struct {
	Chart       HelmChart
	Namespace   string
	ReleaseName string
	ExtraValues map[string]string
}

func helmAction(args ...string) error {
	return sh.Run("helm", args...)
}

func (r HelmRepository) Add() error {
	return helmAction("repo", "add", r.Name, r.URL)
}
func HelmUpdateRepos() error {
	return helmAction("repo", "update")

}
func (r HelmDeployment) Delete() error {
	args := []string{
		"delete",
		r.ReleaseName,
		"-n", r.Namespace,
	}
	return helmAction(args...)
}

func (r HelmDeployment) Deploy() error {
	args := []string{
		"upgrade", "-i",
		r.ReleaseName,
		fmt.Sprintf("%s/%s", r.Chart.Repository.Name, r.Chart.Name),
		"-n", r.Namespace}

	if r.Chart.Version != "" {
		args = append(args, "-v", r.Chart.Version)
	}

	for key, value := range r.ExtraValues {
		args = append(args, "--set", fmt.Sprintf("%s=%s", key, value))
	}

	return helmAction(args...)
}

func ApplyHelmChart(deployment HelmDeployment, matchLabels map[string]string) {
	CreateNamesaceIfNotExists(deployment.Namespace)

	err := deployment.Chart.Repository.Add()
	if err != nil {
		panic(err.Error())
	}

	err = HelmUpdateRepos()
	if err != nil {
		panic(err.Error())
	}

	deployment.Deploy()
	WaitPodReady(deployment.Namespace, matchLabels)

}
