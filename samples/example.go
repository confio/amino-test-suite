package samples

import (
	"time"

	amino "github.com/tendermint/go-amino"

	"github.com/confio/amino-test-suite/common"
)

type SimpleStruct struct {
	Number int64
	String string
	Bytes  []byte
	Time   time.Time
}

type BigStruct struct {
	Sub   []SimpleStruct
	Count uint64
	Info  string
}

func GenerateCases() []common.Example {
	cdc := amino.NewCodec()

	t, _ := time.Parse(time.RFC3339, "2008-11-29T01:02:03Z+0100")
	t2, _ := time.Parse(time.RFC3339, "2019-02-14T14:45:67Z")

	simple := SimpleStruct{
		Number: 29,
		String: "foobar",
		Bytes:  []byte{0xde, 0xad, 0xbe, 0xef},
		Time:   t,
	}

	simple2 := SimpleStruct{
		Number: -81,
		String: "",
		Bytes:  []byte{0x00, 0x01},
		Time:   t2,
	}

	big := BigStruct{
		Sub:   []SimpleStruct{simple, simple2},
		Count: 1234567,
		Info:  "longer and longer and longer",
	}

	return []common.Example{
		{"simple", cdc, simple},
		{"simple2", cdc, simple2},
		{"big", cdc, big},
	}
}
