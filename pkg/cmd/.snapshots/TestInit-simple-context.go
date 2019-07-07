
package cmd

import (
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

type Context struct {
	FS     afero.Fs
	Logger *zap.Logger
}

