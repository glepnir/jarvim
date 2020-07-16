// Package vim provides ...
package vim

import "os"

var (
	ConfPath       = os.ExpandEnv("$HOME/.config/xnvim")
	PluginManage   string
	Leaderkey      string
	LocalLeaderKey string
)
