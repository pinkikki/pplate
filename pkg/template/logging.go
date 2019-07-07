package template

var (
	GlobalGo = `
package logging

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

type Mode int

const (
	Unknown Mode = iota
	Debug
	Verbose
)

func NewMode(mode string) Mode {
	fmt.Print(mode)
	for i, name := range Debug.Names() {
		if mode == name {
			return Mode(i)
		}
	}
	return Unknown
}

func (m Mode) As() string {
	return m.Names()[m]
}

func (m Mode) Names() []string {
	return []string{
		"unknown",
		"debug",
		"verbose",
	}
}

func Setting(mode Mode) {

	config := getConfig(mode)

	logger, err := config.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		return
	}
	zap.ReplaceGlobals(logger)
}

func getConfig(mode Mode) zap.Config {
	switch mode {
	case Debug:
		return DebugConfig()
	case Verbose:
		return VerboseConfig()
	default:
		return DebugConfig()

	}
}
`
	ConfigGo = `
package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	DebugConfig = func() zap.Config {
		cfg := zap.NewDevelopmentConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		return cfg
	}

	VerboseConfig = func() zap.Config {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		return cfg
	}
)
`
)
