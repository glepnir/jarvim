// Package plug provides ...
package vimplug

import (
	"os"
	"strings"
	"sync"

	"github.com/glepnir/jarvim/internal/plugin"
	"github.com/glepnir/jarvim/internal/render"
	"github.com/glepnir/jarvim/internal/vim"
	"github.com/glepnir/jarvim/pkg/color"
	"github.com/glepnir/jarvim/template"
)

// VimPlug struct
type VimPlug struct{}

// Ensure VimPlug struct implement Render interface
var _ render.Render = (*VimPlug)(nil)

// GenerateInit generate init.vim
func (v *VimPlug) GenerateInit() {
	render.WithConfirm(true, vim.ConfPath+"/init.vim", "init.vim", plugin.InitVim)
	os.MkdirAll(vim.ConfModules+"appearance", 0700)
	os.MkdirAll(vim.ConfModules+"program", 0700)
	os.MkdirAll(vim.ConfModules+"enhance", 0700)
	os.MkdirAll(vim.ConfModules+"textobj", 0700)
	os.MkdirAll(vim.ConfModules+"filetype", 0700)
	os.MkdirAll(vim.ConfModules+"languages", 0700)
	os.MkdirAll(vim.ConfModules+"fuzzyfind", 0700)
	os.MkdirAll(vim.ConfModules+"version", 0700)
	os.MkdirAll(vim.ConfModules+"completion", 0700)
}

// GenerateCore will generate core/core.vim
func (v *VimPlug) GenerateCore(LeaderKey, LocalLeaderKey string, leaderkeymap map[string]string) {
	render.ParseTemplate(vim.ConfCore+"core.vim", "core/core.vim", plugin.PlugCore, []string{leaderkeymap[LeaderKey], leaderkeymap[LocalLeaderKey]})
}

// GeneratePlugMan
func (v *VimPlug) GeneratePlugMan() {
	render.WithConfirm(true, vim.ConfCore+"plug.vim", "plug.vim", plugin.VimPlug)
}

// GenerateGeneral will generate core/general.vim
func (v *VimPlug) GenerateGeneral() {
	render.WithConfirm(true, vim.ConfCore+"general.vim", "core/general.vim", plugin.General)
	render.WithConfirm(true, vim.ConfCore+"event.vim", "core/event.vim", plugin.Event)
}

// GenerateAutoloadFunc generate autoload/initself.vim
func (v *VimPlug) GenerateAutoloadFunc() {
	render.WithConfirm(true, vim.ConfAutoload+"initself.vim", "autoload/initself.vim", plugin.AutoloadSourceFile)
	render.WithConfirm(true, vim.ConfAutoload+"initself.vim", "autoload/initself.vim", plugin.AutoloadMkdir)
}

// GenerateTheme will generate autoload/theme.vim
// theme.vim read or write the theme.txt from $CACHE/.vim/theme.txt
func (v *VimPlug) GenerateTheme() {
	render.WithConfirm(true, vim.ConfAutoload+"theme.vim", "autoload/theme.vim", plugin.Theme)
}

// GenerateCacheTheme write the colorscheme into .cache/vim/theme.txt
func (v *VimPlug) GenerateCacheTheme(usercolors []string, colorschememap map[string]string) {
	colors := colorschememap[usercolors[0]]
	render.WriteTemplate(vim.CachePath+"theme.txt", "theme.txt", colors)
}

// GenerateColorscheme generate modules/appearance.toml
func (v *VimPlug) GenerateColorscheme(usercolors []string) {
	render.ParseTemplate(vim.ConfModules+"appearance/plugins.vim", "Colorscheme", plugin.PlugColorscheme, usercolors)
}

// GenerateDevIcons
func (v *VimPlug) GenerateDevIcons() {
	render.WriteTemplate(vim.ConfModules+"appearance/plugins.vim", "Vim-Devicons", plugin.PlugDevicons)
}

