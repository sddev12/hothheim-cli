package utils

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func Banner() {
	banner := figure.NewColorFigure("Hothheim", "gothic", "green", true)
	banner.Print()
	fmt.Println("")

}
