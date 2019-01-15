package common

import amino "github.com/tendermint/go-amino"

// Example is returned from any package to express something that needs to be rendered
type Example struct {
	Label string
	Codec *amino.Codec
	Data  interface{}
}
