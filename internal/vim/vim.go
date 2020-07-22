// Package vim provides ...
package vim

import (
	"os"
)

var (
	ConfPath             = os.ExpandEnv("$HOME/.config/nvim")
	CachePath            = os.ExpandEnv("$HOME/.cache/vim/")
	ConfCore             = ConfPath + "/core/"
	ConfAutoload         = ConfPath + "/autoload/"
	ConfModules          = ConfPath + "/modules/"
	ConfPlugin           = ConfPath + "/plugin/"
	PluginManage         string
	Leaderkey            string
	LocalLeaderKey       string
	Colorscheme          []string
	StartScreenPlugin    bool
	StatusLine           bool
	BufferLine           bool
	Explorer             string
	Database             bool
	Fuzzyfind            bool
	EditorConfig         bool
	IndentPlugin         string
	CommentPlugin        bool
	OutLinePlugin        bool
	TagsPlugin           bool
	QuickRun             bool
	DataTypeFile         []string
	EnhancePlugins       []string
	SandwichPlugin       bool
	VersionControlPlugin []string
	UserLanguages        []string
	CocExtensions        = make([]string, 0)

	LeaderKeyMap = map[string]string{
		"Space":        "\\<Space>",
		"Comma(,)":     ",",
		"Semicolon(;)": ";",
	}

	ColorschemeMap = map[string]string{
		"hardcoreplayers/oceanic-material":   "oceanic_materail",
		"drewtempelmeyer/palenight.vim":      "palenight",
		"gruvbox-community/gruvbox":          "gruvbox",
		"ayu-theme/ayu-vim":                  "ayu",
		"NLKNguyen/papercolor-theme":         "PaperColor",
		"lifepillar/vim-gruvbox8":            "gruvbox8",
		"lifepillar/vim-solarized8":          "solarized8",
		"joshdick/onedark.vim":               "onedark",
		"arcticicestudio/nord-vim":           "nord",
		"rakr/vim-one":                       "one",
		"mhartington/oceanic-next":           "OceanicNext",
		"dracula/vim":                        "dracula",
		"chriskempson/base16-vim":            "base16-default-dark",
		"kristijanhusak/vim-hybrid-material": "hybrid_material",
		"nanotech/jellybeans.vim":            "jellybeans",
	}
)
