package strutil

import (
	"fmt"
)

func Welcome() string {
	hoster := "StevenZY"
	return fmt.Sprintf("%s Says: Hello!", hoster)
}
