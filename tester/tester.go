package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
				fmt.Printf("Running %s... ", dir)
				cmd := exec.Command("go", "run", "./"+dir)
				stdin, err := cmd.StdinPipe()
				if err != nil {
					log.Fatal(err)
				}
				io.WriteString(stdin, lib.FileToString(filepath.Join(dir, "../input")))
				stdin.Close()
				stdout, err := cmd.Output()
				if err != nil {
					log.Fatal(err)
				}
				answer := strings.Trim(lib.FileToString(path), "\n")
				got := strings.TrimSpace(string(stdout))
				if answer != got {
					fmt.Println("FAIL")
					fmt.Printf("  Expected \"%s\", got \"%s\"\n", answer, got)
					success = false
				} else {
					fmt.Println("SUCCESS")
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
