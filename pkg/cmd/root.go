package cmd

import (
	"github.com/pinkikki/pplate/pkg/logging"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewPplateCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "pplate",
		Run: func(cmd *cobra.Command, args []string) {
			// nop
		},
	}
	var loggingMode string
	setLoggingMode(rootCmd, &loggingMode)

	var commands []Command
	commands = append(commands, &InitCommand{})
	for _, c := range commands {
		ctx := &Context{FS: afero.NewOsFs()}
		cc := c.NewCommand(ctx)
		cobra.OnInitialize(func() {
			logging.Setting(logging.NewMode(loggingMode))
			ctx.Logger = zap.L().Named(c.Name())
			c.OnInitialize()
		})
		rootCmd.AddCommand(cc)
	}

	return rootCmd
}

func setLoggingMode(cmd *cobra.Command, loggingMode *string) {
	cmd.PersistentFlags().StringVarP(loggingMode, "logging", "l", "verbose", "output log level")
}
