.PHONY: all run install deps

all: deps install

run:
	go run ./main.go ./template.go out

install:
	go install .

deps:
	dep ensure -vendor-only