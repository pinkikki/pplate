package main

import (
	"fmt"
	"github.com/pinkikki/pplate/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.NewPplateCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
