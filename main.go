package main

import (
	"fmt"
	"time"

	"github.com/avinassh/viper-test/internal/settings"
)

func main() {
	settings, err := settings.ReadSettings()
	if err != nil {
		fmt.Printf("Failed to get settings. Error: %v\n", err)
		return
	}
	settings.SetLastUpdateCheck(time.Now().Unix())
	fmt.Printf("Last update check: %v\n", settings.GetLastUpdateCheck())
}
