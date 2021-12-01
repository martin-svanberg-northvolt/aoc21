package lib

import (
	"path"
	"runtime"
)

func GetFixturePath(name string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("No caller information")
	}
	return path.Join(filename, "../../fixtures", name)
}
