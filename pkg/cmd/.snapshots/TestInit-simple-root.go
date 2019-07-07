
package cmd

import (
	"fmt"
	"github.com/pinkikki/test_module/pkg/logging"
	"github.com/spf13/cobra"
)

func NewPplateCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "test_module",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root called")
		},
	}
	var loggingMode string
	getLoggingMode(rootCmd, &loggingMode)

	var commands []Command
	commands = append(commands, &InitCommand{})
	for _, c := range commands {
		cc := c.NewCommand()
		cobra.OnInitialize(func() {
			logging.Setting(logging.NewMode(loggingMode))
			c.OnInitialize()
		})
		rootCmd.AddCommand(cc)
	}

	return rootCmd
}

func getLoggingMode(cmd *cobra.Command, loggingMode *string) {
	cmd.PersistentFlags().StringVarP(loggingMode, "logging", "l", "verbose", "output log level")
}

