package strutil

import (
	"github.com/google/uuid"
)

func GetUidString() string {
	return uuid.New().String()
}
