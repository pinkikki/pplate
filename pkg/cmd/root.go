package cmd

import (
	"github.com/pinkikki/pplate/pkg/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Logging string
}

func NewPplateCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "pplate",
		Run: func(cmd *cobra.Command, args []string) {
			// nop
		},
	}

	var configPath string
	var envPrefix string
	var loggingMode string
	setConfigPath(rootCmd, &configPath)
	setEnvPrefix(rootCmd, &envPrefix)
	setLoggingMode(rootCmd, &loggingMode)
	viper.BindPFlags(rootCmd.PersistentFlags())

	var commands []Command
	commands = append(commands, &InitCommand{})
	for _, c := range commands {
		cc := c.NewCommand(&Context{})
		rootCmd.AddCommand(cc)
	}
	cobra.OnInitialize(func() {

		viper.SetEnvPrefix(envPrefix)
		viper.AutomaticEnv()

		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			logging.Setting(logging.NewMode(viper.GetString("logging")))
			zap.L().Warn("failed to read config.", zap.Any("err", err))
		}

		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			logging.Setting(logging.NewMode(viper.GetString("logging")))
			zap.L().Fatal("failed to unmarshal config.", zap.Any("err", err))
		}
		logging.Setting(logging.NewMode(config.Logging))
		for _, c := range commands {
			c.OnInitialize()
		}
	})

	return rootCmd
}

func setConfigPath(cmd *cobra.Command, configPath *string) {
	cmd.PersistentFlags().StringVarP(configPath, "config", "c", "config.toml", "pplate settings")
}

func setEnvPrefix(cmd *cobra.Command, envPrefix *string) {
	cmd.PersistentFlags().StringVarP(envPrefix, "envPrefix", "e", "pplate", "environment variable prefix")
}

func setLoggingMode(cmd *cobra.Command, loggingMode *string) {
	cmd.PersistentFlags().StringVarP(loggingMode, "logging", "l", "verbose", "output log level")
}
