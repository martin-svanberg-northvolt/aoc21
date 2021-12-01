package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"log"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	success := true
	err := filepath.Walk("./days",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			base := filepath.Base(path)
			if base == "answer" {
				dir := filepath.Dir(path)
				cmd := exec.Command("go", "run", "./"+dir)
				stdin, err := cmd.StdinPipe()
				if err != nil {
					log.Fatal(err)
				}
				io.WriteString(stdin, lib.FileToString("./"+dir+"/input"))
				stdin.Close()
				stdout, err := cmd.Output()
				if err != nil {
					log.Fatal(err)
				}
				answer := lib.FileToString(path)
				got := string(stdout)
				if answer != string(stdout) {
					fmt.Printf("%s: Expected %s, got %s\n", dir, answer, got)
					success = false
				}
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	if !success {
		os.Exit(1)
	}
}
