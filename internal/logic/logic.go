// Package logic provides ...
package logic

import (
	"os"

	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/render/dein"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
	"github.com/glepnir/jarvis/pkg/util"
)

// Ensure our necessary folders exists
func EnsureFolders() {
	if util.Exist(vim.ConfPath) {
		os.Rename(vim.ConfPath, vim.ConfPath+"-bak")
		color.PrintWarn("Back up your vim config to nvim-bak folder")
	}
	// Ensure our folder exists
	os.MkdirAll(vim.ConfPath+"/core", 0700)
	os.MkdirAll(vim.ConfPath+"/autoload", 0700)
	os.MkdirAll(vim.ConfPath+"/modules", 0700)
}
func LogicInit() {
	vim.PluginManage = PluginManage()
	p := new(dein.Dein)
	r := render.NewRender(p)
	vim.Leaderkey = LeaderKey()
	vim.LocalLeaderKey = LocalLeaderKey()
	vim.Colorscheme = Colorscheme()
	EnsureFolders()
	r.GenerateCore(vim.Leaderkey, vim.LocalLeaderKey)
	r.GenerateGeneral()
	r.GenerateTheme()
	r.GenerateCacheTheme(vim.Colorscheme)
	r.GenerateColorscheme(vim.Colorscheme)
	r.GenerateStatusLine()
}
