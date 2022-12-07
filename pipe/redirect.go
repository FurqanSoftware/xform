package pipe

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type Redirect struct {
	URL string
}

func NewRedirect(ca *yaml.Node) (Step, error) {
	h := Redirect{}
	err := ca.Decode(&h)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h Redirect) Do(v Values) error {
	r := v["@Request"].(*http.Request)
	wr := v["@ResponseWriter"].(http.ResponseWriter)
	http.Redirect(wr, r, h.URL, http.StatusSeeOther)
	return nil
}
