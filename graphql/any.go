package graphql

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

func MarshalAny(v interface{}) Marshaler {
	return WriterFunc(func(w io.Writer) {
		jsonI := jsoniter.ConfigCompatibleWithStandardLibrary

		err := jsonI.NewEncoder(w).Encode(v)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalAny(v interface{}) (interface{}, error) {
	return v, nil
}
