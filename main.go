package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/confio/amino-test-suite/common"
	"github.com/confio/amino-test-suite/cosmos"
	"github.com/confio/amino-test-suite/samples"
)

const perm = 0755

func renderCases(baseDir, subDir string, examples []*common.ExampleData) error {
	outDir := filepath.Join(baseDir, subDir)
	os.MkdirAll(outDir, perm)

	for _, example := range examples {
		jsonFile := filepath.Join(outDir, fmt.Sprintf("%s.json", example.Label))
		err := ioutil.WriteFile(jsonFile, []byte(example.JSON), perm)
		if err != nil {
			return errors.Wrap(err, jsonFile)
		}

		binaryFile := filepath.Join(outDir, fmt.Sprintf("%s.bin", example.Label))
		err = ioutil.WriteFile(binaryFile, []byte(example.BinaryHex), perm)
		if err != nil {
			return errors.Wrap(err, binaryFile)
		}

		// TODO: signBytes
	}
	return nil
}

func main() {
	// TODO: make this configurable
	baseDir := filepath.Join(".", "out")

	examples := []struct {
		label    string
		examples []*common.ExampleData
	}{
		{"samples", samples.GenerateCases()},
		{"cosmos_base_account", cosmos.GenerateBaseAccount()},
		{"cosmos_messages", cosmos.GenerateMessages()},
		{"cosmos_txs", cosmos.GenerateTxs()},
	}

	for _, ex := range examples {
		err := renderCases(baseDir, ex.label, ex.examples)
		if err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			return
		}
	}
	fmt.Printf("Wrote examples to %s\n", baseDir)
}
