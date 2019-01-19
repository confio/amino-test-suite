# amino-test-suite

A generic test suite to evaluate compatibility of non-go amino codec implementations

## Pre-generated test vectors

We generate and commit test vectors to the repo.
You can copy them from [this directory](./out),
or modify the example code and run to generate your own test suite.

## Development / Custom Examples

Checkout out this repository into the proper location and build it:

```shell
go get -d github.com/confio/amino-test-suite
cd ${GOPATH:-~/go}/src/github.com/confio/amino-test-suite
make
```

Run this to generate some js test vectors:

```shell
amino-test-suite ./out
ls ./out
```
