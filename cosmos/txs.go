package cosmos

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/confio/amino-test-suite/common"
)

func GenerateTxs() []*common.ExampleData {
	cdc := amino.NewCodec()
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	addr1 := mustAddr("00cafe00deadbeef00cafe00")
	addr2 := mustAddr("01234567890123456789")
	coin1 := sdk.NewInt64Coin("FOO", 87654321)
	coin2 := sdk.NewInt64Coin("BAR", 9876543210) // this is a bit more than 2**32

	// simple send
	msg := bank.MsgSend{
		Inputs: []bank.Input{
			{Address: addr1, Coins: sdk.Coins{coin2}},
		},
		Outputs: []bank.Output{
			{Address: addr2, Coins: sdk.Coins{coin2}},
		},
	}
	gas := uint64(123456)
	fee := auth.NewStdFee(gas, coin1)
	memo := ""
	msgs := []sdk.Msg{msg}
	unsigned := auth.NewStdTx(msgs, fee, nil, memo)

	// sign these bytes
	signBytes := common.CalcSignBytes(unsigned)
	secret := []byte("fixed key to sign transactions")
	privkey := secp256k1.GenPrivKeySecp256k1(secret)
	pubkey := privkey.PubKey()
	sigBytes, err := privkey.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	// make a std signature
	stdSig := auth.StdSignature{
		PubKey:    pubkey,
		Signature: sigBytes,
	}

	// build a full, signed transaction
	signed := auth.NewStdTx(msgs, fee, []auth.StdSignature{stdSig}, memo)

	res, err := common.RenderAll([]common.Example{
		{"unsigned_send_tx", cdc, unsigned},
		{"signed_send_tx", cdc, signed},
	})
	if err != nil {
		panic(err)
	}
	return res
}
