package kubeconfig

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type KubeConfig struct {
	Contexts []struct {
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
}

func GetContexts() ([]string, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, "", err
	}

	configPath := filepath.Join(home, ".kube", "config")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, "", err
	}

	var kc KubeConfig
	if err := yaml.Unmarshal(data, &kc); err != nil {
		return nil, "", err
	}

	var names []string
	for _, ctx := range kc.Contexts {
		names = append(names, ctx.Name)
	}

	return names, kc.CurrentContext, nil
}
