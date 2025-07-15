package loader

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func LoadYAMLFolder(path string) ([]unstructured.Unstructured, error) {
	var out []unstructured.Unstructured
	err := filepath.Walk(path, func(fp string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(fp, ".yaml") && !strings.HasSuffix(fp, ".yml") {
			return nil
		}
		b, e := ioutil.ReadFile(fp)
		if e != nil {
			return e
		}
		docs := strings.Split(string(b), "---")
		for _, d := range docs {
			var obj map[string]interface{}
			yaml.Unmarshal([]byte(d), &obj)
			if obj != nil {
				out = append(out, unstructured.Unstructured{Object: obj})
			}
		}
		return nil
	})
	return out, err
}
