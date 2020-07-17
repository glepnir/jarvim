// Package logic provides ...
package logic

import (
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/vim"
)

func LogicInit() {
	vim.PluginManage = PluginManage()
	vim.Leaderkey = LeaderKey()
	vim.LocalLeaderKey = LocalLeaderKey()
	render.EnsureFolders()
	render.GenerateCore(vim.Leaderkey, vim.LocalLeaderKey)
	render.GenerateGeneral()
	render.GenerateTheme()
}
