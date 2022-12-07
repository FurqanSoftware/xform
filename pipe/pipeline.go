package pipe

import (
	"git.furqansoftware.net/hjr265/xform/cfg"
)

type Pipeline struct {
	Steps  []Step
	Values Values
}

func NewPipeline(cp cfg.Pipeline) (*Pipeline, error) {
	p := &Pipeline{
		Steps:  []Step{},
		Values: Values{},
	}
	for _, cs := range cp.Steps {
		s, err := Steps[cs.Type](&cs.Args)
		if err != nil {
			return nil, err
		}
		p.Steps = append(p.Steps, s)
	}
	return p, nil
}

func (p *Pipeline) Run() (err error) {
	for _, s := range p.Steps {
		err = s.Do(p.Values)
		if err != nil {
			return StepError{
				Step: s,
				err:  err,
			}
		}
	}
	return nil
}
