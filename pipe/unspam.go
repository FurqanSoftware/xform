package pipe

import (
	"regexp"

	"gopkg.in/yaml.v3"
)

type Unspam struct {
	Field     string
	ShortURLs int `yaml:"shortURLs"`
}

func NewUnspam(ca *yaml.Node) (Step, error) {
	a := Unspam{}
	err := ca.Decode(&a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a Unspam) Do(v Values) error {
	x, ok := v[a.Field].(string)
	if !ok {
		return nil
	}

	if a.ShortURLs != -1 {
		re := regexp.MustCompile(`bit\.ly|tinyurl\.com|goolnk.com`)
		if re.MatchString(x) {
			v.Set("_Unspam/Triggered", true)
		}
	}

	return nil
}

type ErrUnspamTriggered struct{}

func (e ErrUnspamTriggered) Error() string {
	return "unspam triggered"
}
