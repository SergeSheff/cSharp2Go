package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var strSearchPath, strDestinationPath string

	if len(os.Args) > 1 {
		strSearchPath = os.Args[1]
		strDestinationPath = os.Args[2]
	}

	strSearchPath = "C:\\Users\\Serge_Sheff\\Desktop\\121"
	strDestinationPath = "C:\\Users\\Serge_Sheff\\Desktop\\goResult"

	if len(strings.TrimSpace(strSearchPath)) > 0 {
		if len(strings.TrimSpace(strDestinationPath)) > 0 {
			ProcessPath(strSearchPath, nil)
		} else {
			fmt.Println("Destination path cannot be empty")
		}
	} else {
		fmt.Println("Search path cannot be empty")
	}
}
