package metadata

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func encode(metadata *Container) ([]byte, error) {
	return json.Marshal(metadata)
}

func decode(data []byte) (*Container, error) {
	metadata := &Container{}
	err := json.Unmarshal(data, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

// WriteMetadata write metadata
func WriteMetadata(dst string, metadata *Container) error {
	// marshal to json
	b, err := encode(metadata)
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
func ReadMetadata(src string) (*Container, error) {
	b, err := os.ReadFile(src)
	if err != nil {
		return nil, err
	}
	return decode(b)
}

// ReadAll read all metadata
func ReadAll(src string) ([]*Container, error) {
	entry, err := os.ReadDir(src)
	if err != nil {
		return nil, err
	}
	metas := make([]*Container, 0, len(entry))
	for _, v := range entry {
		metadataFilePath := filepath.Join(src, v.Name(), "metadata.json")
		meta, err := ReadMetadata(metadataFilePath)
		if err != nil {
			continue
		}
		metas = append(metas, meta)
	}
	return metas, nil
}
