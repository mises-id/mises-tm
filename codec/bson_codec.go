package codec

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"go.mongodb.org/mongo-driver/bson"
)

// BsonCodecMarshaler defines an interface for codecs that utilize Protobuf for both
// binary and JSON encoding.
type BsonCodecMarshaler interface {
	sdkcodec.Codec
	InterfaceRegistry() types.InterfaceRegistry
}

// BsonCodec defines a codec that utilizes Protobuf for both binary and JSON
// encoding.
type BsonCodec struct {
	interfaceRegistry types.InterfaceRegistry
}

var _ sdkcodec.Codec = &BsonCodec{}
var _ BsonCodecMarshaler = &BsonCodec{}

// NewBsonCodec returns a reference to a new BsonCodec
func NewBsonCodec(interfaceRegistry types.InterfaceRegistry) *BsonCodec {
	return &BsonCodec{interfaceRegistry: interfaceRegistry}
}

// Marshal implements BinaryMarshaler.Marshal method.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.MarshalInterface
func (pc *BsonCodec) Marshal(o sdkcodec.ProtoMarshaler) ([]byte, error) {
	return bson.Marshal(o)
}

// MustMarshal implements BinaryMarshaler.MustMarshal method.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.MarshalInterface
func (pc *BsonCodec) MustMarshal(o sdkcodec.ProtoMarshaler) []byte {
	bz, err := pc.Marshal(o)
	if err != nil {
		panic(err)
	}

	return bz
}

// MarshalLengthPrefixed implements BinaryMarshaler.MarshalLengthPrefixed method.
func (pc *BsonCodec) MarshalLengthPrefixed(o sdkcodec.ProtoMarshaler) ([]byte, error) {
	bz, err := pc.Marshal(o)
	if err != nil {
		return nil, err
	}

	var sizeBuf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(sizeBuf[:], uint64(len(bz)))
	return append(sizeBuf[:n], bz...), nil
}

// MustMarshalLengthPrefixed implements BinaryMarshaler.MustMarshalLengthPrefixed method.
func (pc *BsonCodec) MustMarshalLengthPrefixed(o sdkcodec.ProtoMarshaler) []byte {
	bz, err := pc.MarshalLengthPrefixed(o)
	if err != nil {
		panic(err)
	}

	return bz
}

// Unmarshal implements BinaryMarshaler.Unmarshal method.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterface
func (pc *BsonCodec) Unmarshal(bz []byte, ptr sdkcodec.ProtoMarshaler) error {
	err := bson.Unmarshal(bz, ptr)
	if err != nil {
		return err
	}
	err = types.UnpackInterfaces(ptr, pc.interfaceRegistry)
	if err != nil {
		return err
	}
	return nil
}

// MustUnmarshal implements BinaryMarshaler.MustUnmarshal method.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterface
func (pc *BsonCodec) MustUnmarshal(bz []byte, ptr sdkcodec.ProtoMarshaler) {
	if err := pc.Unmarshal(bz, ptr); err != nil {
		panic(err)
	}
}

// UnmarshalLengthPrefixed implements BinaryMarshaler.UnmarshalLengthPrefixed method.
func (pc *BsonCodec) UnmarshalLengthPrefixed(bz []byte, ptr sdkcodec.ProtoMarshaler) error {
	size, n := binary.Uvarint(bz)
	if n < 0 {
		return fmt.Errorf("invalid number of bytes read from length-prefixed encoding: %d", n)
	}

	if size > uint64(len(bz)-n) {
		return fmt.Errorf("not enough bytes to read; want: %v, got: %v", size, len(bz)-n)
	} else if size < uint64(len(bz)-n) {
		return fmt.Errorf("too many bytes to read; want: %v, got: %v", size, len(bz)-n)
	}

	bz = bz[n:]
	return pc.Unmarshal(bz, ptr)
}

// MustUnmarshalLengthPrefixed implements BinaryMarshaler.MustUnmarshalLengthPrefixed method.
func (pc *BsonCodec) MustUnmarshalLengthPrefixed(bz []byte, ptr sdkcodec.ProtoMarshaler) {
	if err := pc.UnmarshalLengthPrefixed(bz, ptr); err != nil {
		panic(err)
	}
}

// MarshalJSON implements JSONCodec.MarshalJSON method,
// it marshals to JSON using proto codec.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.MarshalInterfaceJSON
func (pc *BsonCodec) MarshalJSON(o proto.Message) ([]byte, error) {
	m, ok := o.(sdkcodec.ProtoMarshaler)
	if !ok {
		return nil, fmt.Errorf("cannot protobuf JSON encode unsupported type: %T", o)
	}

	return sdkcodec.ProtoMarshalJSON(m, pc.interfaceRegistry)
}

