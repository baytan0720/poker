package container

import "poker/internal/metadata"

func Rename(containerIdOrName, newName string) error {
	// check container id
	_, _, metadataPath, meta, err := checkContainer(containerIdOrName)
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

	// update
	err = metadata.WriteMetadata(metadataPath, meta)
	if err != nil {
		return err
	}

	// delete old name
	delete(nameToContainer, oldName)
	nameToContainer[newName] = meta.Id

	return nil
}
