// Package render provides ...
package dein

import (
	"os"

	"github.com/glepnir/jarvis/internal/plugin"
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/vim"
)

type Dein struct{}

var _ render.Render = (*Dein)(nil)

// GenerateCore will generate core/core.vim
func (d *Dein) GenerateCore(LeaderKey, LocalLeaderKey string) {
	keymap := map[string]string{
		"Space":        "\\<Space>",
		"Comma(,)":     ",",
		"Semicolon(;)": ";",
	}
	f, err := os.Create(vim.ConfCore + "core.vim")
	if err != nil {
		render.RollBack(err)
	}
	render.ParseTemplate("core/core.vim", plugin.Core, f, []string{keymap[LeaderKey], keymap[LocalLeaderKey]})
}

// GenerateGeneral will generate core/general.vim
func (d *Dein) GenerateGeneral() {
	render.WriteTemplate(vim.ConfCore+"general.vim", "core/general.vim", plugin.General)
}

// GenerateTheme will generate autoload/theme.vim
// theme.vim read or write the theme.txt from $CACHE/.vim/theme.txt
func (d *Dein) GenerateTheme() {
	render.WriteTemplate(vim.ConfAutoload+"theme.vim", "Autoload/theme.vim", plugin.Theme)
}

func (d *Dein) GenerateCacheTheme(colorschemes []string) {
	colorsmap := map[string]string{
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
	colors := colorsmap[colorschemes[0]]
	render.WriteTemplate(vim.CachePath+"theme.txt", "theme.txt", colors)
}

func (d *Dein) GenerateColorscheme(colorschemes []string) {
	f, err := os.OpenFile(vim.ConfModules+"appearance.toml", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		render.RollBack(err)
	}
	render.ParseTemplate("Colorscheme", plugin.DeinColorscheme, f, colorschemes)
}

func (d *Dein) GenerateDevIcons() {
	render.WriteTemplate(vim.ConfModules+"appearance.toml", "Vim-Devicons", plugin.DeinDevicons)
}

func (d *Dein) GenerateStatusLine() {
	render.WriteTemplate(vim.ConfModules+"appearance.toml", "Statusline", plugin.DeinStatusline)
}
