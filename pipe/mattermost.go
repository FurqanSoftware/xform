package pipe

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Mattermost struct {
	Webhook string
	Text    string
	tplText *template.Template
}

func NewMattermost(ca *yaml.Node) (Step, error) {
	m := Mattermost{}
	err := ca.Decode(&m)
	if err != nil {
		return nil, err
	}
	m.tplText, err = template.New("").Parse(m.Text)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m Mattermost) Do(v Values) error {
	b := bytes.Buffer{}
	err := m.tplText.Execute(&b, v)
	if err != nil {
		return err
	}

	p := struct {
		Text     string `json:"text"`
		Username string `json:"username"`
	}{
		Text:     b.String(),
		Username: "Xform",
	}
	b.Reset()
	err = json.NewEncoder(&b).Encode(p)
	if err != nil {
		return err
	}
	resp, err := http.Post(m.Webhook, "application/json", &b)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
