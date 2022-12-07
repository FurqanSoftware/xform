package pipe

import (
	"bytes"
	"text/template"

	"git.furqansoftware.net/hjr265/xform/cfg"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v3"
)

type Email struct {
	Driver      string
	Host        string
	Port        int
	Username    string
	Password    string
	From        string
	To          string
	Subject     string
	BodyText    string `yaml:"bodyText"`
	tplSubject  *template.Template
	tplBodyText *template.Template
}

func NewEmail(ca *yaml.Node) (Step, error) {
	e := Email{}
	err := ca.Decode(&e)
	if err != nil {
		return nil, err
	}
	e.Password = cfg.Interpolate(e.Password)
	e.tplSubject, err = template.New("").Parse(e.Subject)
	if err != nil {
		return nil, err
	}
	e.tplBodyText, err = template.New("").Parse(e.BodyText)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e Email) Do(v Values) error {
	d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	t, err := d.Dial()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetAddressHeader("To", e.To, "")
	b := bytes.Buffer{}
	e.tplSubject.Execute(&b, v)
	m.SetHeader("Subject", b.String())
	b.Reset()
	e.tplBodyText.Execute(&b, v)
	m.SetBody("text/plain", b.String())
	err = gomail.Send(t, m)
	if err != nil {
		return err
	}

	return nil
}
