package userinterface

import (
	"log"

	"github.com/manifoldco/promptui"
)

func MainMenu() string {
	prompt := promptui.Select{
		Label: "Make your choice wisely...",
		Items: []string{
			"Start server",
			"Stop server",
			"Exit cli",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("ERROR: PromptUI failed to run")
	}

	return result
}
