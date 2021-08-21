package codec_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/mises-id/mises-tm/codec"
)

func createTestInterfaceRegistry() types.InterfaceRegistry {
	interfaceRegistry := types.NewInterfaceRegistry()
	interfaceRegistry.RegisterInterface("testdata.Animal",
		(*testdata.Animal)(nil),
		&testdata.Dog{},
		&testdata.Cat{},
	)

	return interfaceRegistry
}

func TestBsonMarsharlInterface(t *testing.T) {
	cdc := codec.NewBsonCodec(createTestInterfaceRegistry())
	m := interfaceMarshaler{cdc.MarshalInterface, cdc.UnmarshalInterface}
	testInterfaceMarshaling(require.New(t), m, false)
	m = interfaceMarshaler{cdc.MarshalInterfaceJSON, cdc.UnmarshalInterfaceJSON}
	testInterfaceMarshaling(require.New(t), m, false)
}

func TestBsonCodec(t *testing.T) {
	cdc := codec.NewBsonCodec(createTestInterfaceRegistry())
	testMarshaling(t, cdc)
}

type lyingBsonMarshaler struct {
	sdkcodec.ProtoMarshaler
	falseSize int
}

func (lpm *lyingBsonMarshaler) Size() int {
	return lpm.falseSize
}

func TestBsonoCodecUnmarshalLengthPrefixedChecks(t *testing.T) {
	cdc := codec.NewBsonCodec(createTestInterfaceRegistry())
	cat := testdata.Cat{Lives: 9, Moniker: "glowing"}
	// truth := &lyingBsonMarshaler{
	// 	ProtoMarshaler: &cat,
	// 	falseSize:      cat.Size(),
	// }
	//realSize := len(cdc.MustMarshal(truth))

	falseSizes := []int{
		100,
		5,
	}

	for _, falseSize := range falseSizes {
		falseSize := falseSize

		t.Run(fmt.Sprintf("ByMarshaling falseSize=%d", falseSize), func(t *testing.T) {
			lpm := &lyingBsonMarshaler{
				ProtoMarshaler: &cat,
				falseSize:      falseSize,
			}
			var serialized []byte
			require.NotPanics(t, func() { serialized = cdc.MustMarshalLengthPrefixed(lpm) })

			recv := new(testdata.Cat)
			gotErr := cdc.UnmarshalLengthPrefixed(serialized, recv)
			var wantErr error
			wantErr = nil //simply ignor Size()
			// if falseSize > realSize {
			// 	wantErr = fmt.Errorf("not enough bytes to read; want: %d, got: %d", falseSize, realSize)
			// } else {
			// 	wantErr = fmt.Errorf("too many bytes to read; want: %d, got: %d", falseSize, realSize)
			// }
			require.Equal(t, wantErr, gotErr)
		})
	}

	t.Run("Crafted bad uvarint size", func(t *testing.T) {
		crafted := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}
		recv := new(testdata.Cat)
		gotErr := cdc.UnmarshalLengthPrefixed(crafted, recv)
		require.Equal(t, gotErr, errors.New("invalid number of bytes read from length-prefixed encoding: -10"))

		require.Panics(t, func() { cdc.MustUnmarshalLengthPrefixed(crafted, recv) })
	})
}

func mustAny(msg proto.Message) *types.Any {
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		panic(err)
	}
	return any
}

func BenchmarkBsonCodecMarshalLengthPrefixed(b *testing.B) {
	var pCdc = codec.NewBsonCodec(types.NewInterfaceRegistry())
	var msg = &testdata.HasAnimal{
		X: 1000,
		Animal: mustAny(&testdata.HasAnimal{
			X: 2000,
			Animal: mustAny(&testdata.HasAnimal{
				X: 3000,
				Animal: mustAny(&testdata.HasAnimal{
					X: 4000,
					Animal: mustAny(&testdata.HasAnimal{
						X: 5000,
						Animal: mustAny(&testdata.Cat{
							Moniker: "Garfield",
							Lives:   6,
						}),
					}),
				}),
			}),
		}),
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		blob, err := pCdc.MarshalLengthPrefixed(msg)
		if err != nil {
			b.Fatal(err)
		}
		b.SetBytes(int64(len(blob)))
	}
}