// GenerateDashboard
func (v *VimPlug) GenerateDashboard(dashboard bool) {
	render.WithConfirm(dashboard, vim.ConfModules+"appearance/plugins.vim", "dashboard-nvim", plugin.PlugDashboard)
}

// GenerateBufferLine
func (v *VimPlug) GenerateBufferLine(bufferline bool) {
	render.WithConfirm(bufferline, vim.ConfModules+"appearance/plugins.vim", "Vim-Buffer", plugin.PlugBufferLine)
	render.WithConfirm(bufferline, vim.ConfModules+"appearance/config.vim", "vim-buffer keymap", plugin.BuffetKeyMap)
}

// GenerateStatusLine
func (v *VimPlug) GenerateStatusLine(spaceline bool) {
	render.WithConfirm(spaceline, vim.ConfModules+"appearance/plugins.vim", "spacelinev.vim", plugin.PlugStatusline)
	render.WithConfirm(spaceline, vim.ConfModules+"appearance/config.vim", "spacelinev.vim", plugin.PlugStatuslineSetting)
}

// GenerateExplorer
func (v *VimPlug) GenerateExplorer(explorer string) {
	if explorer == "coc-explorer" {
		plugin.PlugCocExplorer = true
		render.WriteTemplate(vim.ConfCore+"pmap.vim", "coc-explorer keymap", plugin.CocExplorerKeyMap)
	} else if explorer == "defx.nvim" {
		render.WriteTemplate(vim.ConfModules+"appearance/plugins.vim", "Defx.nvim", plugin.PlugDefx)
		render.WriteTemplate(vim.ConfModules+"appearance/config.vim", "defx settings", plugin.PlugDefxSetting)
		render.WriteTemplate(vim.ConfModules+"appearance/config.vim", "defx keymap", plugin.DefxKeyMap)
		render.WriteTemplate(vim.ConfModules+"appearance/config.vim", "defx keymap", plugin.DefxFindKeyMap)
	} else {
		render.WriteTemplate(vim.ConfModules+"appearance/plugins.vim", "Nerdtree", plugin.PlugNerdTree)
		render.WriteTemplate(vim.ConfModules+"appearance/config.vim", "nerdtree config", plugin.PlugNerdTreeSetting)
		render.WriteTemplate(vim.ConfModules+"appearance/config.vim", "nerdtree keymap", plugin.NerdTreeKeyMap)
	}
}

// GenerateDatabase
func (v *VimPlug) GenerateDatabase(database bool) {
	if database {
		render.WriteTemplate(vim.ConfAutoload+"initself.vim", "LoadEnv function", plugin.AutoloadLoadEnv)
		render.WriteTemplate(vim.ConfModules+"program/plugins.vim", "Database", plugin.PlugDatabase)
		render.WriteTemplate(vim.ConfModules+"program/config.vim", "database settings", plugin.PlugDatabaseUiSetting)
		render.WriteTemplate(vim.ConfModules+"program/config.vim", "database keymap", plugin.DataBaseKeyMap)
	} else {
		color.PrintWarn("Skip Generate Datbase")
	}
}

// GenerateFuzzyFind
func (v *VimPlug) GenerateFuzzyFind(fuzzyfind bool) {
	render.WithConfirm(fuzzyfind, vim.ConfModules+"fuzzyfind/plugins.vim", "vim-clap", plugin.PlugClap)
	render.WithConfirm(fuzzyfind, vim.ConfModules+"fuzzyfind/config.vim", "vim-clap setting", plugin.PlugClapSetting)
	render.WithConfirm(fuzzyfind, vim.ConfModules+"fuzzyfind/config.vim", "vim-clap keymap", plugin.ClapKeyMap)
	render.WithConfirm(fuzzyfind, vim.ConfModules+"fuzzyfind/config.vim", "coc-clap keymap", plugin.CocClapKeyMap)
}

