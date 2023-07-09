package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintServerStatus(serverStatus string) {
	title := color.New(color.FgCyan)
	var result string

	if serverStatus == "Terminated" {
		value := color.New(color.FgHiRed)
		result = value.Sprint("Stopped")
	} else {
		value := color.New(color.FgGreen)
		result = value.Sprint("Running")
	}

	fmt.Printf("%s: %s\n\n", title.Sprint("Server status"), result)
}
