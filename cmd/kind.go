package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Kind mg.Namespace

func (Kind) Delete() error {
	return sh.Run(
		"kind", "delete", "cluster")
}

func (Kind) Recreate(ctx context.Context) {
	mg.CtxDeps(ctx, Kind.Delete)
	mg.CtxDeps(ctx, Kind.Create)
	mg.CtxDeps(ctx, Kind.InstallIngress)
}
func (Kind) InstallIngress() error {
	url := fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/ingress-nginx/ingress-nginx-%s/deploy/static/provider/kind/deploy.yaml", "2.11.1")
	err := sh.Run("kubectl", "apply", "-f", url, "--wait=true")
	check(err)
	time.Sleep(30 * time.Second)
	return sh.Run("kubectl", "wait", "--namespace", "ingress-nginx",
		"--for=condition=ready", "pod",
		"--selector=app.kubernetes.io/component=controller",
		"--timeout=240s")

}
func (Kind) Create() error {
	kindConfig := `
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
    `

	d1 := []byte(kindConfig)
	err := ioutil.WriteFile("/tmp/kindconfig.yaml", d1, 0644)
	defer os.Remove("/tmp/kindconfig.yaml")
	check(err)
	return sh.Run("kind", "create", "cluster", "--config=/tmp/kindconfig.yaml")
}
