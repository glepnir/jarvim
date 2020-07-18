// Package logic provides ...
package logic

import (
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/util"
)

func LogicInit() {
	r := PluginManage()
	vim.Leaderkey = LeaderKey()
	vim.LocalLeaderKey = LocalLeaderKey()
	vim.Colorscheme = Colorscheme()
	util.EnsureFoldersExist(vim.ConfPath, vim.ConfCore, vim.ConfAutoload, vim.ConfModules, vim.CachePath)
	r.GenerateCore(vim.Leaderkey, vim.LocalLeaderKey)
	r.GenerateGeneral()
	r.GenerateTheme()
	r.GenerateCacheTheme(vim.Colorscheme)
	r.GenerateColorscheme(vim.Colorscheme)
	r.GenerateStatusLine()
}
