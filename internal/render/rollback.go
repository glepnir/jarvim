// Package render provides ...
package render

import (
	"os"

	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
)

func RollBack(err error) {
	color.PrintError(err.Error())
	color.PrintWarn("Rolling back...")
	os.RemoveAll(vim.ConfPath)
	os.Exit(1)
}
