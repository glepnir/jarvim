// Package render provides ...
package dein

import (
	"os"
	"strings"

	"github.com/glepnir/jarvis/internal/plugin"
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
)

type Dein struct{}

var _ render.Render = (*Dein)(nil)

func (d *Dein) GenerateInit() {
	const init = `execute 'source' fnamemodify(expand('<sfile>'), ':h').'/core/core.vim'`
	render.WriteTemplate(vim.ConfPath+"/init.vim", "init.vim", init)
}

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

func (d *Dein) GeneratePlugMan() {
	render.WriteTemplate(vim.ConfCore+"dein.vim", "dein.vim", plugin.Dein)
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

func (d *Dein) GenerateDashboard(dashboard bool) {
	if dashboard {
		render.WriteTemplate(vim.ConfModules+"appearance.toml", "dashboard-nvim", plugin.DeinDashboard)
	} else {
		color.PrintWarn("Skip generate dashboard-nvim config")
	}
}

func (d *Dein) GenerateBufferLine(bufferline bool) {
	if bufferline {
		render.WriteTemplate(vim.ConfModules+"appearance.toml", "Vim-Buffer", plugin.DeinBufferLine)
	} else {
		color.PrintWarn("Skip generate vim-bufferlet config")
	}
}

func (d *Dein) GenerateStatusLine(spaceline bool) {
	if spaceline {
		render.WriteTemplate(vim.ConfModules+"appearance.toml", "Statusline", plugin.DeinStatusline)
	} else {
		color.PrintWarn("Skip generate spaceline.vim config")
	}
}

func (d *Dein) GenerateExplorer(explorer string) {
	if explorer == "coc-explorer" {
	} else if explorer == "defx.nvim" {
		render.WriteTemplate(vim.ConfModules+"appearance.toml", "Defx.nvim", plugin.DeinDefx)
	} else {
		render.WriteTemplate(vim.ConfModules+"appearance.toml", "Nerdtree", plugin.DeinNerdTree)
	}
}

func (d *Dein) GenerateDatabase(database bool) {
	if database {
		render.WriteTemplate(vim.ConfAutoload+"initself.vim", "LoadEnv function", plugin.AutoloadLoadEnv)
		render.WriteTemplate(vim.ConfModules+"database.toml", "Database", plugin.DeinDatabase)
	} else {
		color.PrintWarn("Skip Generate Datbase")
	}
}

func (d *Dein) GenerateFuzzyFind(fuzzyfind bool) {
	if fuzzyfind {
		render.WriteTemplate(vim.ConfModules+"fuzzyfind.toml", "vim-clap", plugin.DeinClap)
	} else {
		color.PrintWarn("Skip generate fuzzyfind vim-clap config ")
	}
}

func (d *Dein) GenerateEditorConfig(editorconfig bool) {
	if editorconfig {
		render.WriteTemplate(vim.ConfModules+"program.toml", "editorconfig", plugin.DeinEditorConfig)
	} else {
		color.PrintWarn("Skip generate editorconfig config")
	}
}

func (d *Dein) GenerateIndentLine(indentplugin string) {
	if indentplugin == "Yggdroot/indentLine" {
		render.WriteTemplate(vim.ConfModules+"program.toml", indentplugin, plugin.DeinIndentLine)
	} else {
		render.WriteTemplate(vim.ConfModules+"program.toml", indentplugin, plugin.DeinIndenGuides)
	}
}

func (d *Dein) GenerateComment(comment bool) {
	if comment {
		render.WriteTemplate(vim.ConfModules+"filetype.toml", "context_filetype.vim", plugin.DeinContextFileType)
		render.WriteTemplate(vim.ConfModules+"program.toml", "Caw.vim", plugin.DeinCaw)
	} else {
		color.PrintWarn("Skip generate caw.vim config")
	}
}

func (d *Dein) GenerateOutLine(outline bool) {
	if outline {
		render.WriteTemplate(vim.ConfModules+"program.toml", "Vista.vim", plugin.DeinVista)
	} else {
		color.PrintWarn("Skip generate vista.vim config")
	}
}

func (d *Dein) GenerateTags(tagsplugin bool) {
	if tagsplugin {
		render.WriteTemplate(vim.ConfModules+"program.toml", "vim-gutentags", plugin.DeinGuTenTags)
	} else {
		color.PrintWarn("Skip generate vim-gutentags config")
	}
}

func (d *Dein) GenerateQuickRun(quickrun bool) {
	if quickrun {
		render.WriteTemplate(vim.ConfModules+"program.toml", "vim-quickrun", plugin.DeinQuickRun)
	} else {
		color.PrintWarn("Skip generate vim-quickrun config")
	}
}

func (d *Dein) GenerateDataTypeFile(datafile []string) {
	datamap := map[string]string{
		"MarkDown":   plugin.DeinMarkDown,
		"Toml":       plugin.DeinToml,
		"Nginx":      plugin.DeinNginx,
		"Json":       plugin.DeinJson,
		"Dockerfile": plugin.DeinDockerFile,
	}
	for _, f := range datafile {
		_, ok := datamap[f]
		if ok {
			render.WriteTemplate(vim.ConfModules+"filetype.toml", f, datamap[f])
		}
	}
}

func (d *Dein) GenerateEnhanceplugin(plugins []string) {
	pluginsmap := map[string]string{
		"accelerated-jk accelerate up-down moving (j and k mapping)": plugin.DeinFastJK,
		"vim-mundo  vim undo tree":                                   plugin.DeinMundo,
		"vim-easymotion fast jump":                                   plugin.DeinEasyMotion,
		"rainbow  rainbow parentheses":                               plugin.DeinRainbow,
		"vim-floterm  vim terminal float":                            plugin.DeinFloaterm,
	}
	for _, v := range plugins {
		_, ok := pluginsmap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"enhance.toml", strings.Split(v, " ")[0], pluginsmap[v])
		}
	}
}
