package json

import (
	"encoding/json"

	"poker/pkg/errno"
)

func Marshal(v any) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, errno.SimpleErr(errno.ErrJsonMarshalFailed)
	}
	return data, nil
}

func Unmarshal(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return errno.SimpleErr(errno.ErrJsonUnmarshalFailed)
	}
	return nil
}
