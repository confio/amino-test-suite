.PHONY: run install deps

run:
	go run ./main.go ./template.go out

install:
	go install .

deps:
	dep ensure