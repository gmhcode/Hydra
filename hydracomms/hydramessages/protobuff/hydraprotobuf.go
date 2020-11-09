package hydraproto

import (
	"errors"

	"google.golang.org/protobuf/proto"
	// "github.com/golang/protobuf/proto"
)

func EncodeProto(obj interface{}) ([]byte, error) {
	if v, ok := obj.(*Ship); ok {
		return proto.Marshal(v)
	}
	return nil, errors.New("Proto: Unknown message type")
}

func DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)
	return pb, proto.Unmarshal(buffer, pb)
}
