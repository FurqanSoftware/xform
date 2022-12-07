package cfg

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Pipeline struct {
	Steps []Step `json:"steps"`
}

type Step struct {
	Type string    `json:"type"`
	Args yaml.Node `json:"args"`
}

var Registry = map[string]Pipeline{}

func LoadPipelines(pdir string) error {
	names, err := filepath.Glob(filepath.Join(pdir, "*.yaml"))
	if err != nil {
		return err
	}
	for _, name := range names {
		b, err := ioutil.ReadFile(name)
		if err != nil {
			return err
		}
		pipe := Pipeline{}
		err = yaml.Unmarshal(b, &pipe)
		if err != nil {
			return err
		}
		key := strings.TrimSuffix(filepath.Base(name), ".yaml")
		Registry[key] = pipe
	}

	keys := []string{}
	for k := range Registry {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	switch len(keys) {
	case 0:
		log.Println("No pipelines were loaded.")
	case 1:
		log.Printf("Loaded pipeline: %s", keys[0])
	default:
		log.Printf("Loaded %d pipeline(s): %v", len(keys), strings.Join(keys, ", "))
	}

	return nil
}
