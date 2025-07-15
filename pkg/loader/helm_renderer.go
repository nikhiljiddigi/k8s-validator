package loader

import (
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"

	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func RenderHelmChart(chartPath, valuesFile string) ([]unstructured.Unstructured, error) {
	settings := cli.New()
	cfg := new(action.Configuration)
	if err := cfg.Init(settings.RESTClientGetter(), "default", os.Getenv("HELM_DRIVER"), nil); err != nil {
		return nil, err
	}
	client := action.NewInstall(cfg)
	client.DryRun, client.ClientOnly, client.Replace = true, true, true
	client.ReleaseName = "validate"

	vals := map[string]interface{}{}
	if valuesFile != "" {
		b, _ := os.ReadFile(valuesFile)
		yaml.Unmarshal(b, &vals)
	}

	chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}
	rel, err := client.Run(chart, vals)
	if err != nil {
		return nil, err
	}

	var out []unstructured.Unstructured
	docs := strings.Split(rel.Manifest, "---")
	for _, d := range docs {
		var obj map[string]interface{}
		yaml.Unmarshal([]byte(d), &obj)
		if obj != nil {
			out = append(out, unstructured.Unstructured{Object: obj})
		}
	}
	return out, nil
}
