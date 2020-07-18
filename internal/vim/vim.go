// Package vim provides ...
package vim

import "os"

var (
	ConfPath       = os.ExpandEnv("$HOME/.config/xnvim")
	CachePath      = os.ExpandEnv("$HOME/.cache/nvim/")
	ConfCore       = ConfPath + "/core/"
	ConfAutoload   = ConfPath + "/autoload/"
	ConfModules    = ConfPath + "/modules/"
	PluginManage   string
	Leaderkey      string
	LocalLeaderKey string
	Colorscheme    []string
	Explorer       string
	Database       bool
)
