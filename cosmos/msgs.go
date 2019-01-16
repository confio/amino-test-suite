package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/gov"
	amino "github.com/tendermint/go-amino"

	"github.com/confio/amino-test-suite/common"
)

// References

// type MsgSend struct {
// 	Inputs  []Input  `json:"inputs"`
// 	Outputs []Output `json:"outputs"`
// }
// type Input struct {
// 	Address sdk.AccAddress `json:"address"`
// 	Coins   sdk.Coins      `json:"coins"`
// }
// type Output struct {
// 	Address sdk.AccAddress `json:"address"`
// 	Coins   sdk.Coins      `json:"coins"`
// }

// type MsgSubmitProposal struct {
// 	Title          string         `json:"title"`           //  Title of the proposal
// 	Description    string         `json:"description"`     //  Description of the proposal
// 	ProposalType   ProposalKind   `json:"proposal_type"`   //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
// 	Proposer       sdk.AccAddress `json:"proposer"`        //  Address of the proposer
// 	InitialDeposit sdk.Coins      `json:"initial_deposit"` //  Initial deposit paid by sender. Must be strictly positive.
// }
// type ProposalKind byte
// const (
// 	ProposalTypeNil             ProposalKind = 0x00
// 	ProposalTypeText            ProposalKind = 0x01
// 	ProposalTypeParameterChange ProposalKind = 0x02
// 	ProposalTypeSoftwareUpgrade ProposalKind = 0x03
// )

func mustAddr(hex string) sdk.AccAddress {
	addr, err := sdk.AccAddressFromHex(hex)
	if err != nil {
		panic(err)
	}
	return addr
}

func GenerateMessages() []common.Example {
	cdc := amino.NewCodec()
	bank.RegisterCodec(cdc)

	addr1 := mustAddr("00cafe00deadbeef00cafe00")
	addr2 := mustAddr("01234567890123456789")
	addr3 := mustAddr("feedf00d0000be2bee77")
	coin1 := sdk.NewInt64Coin("FOO", 87654321)
	coin2 := sdk.NewInt64Coin("BAR", 9876543210)          // this is a bit more than 2**32
	coin3 := sdk.NewInt64Coin("FOO", 1234567890123456789) // this is around 2**60, less than int64, but more than js precision (2**53)

	// simple proposal
	proposal := gov.MsgSubmitProposal{
		Title:          "My first proposal",
		Description:    "This is a first test to see if anyone votes for me",
		ProposalType:   gov.ProposalTypeText,
		Proposer:       addr1,
		InitialDeposit: sdk.Coins{coin1, coin2},
	}

	// proposal missing fields
	half_proposal := gov.MsgSubmitProposal{
		Title:          "",
		Description:    "foo",
		ProposalType:   gov.ProposalTypeSoftwareUpgrade,
		Proposer:       addr1,
		InitialDeposit: nil,
	}

	// simple send
	msg := bank.MsgSend{
		Inputs: []bank.Input{
			{Address: addr1, Coins: sdk.Coins{coin1}},
		},
		Outputs: []bank.Output{
			{Address: addr2, Coins: sdk.Coins{coin1}},
		},
	}

	// swap
	msg2 := bank.MsgSend{
		Inputs: []bank.Input{
			{Address: addr2, Coins: sdk.Coins{coin1}},
			{Address: addr3, Coins: sdk.Coins{coin2}},
		},
		Outputs: []bank.Output{
			{Address: addr2, Coins: sdk.Coins{coin2}},
			{Address: addr3, Coins: sdk.Coins{coin1}},
		},
	}

	// big send
	msg3 := bank.MsgSend{
		Inputs: []bank.Input{
			{Address: addr1, Coins: sdk.Coins{coin2, coin3}},
		},
		Outputs: []bank.Output{
			{Address: addr2, Coins: sdk.Coins{coin2}},
			{Address: addr3, Coins: sdk.Coins{coin3}},
		},
	}

	return []common.Example{
		{"proposal", cdc, proposal},
		{"half_proposal", cdc, half_proposal},
		{"send_msg_small", cdc, msg},
		{"send_msg_swap", cdc, msg2},
		{"send_msg_big", cdc, msg3},
	}
}
