package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/confio/amino-test-suite/common"
	"github.com/confio/amino-test-suite/cosmos"
	"github.com/confio/amino-test-suite/samples"
	"github.com/pkg/errors"
)

const permDir = 0755
const permFile = 0644

func renderCases(tmpl *template.Template, baseDir, label string, examples []*common.ExampleData) error {
	outFile := filepath.Join(baseDir, label+".js")
	f, err := os.Create(outFile)
	if err != nil {
		return errors.Wrap(err, "Create file")
	}
	defer f.Close()

	err = tmpl.Execute(f, examples[0])
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: amino-test-suite <outDir>")
		return
	}

	tmpl, err := GetJSTemplate()
	if err != nil {
		fmt.Printf("ERROR Template: %+v\n", err)
		return
	}

	// prepare outDir
	baseDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR parse dir: %+v\n", err)
		return
	}
	os.MkdirAll(baseDir, permDir)
	fmt.Printf("Writing test cases to %s\n", baseDir)

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
		err := renderCases(tmpl, baseDir, ex.label, ex.examples)
		if err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			return
		}
	}
	fmt.Printf("Wrote examples to %s\n", baseDir)
}
