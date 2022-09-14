package strutil

import (
	"fmt"
)

func Welcome(hoster string) string {
	return fmt.Sprintf("%s Says: Hello!", hoster)
}
