package utils

import (
	"fmt"
	"time"
)

func LoadAnimation(done chan bool) {
	for {

		select {
		case <-done:
			return
		default:
			fmt.Println("Loading |")
			time.Sleep(1 * time.Second)
			ClearLastLine()
			fmt.Println("Loading /")
			time.Sleep(1 * time.Second)
			ClearLastLine()
			fmt.Println("Loading -")
			time.Sleep(1 * time.Second)
			ClearLastLine()
			fmt.Println("Loading \\")
			time.Sleep(1 * time.Second)
			ClearLastLine()
			fmt.Println("Loading |")
			time.Sleep(1 * time.Second)
			ClearLastLine()
			fmt.Println("Loading /")
			time.Sleep(1 * time.Second)
			ClearLastLine()
			fmt.Println("Loading -")
			time.Sleep(1 * time.Second)
			ClearLastLine()
			fmt.Println("Loading \\")
			time.Sleep(1 * time.Second)
			ClearLastLine()
		}
	}
}
