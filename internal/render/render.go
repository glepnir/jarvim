// Package render provides ...
package render

import (
	"os"

	"github.com/glepnir/jarvis/pkg/color"
	"github.com/glepnir/jarvis/pkg/util"
)

var templatesdirectory = "./../../template"
var NeoVimConf = os.Getenv("HOME") + "/.config/nvim"

func GenerateTemplate() {
	if util.Exist(NeoVimConf) {
		os.Rename(NeoVimConf, NeoVimConf+"-bak")
		color.PrintWarn("Back up your vim config to nvim-bak folder")
	}
	util.CopyDir(templatesdirectory, NeoVimConf)
}
