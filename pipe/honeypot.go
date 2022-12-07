package pipe

import "gopkg.in/yaml.v3"

type Honeypot struct {
	Field string
	// Penalty  string
	// Duration time.Duration
}

func NewHoneypot(ca *yaml.Node) (Step, error) {
	h := Honeypot{}
	err := ca.Decode(&h)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h Honeypot) Do(v Values) error {
	x, ok := v[h.Field].(string)
	if !ok {
		return nil
	}
	if x == "" {
		return nil
	}
	return ErrHoneypotTriggered{}
}

type ErrHoneypotTriggered struct{}

func (e ErrHoneypotTriggered) Error() string {
	return "honeypot triggered"
}
