package cfg

import "os"

func Interpolate(v string) string {
	return os.ExpandEnv(v)
}
