package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) > 1 {
		settings.SearchPath = os.Args[1]
		settings.DestinationPath = os.Args[2]
	}

	settings.SearchPath = "C:\\Users\\Serge_Sheff\\Desktop\\121"
	settings.DestinationPath = "C:\\Users\\Serge_Sheff\\Desktop\\goResult"

	if len(strings.TrimSpace(settings.SearchPath)) > 0 {
		if len(strings.TrimSpace(settings.DestinationPath)) > 0 {
			ProcessPath(settings.SearchPath, nil)
		} else {
			fmt.Println("Destination path cannot be empty")
		}
	} else {
		fmt.Println("Search path cannot be empty")
	}
}
