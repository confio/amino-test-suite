package common

import (
	"encoding/hex"

	"github.com/pkg/errors"

	amino "github.com/tendermint/go-amino"
)

// Example is returned from any package to express something that needs to be rendered
type Example struct {
	Label string
	Codec *amino.Codec
	Data  interface{}
}

// Render encodes the data with the codec
func (e Example) Render() (*ExampleData, error) {
	json, err := e.Codec.MarshalJSONIndent(e.Data, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, e.Label)
	}
	bin, err := e.Codec.MarshalBinaryLengthPrefixed(e.Data)
	if err != nil {
		return nil, errors.Wrap(err, e.Label)
	}
	binHex := hex.EncodeToString(bin)
	signBytes := CalcSignBytes(e.Data)
	return &ExampleData{
		Label:     e.Label,
		JSON:      string(json),
		BinaryHex: binHex,
		SignBytes: string(signBytes),
	}, nil
}

// ExampleData is rendered from Example
type ExampleData struct {
	Label     string
	JSON      string
	BinaryHex string
	SignBytes string
}

func RenderAll(examples []Example) ([]*ExampleData, error) {
	var err error
	out := make([]*ExampleData, len(examples))

	for i, ex := range examples {
		out[i], err = ex.Render()
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}
