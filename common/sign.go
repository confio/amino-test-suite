package common

import "github.com/cosmos/cosmos-sdk/x/auth"

func CalcSignBytes(o interface{}) []byte {
	// TODO: don't hardcode these
	chainID := "test-chain"
	accnum := uint64(1234)
	seq := uint64(42)

	var signBytes []byte
	if tx, ok := o.(auth.StdTx); ok {
		signBytes = auth.StdSignBytes(chainID, accnum, seq, tx.Fee, tx.Msgs, tx.Memo)
	}
	return signBytes
}
