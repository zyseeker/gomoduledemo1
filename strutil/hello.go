package strutil

import (
	"fmt"
)

func Welcome() string {
	hoster := "Steven Z"
	return fmt.Sprintf("%s Says: Hello!", hoster)
}
