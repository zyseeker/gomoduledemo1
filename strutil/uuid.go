package strutil

import (
	"github.com/google/uuid"
)

func GetUidString(prefix string) string {
	return prefix + uuid.New().String()
}
