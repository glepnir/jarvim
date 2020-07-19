// Package render provides ...
package dein

import (
	"os"
	"strings"

	"github.com/glepnir/jarvis/internal/plugin"
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
	"github.com/glepnir/jarvis/pkg/util"
)

type Dein struct{}

var _ render.Render = (*Dein)(nil)

func (d *Dein) GenerateInit() {
	render.WithConfirm(true, vim.ConfPath+"/init.vim", "init.vim", plugin.InitVim)
}

// GenerateCore will generate core/core.vim
func (d *Dein) GenerateCore(LeaderKey, LocalLeaderKey string, leaderkeymap map[string]string) {
	render.ParseTemplate(vim.ConfCore+"core.vim", "core/core.vim", plugin.Core, []string{leaderkeymap[LeaderKey], leaderkeymap[LocalLeaderKey]})
}

func (d *Dein) GeneratePlugMan() {
	render.WithConfirm(true, vim.ConfCore+"dein.vim", "dein.vim", plugin.Dein)
}

// GenerateGeneral will generate core/general.vim
func (d *Dein) GenerateGeneral() {
	render.WithConfirm(true, vim.ConfCore+"general.vim", "core/general.vim", plugin.General)
	render.WithConfirm(true, vim.ConfCore+"event.vim", "core/event.vim", plugin.Event)
}

func (d *Dein) GenerateAutoloadFunc() {
	render.WithConfirm(true, vim.ConfAutoload+"initself.vim", "autoload/initself.vim", plugin.AutoloadSourceFile)
	render.WithConfirm(true, vim.ConfAutoload+"initself.vim", "autoload/initself.vim", plugin.AutoloadMkdir)
}

// GenerateTheme will generate autoload/theme.vim
// theme.vim read or write the theme.txt from $CACHE/.vim/theme.txt
func (d *Dein) GenerateTheme() {
	render.WithConfirm(true, vim.ConfAutoload+"theme.vim", "autoload/theme.vim", plugin.Theme)
}

func (d *Dein) GenerateCacheTheme(usercolors []string, colorschememap map[string]string) {
	colors := colorschememap[usercolors[0]]
	render.WriteTemplate(vim.CachePath+"theme.txt", "theme.txt", colors)
}

func (d *Dein) GenerateColorscheme(usercolors []string) {
	render.ParseTemplate(vim.ConfModules+"appearance.toml", "Colorscheme", plugin.DeinColorscheme, usercolors)
}

func (d *Dein) GenerateDevIcons() {
	render.WriteTemplate(vim.ConfModules+"appearance.toml", "Vim-Devicons", plugin.DeinDevicons)
}

func (d *Dein) GenerateDashboard(dashboard bool) {
	render.WithConfirm(dashboard, vim.ConfModules+"appearance.toml", "dashboard-nvim", plugin.DeinDashboard)
}

func (d *Dein) GenerateBufferLine(bufferline bool) {
	render.WithConfirm(bufferline, vim.ConfModules+"appearance.toml", "Vim-Buffer", plugin.DeinBufferLine)
}

func (d *Dein) GenerateStatusLine(spaceline bool) {
	render.WithConfirm(spaceline, vim.ConfModules+"appearance.toml", "Statusline", plugin.DeinStatusline)
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
	render.WithConfirm(fuzzyfind, vim.ConfModules+"fuzzyfind.toml", "vim-clap", plugin.DeinClap)
}

func (d *Dein) GenerateEditorConfig(editorconfig bool) {
	render.WithConfirm(editorconfig, vim.ConfModules+"program.toml", "editorconfig", plugin.DeinEditorConfig)
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
	render.WithConfirm(outline, vim.ConfModules+"program.toml", "Vista.vim", plugin.DeinVista)
}

func (d *Dein) GenerateTags(tagsplugin bool) {
	render.WithConfirm(tagsplugin, vim.ConfModules+"program.toml", "vim-gutentags", plugin.DeinGuTenTags)
}

func (d *Dein) GenerateQuickRun(quickrun bool) {
	render.WithConfirm(quickrun, vim.ConfModules+"program.toml", "vim-quickrun", plugin.DeinQuickRun)
}

func (d *Dein) GenerateDataTypeFile(datafile []string, datafilemap map[string]string) {
	for _, f := range datafile {
		_, ok := datafilemap[f]
		if ok {
			render.WriteTemplate(vim.ConfModules+"filetype.toml", f, datafilemap[f])
		}
	}
}

func (d *Dein) GenerateEnhanceplugin(plugins []string, enhancepluginmap map[string]string) {
	for _, v := range plugins {
		_, ok := enhancepluginmap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"enhance.toml", strings.Split(v, " ")[0], enhancepluginmap[v])
		}
	}
}

func (d *Dein) GenerateSandWich(sandwich bool) {
	render.WithConfirm(sandwich, vim.ConfModules+"textobj.toml", "vim-sandwich", plugin.DeinSandWich)
}

func (d *Dein) GenerateTextObj() {
	render.WithConfirm(true, vim.ConfModules+"textobj.toml", "textobj plugins", plugin.DeinTextObj)
}

func (d *Dein) GenerateVersionControl(userversion []string, versionmap map[string]string) {
	for i, v := range userversion {
		_, ok := versionmap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"version.toml", userversion[i], versionmap[v])
		}
	}
	render.WriteTemplate(vim.ConfModules+"version.toml", "committia.vim", plugin.DeinCommita)
}

func (d *Dein) GeneratePluginFolder() {
	err := util.CopyDir("./../../plugin", vim.ConfPlugin)
	if err != nil {
		color.PrintError("Copy plugin folder to your vim config path failed")
		os.Exit(0)
	}
	color.PrintSuccess("Generate plugin folder success")
}

func (d *Dein) GenerateLanguagePlugin(UserLanguages []string, LanguagesPluginMap map[string]string) {
	for i, k := range UserLanguages {
		v, ok := LanguagesPluginMap[k]
		if ok {
			render.WriteTemplate(vim.ConfModules+"languages.toml", UserLanguages[i], v)
		}
	}
}
