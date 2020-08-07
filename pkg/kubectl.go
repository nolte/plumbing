package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func kubeClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homeDir(), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func CreateNamesaceIfNotExists(namespaceName string) (*core.Namespace, error) {
	clientset := kubeClient()

	namespace := v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: namespaceName},
	}
	obj, err := clientset.CoreV1().Namespaces().Create(&namespace)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func WaitPodReady(namespaceName string, matchLabels map[string]string) error {
	// TODO: Switch implementation to go based version

	args := []string{"wait", "--namespace", namespaceName, "--for=condition=ready", "pod", "--timeout=680s"}
	for key, value := range matchLabels {
		args = append(args, fmt.Sprintf("--selector=%s=%s", key, value))
	}

	return sh.Run("kubectl", args...)
}
