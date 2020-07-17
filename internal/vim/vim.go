// Package vim provides ...
package vim

import "os"

var (
	ConfPath       = os.ExpandEnv("$HOME/.config/xnvim")
	CachePath      = os.ExpandEnv("$HOME/.cache/vim")
	PluginManage   string
	Leaderkey      string
	LocalLeaderKey string
	Colorscheme    []string
)