// GenerateIndentLine
func (v *VimPlug) GenerateIndentLine(indentplugin string) {
	if indentplugin == "Yggdroot/indentLine" {
		render.WriteTemplate(vim.ConfModules+"program/plugins.vim", indentplugin, plugin.PlugIndentLine)
		render.WriteTemplate(vim.ConfModules+"program/config.vim", indentplugin, plugin.PlugIndentLineSetting)
	} else {
		render.WriteTemplate(vim.ConfModules+"program/plugins.vim", indentplugin, plugin.PlugIndentGuides)
		render.WriteTemplate(vim.ConfModules+"program/config.vim", indentplugin, plugin.PlugIndentGuidesSetting)
	}
}

// GenerateComment
func (v *VimPlug) GenerateComment(comment bool) {
	if comment {
		render.WriteTemplate(vim.ConfModules+"filetype/plugins.vim", "context_filetype.vim", plugin.PlugContextFileType)
		render.WriteTemplate(vim.ConfModules+"program/plugins.vim", "Caw.vim", plugin.PlugCawvim)
	} else {
		color.PrintWarn("Skip generate caw.vim config")
	}
}

// GenerateOutLine
func (v *VimPlug) GenerateOutLine(outline bool) {
	render.WithConfirm(outline, vim.ConfModules+"program/plugins.vim", "Vista.vim", plugin.PlugVista)
	render.WithConfirm(outline, vim.ConfModules+"program/config.vim", "Vista.vim", plugin.PlugVistaSetting)
	render.WithConfirm(outline, vim.ConfModules+"program/config.vim", "vista.vim keymap", plugin.VistaKeyMap)
}

// GenerateTags
func (v *VimPlug) GenerateTags(tagsplugin bool) {
	render.WithConfirm(tagsplugin, vim.ConfModules+"program/plugins.vim", "vim-gutentags", plugin.PlugGuTenTags)
	render.WithConfirm(tagsplugin, vim.ConfModules+"program/config.vim", "vim-gutentags", plugin.PlugGuTenTagsSetting)
}

// GenerateQuickRun
func (v *VimPlug) GenerateQuickRun(quickrun bool) {
	render.WithConfirm(quickrun, vim.ConfModules+"program/plugins.vim", "vim-quickrun", plugin.PlugQuickRun)
	render.WithConfirm(quickrun, vim.ConfModules+"program/config.vim", "vim-quickrun", plugin.PlugQuickRunSetting)
	render.WithConfirm(quickrun, vim.ConfModules+"program/config.vim", "quickrun keymap", plugin.QuickRunKeyMap)
}

// GenerateEditorConfig
func (v *VimPlug) GenerateEditorConfig(editorconfig bool) {
	render.WithConfirm(editorconfig, vim.ConfModules+"program/plugins.vim", "editorconfig", plugin.PlugEditorConfig)
}

// GenerateDataTypeFile
func (v *VimPlug) GenerateDataTypeFile(datafile []string, datafilemap map[string]string) {
	for _, f := range datafile {
		_, ok := datafilemap[f]
		if ok {
			render.WriteTemplate(vim.ConfModules+"filetype/plugins.vim", f, datafilemap[f])
		}
		if f == "MarkDown" {
			render.WriteTemplate(vim.ConfModules+"filetype/config.vim", "Markdwon settings", plugin.PlugMarkDownSetting)
		}
	}
}

// GenerateEnhanceplugin
func (v *VimPlug) GenerateEnhanceplugin(plugins []string, enhancepluginmap map[string]string) {
	var enhancekeymap = map[string]string{
		"vim-mundo":      plugin.MundoKeyMap,
		"vim-easymotion": plugin.EasyMotionKeyMap,
		"vim-floterm":    plugin.FloatermKeyMap,
	}
	var enhancesetting = map[string]string{
		"vim-floaterm":   plugin.PlugFloatermSetting,
		"rainbow":        plugin.PlugRainbowSetting,
		"vim-easymotion": plugin.PlugEasyMotionSetting,
		"accelerated-jk": plugin.PlugFastJKSetting,
	}
	for _, v := range plugins {
		pluginname := strings.Split(v, " ")[0]
		i, ok := enhancepluginmap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"enhance/plugins.vim", pluginname, i)
		}
		k, ok := enhancesetting[pluginname]
		if ok {
			render.WriteTemplate(vim.ConfModules+"enhance/config.vim", pluginname, k)
		}
		j, ok := enhancekeymap[pluginname]
		if ok {
			render.WriteTemplate(vim.ConfModules+"enhance/config.vim", pluginname+"keymap", j)
		}
	}
}

