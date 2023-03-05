package metadata

import (
	jsoniter "github.com/json-iterator/go"
	"os"
	"poker/internal/types"
)

func encodeMetadata(metadata *types.ContainerMetadata) ([]byte, error) {
	return jsoniter.Marshal(metadata)
}

func decodeMetadata(data []byte) (*types.ContainerMetadata, error) {
	metadata := &types.ContainerMetadata{}
	err := jsoniter.Unmarshal(data, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

// WriteMetadata write metadata
func WriteMetadata(dst string, metadata *types.ContainerMetadata) error {
	// marshal to json
	b, err := encodeMetadata(metadata)
	if err != nil {
		return err
	}

	// write in
	if err := os.WriteFile(dst, b, 0777); err != nil {
		return err
	}

	return nil
}

// ReadMetadata read metadata
func ReadMetadata(src string) (*types.ContainerMetadata, error) {
	b, err := os.ReadFile(src)
	if err != nil {
		return nil, err
	}
	return decodeMetadata(b)
}