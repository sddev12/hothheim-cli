package userinterface

import (
	"log"

	"github.com/manifoldco/promptui"
)

func EnterPassword() string {
	prompt := promptui.Prompt{
		Label: "Hello friend. The password is...",
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal("Unable to complete password entry prompt")
	}

	return result
}
