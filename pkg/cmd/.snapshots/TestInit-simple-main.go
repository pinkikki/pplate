
package main

import (
	"fmt"
	"github.com/pinkikki/test_module/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.NewPplateCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

