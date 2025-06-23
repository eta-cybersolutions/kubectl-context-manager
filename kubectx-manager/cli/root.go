package cli

import (
	"log"
	"os"
	"os/exec"

	"github.com/eta-cybersolutions/kubectl-context-manager/kubectx-manager/internal/kubeconfig"

	"github.com/manifoldco/promptui"
)

func Execute() {
	contexts, current, err := kubeconfig.GetContexts()
	if err != nil {
		log.Fatalf("Error loading contexts: %v", err)
	}

	type ctxItem struct {
		Name     string
		IsActive bool
	}

	var items []ctxItem
	for _, ctx := range contexts {
		items = append(items, ctxItem{
			Name:     ctx,
			IsActive: ctx == current,
		})
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉  {{ .Name | cyan }}",
		Inactive: `   {{ if .IsActive }}{{ .Name | green }} (current){{ else }}{{ .Name }}{{ end }}`,
		Selected: "✔️  {{ .Name }}",
	}

	prompt := promptui.Select{
		Label:     "Select Kubernetes Context",
		Items:     items,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	selected := items[i].Name
	cmd := exec.Command("kubectl", "config", "use-context", selected)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to switch context: %v", err)
	}
}
