package pipe

type Values map[string]interface{}

func (v Values) Get(k string) interface{} {
	return v[k]
}

func (v Values) Set(k string, x interface{}) {
	v[k] = x
}
