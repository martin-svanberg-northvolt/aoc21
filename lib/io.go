package lib

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		out = append(out, strings.TrimSpace(line))
		if err == io.EOF {
			break
		}
	}
	return out
}
