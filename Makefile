.PHONY: run install deps

run:
	go run ./main.go out

install:
	go install .

deps:
	dep ensure