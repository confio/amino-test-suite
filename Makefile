.PHONY: run install deps

run:
	go run ./main.go

install:
	go install .

deps:
	dep ensure