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
	render.WithConfirm(bufferline, vim.ConfCore+"pmap.vim", "vim-buffer keymap", plugin.BuffetKeyMap)
}

func (d *Dein) GenerateStatusLine(spaceline bool) {
	render.WithConfirm(spaceline, vim.ConfModules+"appearance.toml", "Statusline", plugin.DeinStatusline)
}

func (d *Dein) GenerateExplorer(explorer string) {
	if explorer == "coc-explorer" {
		plugin.DeinCocExplorer = true
		render.WriteTemplate(vim.ConfCore+"pmap.vim", "coc-explorer keymap", plugin.CocExplorerKeyMap)
	} else if explorer == "defx.nvim" {
		render.WriteTemplate(vim.ConfModules+"appearance.toml", "Defx.nvim", plugin.DeinDefx)
		render.WriteTemplate(vim.ConfCore+"pmap.vim", "defx keymap", plugin.DefxKeyMap)
		render.WriteTemplate(vim.ConfCore+"pmap.vim", "defx keymap", plugin.DefxFindKeyMap)
	} else {
		render.WriteTemplate(vim.ConfModules+"appearance.toml", "Nerdtree", plugin.DeinNerdTree)
		render.WriteTemplate(vim.ConfCore+"pmap.vim", "nerdtree keymap", plugin.NerdTreeKeyMap)
	}
}

func (d *Dein) GenerateDatabase(database bool) {
	if database {
		render.WriteTemplate(vim.ConfAutoload+"initself.vim", "LoadEnv function", plugin.AutoloadLoadEnv)
		render.WriteTemplate(vim.ConfModules+"database.toml", "Database", plugin.DeinDatabase)
		render.WriteTemplate(vim.ConfCore+"pmap.vim", "database keymap", plugin.DataBaseKeyMap)
	} else {
		color.PrintWarn("Skip Generate Datbase")
	}
}

func (d *Dein) GenerateFuzzyFind(fuzzyfind bool) {
	render.WithConfirm(fuzzyfind, vim.ConfModules+"fuzzyfind.toml", "vim-clap", plugin.DeinClap)
	render.WithConfirm(fuzzyfind, vim.ConfCore+"pmap.vim", "vim-clap keymap", plugin.ClapKeyMap)
	render.WithConfirm(fuzzyfind, vim.ConfCore+"pmap.vim", "coc-clap keymap", plugin.CocClapKeyMap)
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
		render.WriteTemplate(vim.ConfCore+"pmap.vim", "caw.vim keymap", plugin.CawKeyMap)
	} else {
		color.PrintWarn("Skip generate caw.vim config")
	}
}

func (d *Dein) GenerateOutLine(outline bool) {
	render.WithConfirm(outline, vim.ConfModules+"program.toml", "Vista.vim", plugin.DeinVista)
	render.WithConfirm(outline, vim.ConfCore+"pmap.vim", "vista.vim keymap", plugin.VistaKeyMap)
}

func (d *Dein) GenerateTags(tagsplugin bool) {
	render.WithConfirm(tagsplugin, vim.ConfModules+"program.toml", "vim-gutentags", plugin.DeinGuTenTags)
}

func (d *Dein) GenerateQuickRun(quickrun bool) {
	render.WithConfirm(quickrun, vim.ConfModules+"program.toml", "vim-quickrun", plugin.DeinQuickRun)
	render.WithConfirm(quickrun, vim.ConfCore+"pmap.vim", "quickrun keymap", plugin.QuickRunKeyMap)
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
	var enhancekeymap = map[string]string{
		"vim-mundo":      plugin.MundoKeyMap,
		"vim-easymotion": plugin.EasyMotionKeyMap,
		"vim-floterm":    plugin.FloatermKeyMap,
	}
	render.WriteTemplate(vim.ConfModules+"enhance.toml", "dein.vim", plugin.DeinDein)
	for _, v := range plugins {
		pluginname := strings.Split(v, " ")[0]
		i, ok := enhancepluginmap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"enhance.toml", pluginname, i)
		}
		j, ok := enhancekeymap[pluginname]
		if ok {
			render.WriteTemplate(vim.ConfCore+"pmap.vim", pluginname+"keymap", j)
		}
	}
}

func (d *Dein) GenerateSandWich(sandwich bool) {
	render.WithConfirm(sandwich, vim.ConfModules+"textobj.toml", "vim-sandwich", plugin.DeinSandWich)
	render.WithConfirm(sandwich, vim.ConfCore+"pmap.vim", "vim-sandwich keymap", plugin.SandWichKeyMap)
}

func (d *Dein) GenerateTextObj() {
	render.WithConfirm(true, vim.ConfModules+"textobj.toml", "textobj plugins", plugin.DeinTextObj)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "textobj vim", plugin.NiceBlockKeyMap)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "textobj vim", plugin.VimExpandRegionKeyMap)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "textobj vim", plugin.DsfKeyMap)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "textobj vim", plugin.SplitJoinKeyMap)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "textobj vim", plugin.OperatorReplaceKeyMap)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "textobj vim", plugin.MultiBlockKeyMap)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "textobj vim", plugin.TextObjFunctionKeyMap)
}

func (d *Dein) GenerateVersionControl(userversion []string, versionmap map[string]string) {
	versionkeymap := map[string]string{
		"jreybert/vimagit":   plugin.VimagitKeyMap,
		"tpope/vim-fugitive": plugin.FugiTiveKeyMap,
	}

	for i, v := range userversion {
		_, ok := versionmap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"version.toml", userversion[i], versionmap[v])
		}
		_, ok = versionkeymap[v]
		if ok {
			render.WriteTemplate(vim.ConfCore+"pmap.vim", v+"keymap", versionkeymap[v])
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
	data := []interface{}{plugin.DeinCoC, plugin.DeinCocExplorer}
	for i, k := range UserLanguages {
		v, ok := LanguagesPluginMap[k]
		if ok {
			render.WriteTemplate(vim.ConfModules+"languages.toml", UserLanguages[i], v)
		}
	}
	render.WriteTemplate(vim.ConfAutoload+"initself.vim", "autoload coc function", plugin.AutoloadCoc)
	render.ParseTemplate(vim.ConfModules+"completion.toml", "coc.nvim", plugin.DeinCoC, data)
	render.WriteTemplate(vim.ConfCore+"pmap.vim", "coc.nvim keymap", plugin.CocKeyMap)
}

func (d *Dein) GenerateCocJson() {
	render.WriteTemplate(vim.ConfPath+"/coc-settings.json", "coc-settings.json file", plugin.CocJson)
}

func (d *Dein) GenerateVimMap() {
	render.WriteTemplate(vim.ConfCore+"vmap.vim", "vim map", plugin.VimKeyMap)
}

func (d *Dein) GenerateInstallScripts() {
	render.WriteTemplate(vim.ConfPath+"/Makefile", "Makefile", plugin.DeinMakeFile)
	render.WriteTemplate(vim.ConfPath+"/install.sh", "install.sh", plugin.DeinInstallShell)
}
