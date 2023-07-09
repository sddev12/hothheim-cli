package main

import (
	"os"

	"hothheim-cli/actions"
	"hothheim-cli/userinterface"
	"hothheim-cli/utils"
)

func main() {

	utils.Banner()

	pass := userinterface.EnterPassword()

	actions.ExecuteServerInfo(pass)

	for {
		answer := userinterface.MainMenu()

		switch {
		case answer == "Start server":
			actions.ExecuteStartServer(pass)

		case answer == "Stop server":
			actions.ExecuteStopServer(pass)

		case answer == "Exit cli":
			os.Exit(0)
		}

	}
}
