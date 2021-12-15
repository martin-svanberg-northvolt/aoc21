new:
	mkdir -p days/$(day)/a
	mkdir -p days/$(day)/b
	echo "package main\n\nfunc main() {\n\n}" > days/$(day)/a/main.go
	echo "package main\n\nfunc main() {\n\n}" > days/$(day)/b/main.go
	curl --cookie "${AOC_SESSION_COOKIE}" https://adventofcode.com/2021/day/$(day)/input > days/$(day)/input

run:
	@go run days/$(day)/$(part)/main.go < days/$(day)/input

try:
	@go run days/$(day)/$(part)/main.go < days/$(day)/test-input

test:
	@go run ./tester

.PHONY: new run try test
