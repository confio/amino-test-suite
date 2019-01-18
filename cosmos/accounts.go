package cosmos

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/confio/amino-test-suite/common"
)

// Reference...
// // BaseAccount - a base account structure.
// // This can be extended by embedding within in your AppAccount.
// // There are examples of this in: examples/basecoin/types/account.go.
// // However one doesn't have to use BaseAccount as long as your struct
// // implements Account.
// type BaseAccount struct {
// 	Address sdk.AccAddress `json:"address"`
// 	Coins   sdk.Coins      `json:"coins"`
// 	PubKey  crypto.PubKey  `json:"public_key"`
// 	AccountNumber uint64         `json:"account_number"`
// 	Sequence      uint64         `json:"sequence"`
// }

func GenerateBaseAccount() []*common.ExampleData {
	cdc := amino.NewCodec()
	auth.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	addr, err := sdk.AccAddressFromHex("00cafe00deadbeef00cafe00")
	if err != nil {
		panic(err)
	}
	coins := sdk.Coins{sdk.NewInt64Coin("IOV", 87654321), sdk.NewInt64Coin("ATOM", 100200300)}

	secret := []byte("not so secret string - thus deterministic")
	privkey := secp256k1.GenPrivKeySecp256k1(secret)
	pubkey := privkey.PubKey()

	account := auth.BaseAccount{
		Address:       addr,
		Coins:         coins,
		PubKey:        pubkey,
		AccountNumber: 1234,
		Sequence:      77777,
	}

	res, err := common.RenderAll([]common.Example{
		{"base_account", cdc, account},
	})
	if err != nil {
		panic(err)
	}
	return res
}
