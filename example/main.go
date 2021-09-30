package main

import (
	"fmt"

	"github.com/nikx-gupta/buildagent"
)

func main() {
	err := buildagent.Run()
	if err != nil {
		fmt.Printf("Error Starting build agent: %s \n", err.Error())
	}
}
