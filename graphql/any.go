package graphql

import (
	"io"

	"github.com/bytedance/sonic"
)

func MarshalAny(v any) Marshaler {
	return WriterFunc(func(w io.Writer) {
		err := sonic.ConfigFastest.NewEncoder(w).Encode(v)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalAny(v any) (any, error) {
	return v, nil
}
