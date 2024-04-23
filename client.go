package goclient

import (
	"fmt"
)

func Client() string {
	return fmt.Sprintf("This is the %s talking", "client")
}
