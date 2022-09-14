package strutil

import (
	"fmt"
)

func Welcome() string {
	hoster := "StevenZ"
	return fmt.Sprintf("%s Says: Hello!", hoster)
}
