package utils

import (
	"fmt"
)

// Debug - consistent debugging
func Debug(msg interface{}) {
	fmt.Printf("\n%v\n", msg)
}
