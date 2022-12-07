package pipe

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Dump struct{}

func NewDump(ca *yaml.Node) (Step, error) {
	h := Dump{}
	err := ca.Decode(&h)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h Dump) Do(v Values) error {
	log.Println(v)
	return nil
}
