package container

import "poker/internal/metadata"

func Rename(containerIdOrName, newName string) error {
	// check container id
	containerId := checkName(containerIdOrName)
	containerPath, err := findPath(containerId)
	if err != nil {
		return err
	}
	metadataFilePath := containerPath + "/metadata.json"

	// read metadata
	meta, err := metadata.ReadMetadata(metadataFilePath)
	if err != nil {
		return err
	}

	// check if the name available
	err = checkNameAvailable(newName)
	if err != nil {
		return err
	}

	oldName := meta.Name
	meta.Name = newName

	// delete old name
	delete(nameToContainer, oldName)

	// update
	err = metadata.WriteMetadata(metadataFilePath, meta)
	if err != nil {
		return err
	}

	nameToContainer[newName] = meta.Id
	return nil
}
