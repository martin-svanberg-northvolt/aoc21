package lib

import (
	"io/ioutil"
	"log"
)

func FileToString(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
