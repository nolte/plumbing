package pkg

import (
	"fmt"
	"log"

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
		args = append(args, "--version", r.Chart.Version)
	}

	for key, value := range r.ExtraValues {
		args = append(args, "--set", fmt.Sprintf("%s=%s", key, value))
	}

	return helmAction(args...)
}

func ApplyHelmChart(deployment HelmDeployment, matchLabels map[string]string) {
	_, err := CreateNamesaceIfNotExists(deployment.Namespace)
	if err != nil {
		log.Printf("Namespace allways Exists")
	}

	err = deployment.Chart.Repository.Add()
	CheckError(err)

	err = HelmUpdateRepos()
	CheckError(err)

	err = deployment.Deploy()
	CheckError(err)
	err = WaitPodReady(deployment.Namespace, matchLabels)
	CheckError(err)
}
