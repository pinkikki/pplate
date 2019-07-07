package cmd_test

import (
	"github.com/pinkikki/pplate/pkg/logging"
	"os"
	"path/filepath"
	"testing"

	"github.com/bradleyjkemp/cupaloy"

	"github.com/pinkikki/pplate/pkg/cmd"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() {
	logging.Setting(logging.NewMode("debug"))
}

func teardown() {
}

func executeSubCommand(command *cobra.Command, args ...string) error {
	root := &cobra.Command{
		Use: "root",
		Run: func(_ *cobra.Command, args []string) {
			// nop
		},
	}
	root.SetArgs(args)
	root.AddCommand(command)
	return root.Execute()
}

func TestInit(t *testing.T) {

	moduleName := "test_module"

	testCases := []struct {
		name  string
		files []string
	}{
		{"simple",
			[]string{".gitignore",
				"cmd/" + moduleName + "/main.go",
				"cmd/main.go",
				"pkg/cmd/root.go",
				"pkg/cmd/init.go",
				"pkg/cmd/command.go",
				"pkg/cmd/context.go",
				"pkg/logging/global.go",
				"pkg/logging/config.go"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			initCommand := &cmd.InitCommand{}
			fs := afero.NewMemMapFs()
			ctx := &cmd.Context{FS: fs, Logger: zap.L().Named(initCommand.Name())}
			command := initCommand.NewCommand(ctx)
			cobra.OnInitialize(func() {
				initCommand.OnInitialize()
			})
			err := executeSubCommand(command, "init", moduleName)
			if err != nil {
				t.Fatalf("failed to execute init command: %v", err)
			}
			dir, err := os.Getwd()
			if err != nil {
				t.Fatalf("failed to get current directory: %v", err)
			}

			files := make(map[string]struct{})
			for _, f := range tc.files {
				files[filepath.Join(dir, moduleName, f)] = struct{}{}
			}

			err = afero.Walk(fs, dir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					t.Fatal(err)
				}
				if info.IsDir() {
					return nil
				}
				if _, ok := files[path]; ok {
					t.Run(info.Name(), func(t *testing.T) {
						data, err := afero.ReadFile(fs, path)
						if err != nil {
							t.Fatalf("failed to read file: %v", err)
						}
						cupaloy.SnapshotT(t, string(data))
					})
				} else {
					t.Fatalf("unexpected file was created: %v", path)
				}

				return nil
			})
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestInitError(t *testing.T) {
	testCases := []struct {
		name string
	}{{"no args"}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			initCommand := &cmd.InitCommand{}
			fs := afero.NewMemMapFs()
			ctx := &cmd.Context{FS: fs, Logger: zap.L().Named(initCommand.Name())}
			command := initCommand.NewCommand(ctx)
			cobra.OnInitialize(func() {
				initCommand.OnInitialize()
			})
			err := executeSubCommand(command, "init")
			if err == nil {
				t.Fatalf("command arguments must be required : %v", err)
			}
		})
	}
}
