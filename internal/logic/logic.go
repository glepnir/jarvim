// Package logic provides ...
package logic

import (
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/render/dein"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/util"
)

func LogicInit() {
	vim.PluginManage = PluginManage()
	p := new(dein.Dein)
	r := render.NewRender(p)
	vim.Leaderkey = LeaderKey()
	vim.LocalLeaderKey = LocalLeaderKey()
	vim.Colorscheme = Colorscheme()
	util.EnsureFoldersExist(vim.ConfPath, vim.ConfCore, vim.ConfAutoload, vim.ConfModules)
	r.GenerateCore(vim.Leaderkey, vim.LocalLeaderKey)
	r.GenerateGeneral()
	r.GenerateTheme()
	r.GenerateCacheTheme(vim.Colorscheme)
	r.GenerateColorscheme(vim.Colorscheme)
	r.GenerateStatusLine()
}
