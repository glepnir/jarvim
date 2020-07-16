// Package logic provides ...
package logic

import (
	"github.com/glepnir/jarvis/internal/render"
)

func LogicInit() {
	// vim.PluginManage = PluginManage()
	// vim.Leaderkey = LeaderKey()
	// vim.LocalLeaderKey = LocalLeaderKey()
	// render.GenerateLeaderkey(vim.Leaderkey, vim.LocalLeaderKey)
	render.GenerateGeneral()
}
