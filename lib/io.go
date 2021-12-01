package lib

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func FileToString(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func GetInput() []string {
	out := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			out = append(out, text)
		} else {
			break
		}

	}
	return out
}