// MustMarshalJSON implements JSONCodec.MustMarshalJSON method,
// it executes MarshalJSON except it panics upon failure.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.MarshalInterfaceJSON
func (pc *BsonCodec) MustMarshalJSON(o proto.Message) []byte {
	bz, err := pc.MarshalJSON(o)
	if err != nil {
		panic(err)
	}

	return bz
}

// UnmarshalJSON implements JSONCodec.UnmarshalJSON method,
// it unmarshals from JSON using proto codec.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterfaceJSON
func (pc *BsonCodec) UnmarshalJSON(bz []byte, ptr proto.Message) error {
	m, ok := ptr.(sdkcodec.ProtoMarshaler)
	if !ok {
		return fmt.Errorf("cannot protobuf JSON decode unsupported type: %T", ptr)
	}

	unmarshaler := jsonpb.Unmarshaler{AnyResolver: pc.interfaceRegistry}
	err := unmarshaler.Unmarshal(strings.NewReader(string(bz)), m)
	if err != nil {
		return err
	}

	return types.UnpackInterfaces(ptr, pc.interfaceRegistry)
}

// MustUnmarshalJSON implements JSONCodec.MustUnmarshalJSON method,
// it executes UnmarshalJSON except it panics upon failure.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterfaceJSON
func (pc *BsonCodec) MustUnmarshalJSON(bz []byte, ptr proto.Message) {
	if err := pc.UnmarshalJSON(bz, ptr); err != nil {
		panic(err)
	}
}

// MarshalInterface is a convenience function for proto marshalling interfaces. It packs
// the provided value, which must be an interface, in an Any and then marshals it to bytes.
// NOTE: to marshal a concrete type, you should use Marshal instead
func (pc *BsonCodec) MarshalInterface(i proto.Message) ([]byte, error) {
	if err := assertNotNil(i); err != nil {
		return nil, err
	}
	any, err := types.NewAnyWithValue(i)
	if err != nil {
		return nil, err
	}

	return pc.Marshal(any)
}

// UnmarshalInterface is a convenience function for proto unmarshaling interfaces. It
// unmarshals an Any from bz bytes and then unpacks it to the `ptr`, which must
// be a pointer to a non empty interface with registered implementations.
// NOTE: to unmarshal a concrete type, you should use Unmarshal instead
//
// Example:
//    var x MyInterface
//    err := cdc.UnmarshalInterface(bz, &x)
func (pc *BsonCodec) UnmarshalInterface(bz []byte, ptr interface{}) error {
	any := &types.Any{}
	err := pc.Unmarshal(bz, any)
	if err != nil {
		return err
	}

	return pc.UnpackAny(any, ptr)
}

// MarshalInterfaceJSON is a convenience function for proto marshalling interfaces. It
// packs the provided value in an Any and then marshals it to bytes.
// NOTE: to marshal a concrete type, you should use MarshalJSON instead
func (pc *BsonCodec) MarshalInterfaceJSON(x proto.Message) ([]byte, error) {
	any, err := types.NewAnyWithValue(x)
	if err != nil {
		return nil, err
	}
	return pc.MarshalJSON(any)
}

// UnmarshalInterfaceJSON is a convenience function for proto unmarshaling interfaces.
// It unmarshals an Any from bz bytes and then unpacks it to the `iface`, which must
// be a pointer to a non empty interface, implementing proto.Message with registered implementations.
// NOTE: to unmarshal a concrete type, you should use UnmarshalJSON instead
//
// Example:
//    var x MyInterface  // must implement proto.Message
//    err := cdc.UnmarshalInterfaceJSON(&x, bz)
func (pc *BsonCodec) UnmarshalInterfaceJSON(bz []byte, iface interface{}) error {
	any := &types.Any{}
	err := pc.UnmarshalJSON(bz, any)
	if err != nil {
		return err
	}
	return pc.UnpackAny(any, iface)
}

// UnpackAny implements AnyUnpacker.UnpackAny method,
// it unpacks the value in any to the interface pointer passed in as
// iface.
func (pc *BsonCodec) UnpackAny(any *types.Any, iface interface{}) error {
	return pc.interfaceRegistry.UnpackAny(any, iface)
}

// InterfaceRegistry returns InterfaceRegistry
func (pc *BsonCodec) InterfaceRegistry() types.InterfaceRegistry {
	return pc.interfaceRegistry
}

func assertNotNil(i interface{}) error {
	if i == nil {
		return errors.New("can't marshal <nil> value")
	}
	return nil
}
