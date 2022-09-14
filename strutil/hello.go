package strutil

import (
	"fmt"
)

func Welcome() string {
	hoster := "Steven"
	return fmt.Sprintf("%s Says: Hello!", hoster)
}
