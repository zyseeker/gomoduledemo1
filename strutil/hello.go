package strutil

import (
	"fmt"
)

func Welcome() string {
	hoster := "Steven Z1"
	return fmt.Sprintf("%s Says: Hello!", hoster)
}