// GenerateSandWich
func (v *VimPlug) GenerateSandWich(sandwich bool) {
	render.WithConfirm(sandwich, vim.ConfModules+"textobj/plugins.vim", "vim-sandwich", plugin.PlugSandWich)
	render.WithConfirm(sandwich, vim.ConfModules+"textobj/config.vim", "vim-sandwich", plugin.PlugSandWichSetting)
}

// GenerateTextObj
func (v *VimPlug) GenerateTextObj() {
	render.WithConfirm(true, vim.ConfModules+"textobj/plugins.vim", "textobj plugins", plugin.PlugTextObj)
	render.WithConfirm(true, vim.ConfModules+"textobj/config.vim", "textobj plugins setting", plugin.PlugTextObjSetting)
	render.WriteTemplate(vim.ConfModules+"textobj/config.vim", "textobj vim keymap", plugin.NiceBlockKeyMap)
	render.WriteTemplate(vim.ConfModules+"textobj/config.vim", "textobj vim keymap", plugin.VimExpandRegionKeyMap)
	render.WriteTemplate(vim.ConfModules+"textobj/config.vim", "textobj vim keymap", plugin.DsfKeyMap)
	render.WriteTemplate(vim.ConfModules+"textobj/config.vim", "textobj vim keymap", plugin.SplitJoinKeyMap)
	render.WriteTemplate(vim.ConfModules+"textobj/config.vim", "textobj vim keymap", plugin.OperatorReplaceKeyMap)
	render.WriteTemplate(vim.ConfModules+"textobj/config.vim", "textobj vim keymap", plugin.MultiBlockKeyMap)
	render.WriteTemplate(vim.ConfModules+"textobj/config.vim", "textobj vim keymap", plugin.TextObjFunctionKeyMap)
}

// GenerateVersionControl
func (v *VimPlug) GenerateVersionControl(userversion []string, versionmap map[string]string) {
	versionkeymap := map[string]string{
		"jreybert/vimagit":   plugin.VimagitKeyMap,
		"tpope/vim-fugitive": plugin.FugiTiveKeyMap,
	}

	for i, v := range userversion {
		_, ok := versionmap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"version/plugins.vim", userversion[i], versionmap[v])
			if v == "vim-fugitive" {
				render.WriteTemplate(vim.ConfModules+"version/config.vim", "vim-fugtive setting", plugin.PlugFugTiveSetting)
			} else {
				render.WriteTemplate(vim.ConfModules+"version/config.vim", "vim-fugtive setting", plugin.PlugVimagitSetting)
			}
		}
		_, ok = versionkeymap[v]
		if ok {
			render.WriteTemplate(vim.ConfModules+"version/config.vim", v+"keymap", versionkeymap[v])
		}
	}
	render.WriteTemplate(vim.ConfModules+"version/plugins.vim", "committia.vim", plugin.PlugCommita)
	render.WriteTemplate(vim.ConfModules+"version/plugins.vim", "committia.vim", plugin.PlugCommitaSetting)

}

// GeneratePluginFolder
func (v *VimPlug) GeneratePluginFolder() {
	render.WriteTemplate(vim.ConfPlugin+"bufkill.vim", "bufkill.vim", template.Bufkill)
	render.WriteTemplate(vim.ConfPlugin+"difftools.vim", "difftools.vim", template.Difftools)
	render.WriteTemplate(vim.ConfPlugin+"hlsearch.vim", "hlsearch.vim", template.Hlsearchvim)
	render.WriteTemplate(vim.ConfPlugin+"nicefold.vim", "nicefold.vim", template.Nicefold)
	render.WriteTemplate(vim.ConfPlugin+"whitespace.vim", "whitespace.vim", template.Whitespace)
	color.PrintSuccess("Generate plugin folder success")
}

