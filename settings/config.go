package config

import (
	"fmt"
	"time"
)

var (
	// style
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Green = "\033[32m"
	Reset = "\033[0m"
	Time  = time.Now().Format("2006-01-02 15:04:05")
)

func Help() {
	fmt.Println("Help.")
}
