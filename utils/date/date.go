package dateutils

import "time"

const (
	// FORMAT time
	FORMAT    = "2006-01-02T15:04:05Z"
	FORMAT_DB = "2006-01-02 15:04:05"
)

// GetNow Commnent
func GetNow() time.Time {
	return time.Now()
}

// GetFormatedNow Commnent
func GetFormatedNow() string {
	return GetNow().Format(FORMAT)
}

// GetDBFormatedNow Commnent
func GetDBFormatedNow() string {
	return GetNow().Format(FORMAT_DB)
}
