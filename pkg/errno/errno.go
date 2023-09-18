package errno

const (
	OK int64 = 0
)

// internal error
const (
	ErrUnknown int64 = iota + 10000
	ErrWriteFileFailed
	ErrReadFileFailed
	ErrJsonMarshalFailed
	ErrJsonUnmarshalFailed
)

// client error
const (
	ErrContainerNotFound int64 = iota + 20000
	ErrContainerNameConflict
	ErrContainerNameTooLong
	ErrRemoveRunningContainer
	ErrRemoveContainerDirFailed
	ErrImageEmpty
	ErrImageNotFound
)

// container error
const ()

// image error
const ()

var ErrnoMap = map[int64]string{
	OK: "OK",

	ErrUnknown:         "unknown error, please check log",
	ErrWriteFileFailed: "write file failed",
	ErrReadFileFailed:  "read file failed",

	ErrContainerNotFound:        "container not found",
	ErrContainerNameConflict:    "container name conflict",
	ErrContainerNameTooLong:     "container name too long, max length is 63",
	ErrRemoveRunningContainer:   "container is still running, please stop it first",
	ErrRemoveContainerDirFailed: "remove container dir failed",
	ErrImageEmpty:               "must specify an image",
	ErrImageNotFound:            "image not found",
}

var ErrnoFormatMap = map[int64]string{
	OK: "OK",

	ErrUnknown:             "unknown error, please check log",
	ErrWriteFileFailed:     "write file %s failed",
	ErrReadFileFailed:      "read file %s failed",
	ErrJsonMarshalFailed:   "json marshal failed",
	ErrJsonUnmarshalFailed: "json unmarshal failed",

	ErrContainerNotFound:        "container %s not found",
	ErrContainerNameConflict:    "container name %s conflict",
	ErrContainerNameTooLong:     "container name %s too long, max length is 63",
	ErrRemoveRunningContainer:   "container %s is still running, please stop it first",
	ErrRemoveContainerDirFailed: "remove container dir %s failed",
	ErrImageEmpty:               "must specify an image",
	ErrImageNotFound:            "image %s not found",
}
