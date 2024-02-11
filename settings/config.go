package config

import "fmt"

var (
	// style
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Help() {
	fmt.Println("Help.")
}
