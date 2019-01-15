package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/confio/amino-test-suite/common"
	"github.com/confio/amino-test-suite/samples"
)

const perm = 0755

func renderJSON(outFile string, example common.Example) error {
	data, err := example.Codec.MarshalJSONIndent(example.Data, "", "  ")
	if err != nil {
		return errors.Wrap(err, example.Label)
	}
	err = ioutil.WriteFile(outFile, data, perm)
	return errors.Wrap(err, outFile)
}

func renderBinary(outFile string, example common.Example) error {
	return nil
}

func renderCases(baseDir, subDir string, examples []common.Example) error {
	outDir := filepath.Join(baseDir, subDir)
	os.MkdirAll(outDir, perm)

	for _, example := range examples {
		jsonFile := filepath.Join(outDir, fmt.Sprintf("%s.json", example.Label))
		err := renderJSON(jsonFile, example)
		if err != nil {
			return err
		}

		binaryFile := filepath.Join(outDir, fmt.Sprintf("%s.bin", example.Label))
		err = renderBinary(binaryFile, example)
		if err != nil {
			return err
		}

		// TODO: signBytes
	}
	return nil
}

func main() {
	// TODO: make this configurable
	baseDir := filepath.Join(".", "out")

	err := renderCases(baseDir, "samples", samples.GenerateCases())
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}
}
