package cmd

import "github.com/spf13/cobra"

type Command interface {
	NewCommand(ctx *Context) *cobra.Command
	OnInitialize()
	Name() string
}