// GenerateLanguagePlugin
func (v *VimPlug) GenerateLanguagePlugin(UserLanguages []string, LanguagesPluginMap map[string]string) {
	pluglsp := map[string]string{
		plugin.PlugCFamily:    plugin.PlugCFamilyLsp,
		plugin.PlugR:          plugin.PlugRLsp,
		plugin.PlugJavascript: plugin.PlugJavascriptLsp,
		plugin.PlugTypescript: plugin.PlugTypescriptLsp,
		plugin.PlugDart:       plugin.PlugDartLsp,
		plugin.PlugVue:        plugin.PlugVueLsp,
		plugin.PlugGo:         plugin.PlugGoLsp,
		plugin.PlugRust:       plugin.PlugRustLsp,
		plugin.PlugHaskell:    plugin.PlugHaskellLsp,
		plugin.PlugPhp:        plugin.PlugPhpLsp,
		plugin.PlugRuby:       plugin.PlugRubyLsp,
		plugin.PlugScala:      plugin.PlugScalaLsp,
		plugin.PlugShell:      plugin.PlugShellLsp,
		plugin.PlugLua:        plugin.PlugLuaLsp,
		plugin.PlugPython:     plugin.PlugPythonLsp,
		plugin.PlugHtml:       plugin.PlugHtmlLsp,
		plugin.PlugCss:        plugin.PlugCssLsp,
		plugin.PlugReact:      plugin.PlugReactLsp,
	}

	var once sync.Once
	needemmet := []string{"React", "Vue", "Html"}
	data := []interface{}{plugin.DeinCoC, plugin.PlugCocExplorer}
	for i, k := range UserLanguages {
		v, ok := LanguagesPluginMap[k]
		if ok {
			render.WriteTemplate(vim.ConfModules+"languages/plugins.vim", UserLanguages[i], v)
		}
		l, ok := pluglsp[v]
		if ok {
			if k == "Typescript" || k == "Javascript" {
				once.Do(func() {
					render.WriteTemplate(vim.ConfModules+"languages/config.vim", "Js Ts lsp setting", plugin.PlugTypescriptLsp)
				})

			} else {
				render.WriteTemplate(vim.ConfModules+"languages/config.vim", UserLanguages[i]+"lsp setting", l)
			}
		}
		for _, j := range needemmet {
			once.Do(func() {
				if j == k {
					render.WriteTemplate(vim.ConfModules+"program/plugins.vim", "emmet plugins", plugin.PlugEmmet)
					render.WriteTemplate(vim.ConfModules+"program/config.vim", "emmet plugins setting", plugin.PlugEmmetSetting)
				}
			})
		}
	}
	render.WriteTemplate(vim.ConfAutoload+"initself.vim", "autoload coc function", plugin.AutoloadCoc)
	render.WriteTemplate(vim.ConfModules+"completion/plugins.vim", "completion plugins", plugin.PlugCoc)
	render.ParseTemplate(vim.ConfModules+"completion/config.vim", "coc.nvim", plugin.PlugCocSetting, data)
	render.WriteTemplate(vim.ConfModules+"completion/config.vim", "coc.nvim keymap", plugin.CocKeyMap)
}

// GenerateCocJson
func (v *VimPlug) GenerateCocJson() {
	render.WriteTemplate(vim.ConfPath+"/coc-settings.json", "coc-settings.json file", plugin.CocJson)
}

// GenerateVimMap
func (v *VimPlug) GenerateVimMap() {
	render.WriteTemplate(vim.ConfCore+"vmap.vim", "vim map", plugin.VimKeyMap)
}

// GenerateInstallScripts
func (v *VimPlug) GenerateInstallScripts() {
	render.WriteTemplate(vim.ConfPath+"/Makefile", "Makefile", plugin.PlugMakefile)
	render.WriteTemplate(vim.ConfPath+"/install.sh", "install.sh", plugin.PlugInstallShell)
}
