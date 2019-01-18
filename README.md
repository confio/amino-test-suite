# amino-test-suite

A generic test suite to evaluate compatibility of non-go amino codec implementations

## Read test vectors

We generate and commit these vectors to the repo

## Usage

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
