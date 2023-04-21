package container

const (
	// --- Path ---

	IMAGE_PATH         = "/var/lib/poker/images"
	CONTAINER_PATH     = "/var/lib/poker/containers"
	METADATA_PREFIX    = CONTAINER_PATH
	CONTAINER_LOG_PATH = "/var/log/poker"
	DAEMON_LOG_PATH

	// --- ID ---
	MAX_CONTAINERID  = 64
	ID_RAND_SOURCE   = "abcdefghijklmnopqrstuvwxyz0123456789"
	NAME_RAND_SOURCE = "abcdefghijklmnopqrstuvwxyz"
)
