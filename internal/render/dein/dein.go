// Package render provides ...
package dein

import (
	"os"
	"text/template"

	"github.com/glepnir/jarvis/internal/plugin"
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
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
	f, err := os.Create(vim.ConfPath + "/core/core.vim")
	if err != nil {
		render.RollBack(err)
	}
	tpl := template.Must(template.New("").Parse(plugin.Core))
	err = tpl.Execute(f, []string{keymap[LeaderKey], keymap[LocalLeaderKey]})
	if err != nil {
		render.RollBack(err)
	}
	color.PrintSuccess("Generate Core.vim Success")
}

// GenerateGeneral will generate core/general.vim
func (d *Dein) GenerateGeneral() {
	f, err := os.OpenFile(vim.ConfPath+"/core/general.vim", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		render.RollBack(err)
	}
	_, err = f.WriteString(plugin.General)
	if err != nil {
		render.RollBack(err)
	}
	color.PrintSuccess("Generate general.vim success")
}

// GenerateTheme will generate autoload/theme.vim
// theme.vim read or write the theme.txt from $CACHE/.vim/theme.txt
func (d *Dein) GenerateTheme() {
	f, err := os.OpenFile(vim.ConfPath+"/autoload/theme.vim", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		render.RollBack(err)
	}
	_, err = f.WriteString(plugin.Theme)
	if err != nil {
		render.RollBack(err)
	}
	color.PrintSuccess("Generate theme.vim success")
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
	f, err := os.OpenFile(vim.CachePath+"/theme.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		render.RollBack(err)
	}
	_, err = f.WriteString(colors)
	if err != nil {
		render.RollBack(err)
	}
	color.PrintSuccess("Write colorscheme to theme.txt success")
}

func (d *Dein) GenerateColorscheme(colorschemes []string) {
	f, err := os.OpenFile(vim.ConfPath+"/modules/appearance.toml", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		render.RollBack(err)
	}
	tpl := template.Must(template.New("").Parse(plugin.DeinColorscheme))
	tpl.Execute(f, colorschemes)
	color.PrintSuccess("Generate Colorscheme Plugins success")
}

func (d *Dein) GenerateStatusLine() {
	f, err := os.OpenFile(vim.ConfPath+"/modules/appearance.toml", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		render.RollBack(err)
	}
	_, err = f.WriteString(plugin.DeinStatusline)
	if err != nil {
		render.RollBack(err)
	}
	color.PrintSuccess("Generate statusline success")
}
