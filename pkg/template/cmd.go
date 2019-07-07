package template

var (
	MainGo = `
package main

import (
	"fmt"
	"github.com/pinkikki/{{ .ModuleName }}/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.NewPplateCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
`
	RootGo = `
package cmd

import (
	"fmt"
	"github.com/pinkikki/{{ .ModuleName }}/pkg/logging"
	"github.com/spf13/cobra"
)

func NewPplateCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "{{ .ModuleName }}",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root called")
		},
	}
	var loggingMode string
	setLoggingMode(rootCmd, &loggingMode)

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

func setLoggingMode(cmd *cobra.Command, loggingMode *string) {
	cmd.PersistentFlags().StringVarP(loggingMode, "logging", "l", "verbose", "output log level")
}
`
	InitGo = `
package cmd

import (
	"github.com/spf13/cobra"
)

type InitCommand struct {
	*Context
}

func (c *InitCommand) NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "",
		Long:  "",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c.Logger.Info("init start")

			return nil
		},
	}
}

func (c *InitCommand) OnInitialize() {

}
`
	CommandGo = `
package cmd

import "github.com/spf13/cobra"

type Command interface {
	NewCommand() *cobra.Command
	OnInitialize()
}
`
	ContextGo = `
package cmd

import (
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

type Context struct {
	FS     afero.Fs
	Logger *zap.Logger
}
`
)
