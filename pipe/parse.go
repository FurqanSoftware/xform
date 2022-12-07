package pipe

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Parse struct {
	Fields []ParseField
}

func NewParse(ca *yaml.Node) (Step, error) {
	p := Parse{}
	err := ca.Decode(&p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p Parse) Do(v Values) error {
	r := v.Get("@Request").(*http.Request)

	ctype := r.Header.Get("Content-Type")
	mtype, _, err := mime.ParseMediaType(ctype)
	if err != nil {
		return err
	}

	w := Values{}
	switch mtype {
	case "application/json":
		err = json.NewDecoder(r.Body).Decode(&w)
		if err != nil {
			return err
		}

	case "multipart/form-data":
		err = r.ParseMultipartForm(100 * 1024)
		if err != nil {
			return err
		}
		for k := range r.Form {
			w.Set(k, r.Form.Get(k))
		}

	case "application/x-www-form-urlencoded":
		err = r.ParseForm()
		if err != nil {
			return err
		}
		for k := range r.Form {
			w.Set(k, r.Form.Get(k))
		}

	default:
		return errors.New("missing or unknown content type")
	}
	err = p.validate(w)
	if err != nil {
		return ErrBadParse{err}
	}

	for _, f := range p.Fields {
		v.Set(f.As, w[f.Key])
	}

	return nil
}

func (p *Parse) validate(v Values) error {
	fs := map[string]bool{}
	for _, f := range p.Fields {
		fs[f.Key] = true
	}
	for k := range v {
		if !fs[k] {
			return ErrParseExtraField{
				Field: k,
			}
		}
	}

	for _, f := range p.Fields {
		x, ok := v[f.Key]
		if !ok {
			if f.Required {
				return ErrParseMissingField{
					Field: f.Key,
				}
			}
			continue
		}

		switch f.Type {
		case "string":
			_, ok := x.(string)
			if !ok {
				return ErrParseIncorrectType{
					Field: f.Key,
					Want:  f.Type,
					Got:   "?",
				}
			}

		case "email":
			_, ok := x.(string)
			if !ok {
				return ErrParseIncorrectType{
					Field: f.Key,
					Want:  f.Type,
					Got:   "?",
				}
			}
		}
	}

	return nil
}

type ParseField struct {
	Key      string
	As       string
	Type     string
	Required bool
}

type ErrBadParse struct {
	err error
}

func (e ErrBadParse) Error() string {
	return e.err.Error()
}

func (e ErrBadParse) Unwrap() error {
	return e.err
}

type ErrParseExtraField struct {
	Field string
}

func (e ErrParseExtraField) Error() string {
	return fmt.Sprintf("unexpected field %s", e.Field)
}

type ErrParseMissingField struct {
	Field string
}

func (e ErrParseMissingField) Error() string {
	return fmt.Sprintf("missing field %s", e.Field)
}

type ErrParseIncorrectType struct {
	Field string
	Want  string
	Got   string
}

func (e ErrParseIncorrectType) Error() string {
	return fmt.Sprintf("incorrect type in field %s; want %s, got %s", e.Field, e.Want, e.Got)
}
