
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

