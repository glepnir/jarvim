// Package render provides ...
package render

import (
	"os"
	"text/template"

	"github.com/glepnir/jarvis/internal/plugin"
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
}

// GenerateCore will generate core/core.vim
func GenerateCore(LeaderKey, LocalLeaderKey string) {
	keymap := map[string]string{
		"Space":        "\\<Space>",
		"Comma(,)":     ",",
		"Semicolon(;)": ";",
	}
	f, err := os.Create(vim.ConfPath + "/core/core.vim")
	if err != nil {
		RollBack(err)
	}
	tpl := template.Must(template.New("").Parse(plugin.Core))
	err = tpl.Execute(f, []string{keymap[LeaderKey], keymap[LocalLeaderKey]})
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess("Generate Core.vim Success")
}

// GenerateGeneral will generate core/general.vim
func GenerateGeneral() {
	f, err := os.OpenFile(vim.ConfPath+"/core/general.vim", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		RollBack(err)
	}
	_, err = f.WriteString(plugin.General)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess("Generate general.vim success")
}

// GenerateTheme will generate autoload/theme.vim
// theme.vim read or write the theme.txt from $CACHE/.vim/theme.txt
func GenerateTheme() {
	f, err := os.OpenFile(vim.ConfPath+"/autoload/theme.vim", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		RollBack(err)
	}
	_, err = f.WriteString(plugin.Theme)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess("Generate theme.vim success")
}

func GenerateColorscheme(colorschemes []string) {
	colorsmap := map[string]string{
		"hardcoreplayers/oceanic-material(author is me)": "oceanic_materail",
		"drewtempelmeyer/palenight.vim":                  "palenight",
		"gruvbox-community/gruvbox":                      "gruvbox",
		"ayu-theme/ayu-vim":                              "ayu",
		"NLKNguyen/papercolor-theme":                     "PaperColor",
		"lifepillar/vim-gruvbox8":                        "gruvbox8",
		"lifepillar/vim-solarized8":                      "solarized8",
		"joshdick/onedark.vim":                           "onedark",
		"arcticicestudio/nord-vim":                       "nord",
		"rakr/vim-one":                                   "one",
		"mhartington/oceanic-next":                       "OceanicNext",
		"dracula/vim":                                    "dracula",
		"chriskempson/base16-vim":                        "base16-default-dark",
		"kristijanhusak/vim-hybrid-material":             "hybrid_material",
		"nanotech/jellybeans.vim":                        "jellybeans",
	}
	colors := colorsmap[colorschemes[0]]
	f, err := os.OpenFile(vim.CachePath+"/theme.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		RollBack(err)
	}
	_, err = f.WriteString(colors)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess("Write colorscheme to theme.txt success")
}
