package cmd

import (
	"bytes"
	"github.com/pinkikki/pplate/pkg/template"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type InitCommand struct {
	*Context
}

func (c *InitCommand) NewCommand(ctx *Context) *cobra.Command {
	c.Context = ctx
	return &cobra.Command{
		Use:   "init",
		Short: "template",
		Long:  `template`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c.Logger.Info("init start")

			moduleName := args[0]
			wd, err := os.Getwd()

			if err != nil {
				return err
			}
			modulePath := filepath.Join(wd, moduleName)
			c.Logger.Debug("generate project", zap.String("modulePath", modulePath))

			vars := struct {
				ModuleName string
			}{moduleName}

			params := []*param{
				{path: filepath.Join(modulePath, ".gitignore"), template: template.CreateTemplate(".gitignore", template.Gitignore), withFile: true},
				{path: filepath.Join(modulePath, "cmd", moduleName, "main.go"), template: template.CreateTemplate("main.go", template.MainGo), withFile: true},
				{path: filepath.Join(modulePath, "pkg", "cmd", "root.go"), template: template.CreateTemplate("root.go", template.RootGo), withFile: true},
				{path: filepath.Join(modulePath, "pkg", "cmd", "init.go"), template: template.CreateTemplate("init.go", template.InitGo), withFile: true},
				{path: filepath.Join(modulePath, "pkg", "cmd", "command.go"), template: template.CreateTemplate("command.go", template.CommandGo), withFile: true},
				{path: filepath.Join(modulePath, "pkg", "cmd", "context.go"), template: template.CreateTemplate("context.go", template.ContextGo), withFile: true},
				{path: filepath.Join(modulePath, "pkg", "logging", "global.go"), template: template.CreateTemplate("global.go", template.GlobalGo), withFile: true},
				{path: filepath.Join(modulePath, "pkg", "logging", "config.go"), template: template.CreateTemplate("config.go", template.ConfigGo), withFile: true},
			}

			for _, p := range params {
				err := p.generate(c.Context, vars)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func (c *InitCommand) OnInitialize() {
	// nop
}

func (c *InitCommand) Name() string {
	return "InitCommand"
}

func (p *param) generate(ctx *Context, vars interface{}) error {
	dir := filepath.Dir(p.path)
	if ok, err := afero.DirExists(ctx.FS, dir); err != nil {
		return err
	} else if !ok {
		zap.L().Debug("create a directory", zap.String("dir", dir))
		err = ctx.FS.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	if p.withFile {
		buf := new(bytes.Buffer)
		err := p.template.Execute(buf, vars)
		if err != nil {
			return err
		}
		data := buf.Bytes()

		zap.L().Debug("create a new flie", zap.String("path", p.path))
		err = afero.WriteFile(ctx.FS, p.path, data, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

type param struct {
	path     string
	template *template.Template
	withFile bool
}
