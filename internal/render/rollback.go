// Package render provides ...
package render

import (
	"os"

	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
)

// RollBack will remove the config folder when
// some errors disappear
func RollBack(err error) {
	color.PrintError(err.Error())
	color.PrintWarn("Rolling back...")
	os.RemoveAll(vim.ConfPath)
	os.Exit(1)
}
