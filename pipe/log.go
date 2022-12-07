package pipe

import (
	"bytes"
	"log"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Log struct {
	Text    string
	tplText *template.Template
}

func NewLog(ca *yaml.Node) (Step, error) {
	h := Log{}
	err := ca.Decode(&h)
	if err != nil {
		return nil, err
	}
	h.tplText, err = template.New("").Parse(h.Text)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h Log) Do(v Values) error {
	b := bytes.Buffer{}
	err := h.tplText.Execute(&b, v)
	if err != nil {
		return err
	}
	log.Println(b.String())
	return nil
}
