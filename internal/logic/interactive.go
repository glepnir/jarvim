// Copyright 2020 The jarvim Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logic

import (
	"github.com/glepnir/jarvim/internal/plugin"
	"github.com/glepnir/jarvim/internal/render"
	"github.com/glepnir/jarvim/internal/render/dein"
	"github.com/glepnir/jarvim/internal/render/vimplug"
	"github.com/glepnir/jarvim/internal/vim"
	"github.com/glepnir/jarvim/pkg/cli"
)

// PluginManage return the plugin management plugin
// that user select
func PluginManage() render.Render {
	message := "What is plugin manage do you use?"
	options := []string{"dein", "vim-plug"}
	pm := cli.SingleSelectTemplate(message, options)
	if pm == "dein" {
		return new(dein.Dein)
	} else {
		return new(vimplug.VimPlug)
	}
}

// NewDatFileMap according the render type return
// datafilemap
func NewDataFileMap(r render.Render) map[string]string {
	_, ok := r.(*dein.Dein)
	if ok {
		return map[string]string{
			"MarkDown":   plugin.DeinMarkDown,
			"Toml":       plugin.DeinToml,
			"Nginx":      plugin.DeinNginx,
			"Json":       plugin.DeinJson,
			"Dockerfile": plugin.DeinDockerFile,
		}
	}
	return map[string]string{
		"MarkDown":   plugin.PlugMarkDown,
		"Toml":       plugin.PlugToml,
		"Nginx":      plugin.PlugNginx,
		"Json":       plugin.PlugJson,
		"Dockerfile": plugin.PlugDockerFile,
	}

}

// NewENewEnhancePluginMap return the enhance plugin map
// according plugin management type
func NewEnhancePluginMap(r render.Render) map[string]string {
	_, ok := r.(*dein.Dein)
	if ok {
		return map[string]string{
			"accelerated-jk accelerate up-down moving (j and k mapping)": plugin.DeinFastJK,
			"vim-mundo  vim undo tree":                                   plugin.DeinMundo,
			"vim-easymotion fast jump":                                   plugin.DeinEasyMotion,
			"rainbow  rainbow parentheses":                               plugin.DeinRainbow,
			"vim-floterm  vim terminal float":                            plugin.DeinFloaterm,
		}
	}
	return map[string]string{
		"accelerated-jk accelerate up-down moving (j and k mapping)": plugin.PlugFastJK,
		"vim-mundo  vim undo tree":                                   plugin.PlugMundo,
		"vim-easymotion fast jump":                                   plugin.PlugEasyMotion,
		"rainbow  rainbow parentheses":                               plugin.PlugRainbow,
		"vim-floterm  vim terminal float":                            plugin.PlugFloaterm,
	}
}

// NewVersionPluginMap return the version control map
func NewVersionPluginMap(r render.Render) map[string]string {
	_, ok := r.(*dein.Dein)
	if ok {
		return map[string]string{
			"jreybert/vimagit":     plugin.DeinVimagt,
			"tpope/vim-fugitive":   plugin.DeinFugiTive,
			"lambdalisue/gina.vim": plugin.DeinGina,
		}
	}
	return map[string]string{
		"jreybert/vimagit":     plugin.PlugVimagit,
		"tpope/vim-fugitive":   plugin.PlugFugTive,
		"lambdalisue/gina.vim": plugin.PlugGina,
	}

}

// NewLanguagePlugMap return the lanuages config map
func NewLanguagePlugMap(r render.Render) map[string]string {
	_, ok := r.(*dein.Dein)
	if ok {
		return map[string]string{
			"C-family":   plugin.DeinCFamily,
			"R":          plugin.DeinR,
			"Javascript": plugin.DeinJavascript,
			"Typescript": plugin.DeinTypescript,
			"Dart":       plugin.DeinDart,
			"React":      plugin.DeinReact,
			"Vue":        plugin.DeinVue,
			"Go":         plugin.DeinGo,
			"Rust":       plugin.DeinRust,
			"Haskell":    plugin.DeinHaskell,
			"Php":        plugin.DeinPhp,
			"Ruby":       plugin.DeinRuby,
			"Scala":      plugin.DeinScala,
			"Shell":      plugin.DeinShell,
			"Lua":        plugin.DeinLua,
			"Python":     plugin.DeinPython,
			"Html":       plugin.DeinHtml,
			"Css":        plugin.DeinCss,
			"Less":       plugin.DeinLess,
			"Sass scss":  plugin.DeinSass,
			"Stylus":     plugin.DeinStylus,
		}
	}
	return map[string]string{
		"C-family":   plugin.PlugCFamily,
		"R":          plugin.PlugR,
		"Javascript": plugin.PlugJavascript,
		"Typescript": plugin.PlugTypescript,
		"Dart":       plugin.PlugDart,
		"React":      plugin.PlugReact,
		"Vue":        plugin.PlugVue,
		"Go":         plugin.PlugGo,
		"Rust":       plugin.PlugRust,
		"Haskell":    plugin.PlugHaskell,
		"Php":        plugin.PlugPhp,
		"Ruby":       plugin.PlugRuby,
		"Scala":      plugin.PlugScala,
		"Shell":      plugin.PlugShell,
		"Lua":        plugin.PlugLua,
		"Python":     plugin.PlugPython,
		"Html":       plugin.PlugHtml,
		"Css":        plugin.PlugCss,
		"Less":       plugin.PlugLess,
		"Sass scss":  plugin.PlugSass,
		"Stylus":     plugin.PlugStylus,
	}

}

// Leaderkey get the user LeaderKey
func LeaderKey() string {
	message := "What is your Leader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

// LocalLeaderKey get the user LocalLeaderKey
func LocalLeaderKey() string {
	message := "What is your LocalLeader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

// Colorscheme get the user colorshemes
func Colorscheme() []string {
	questionname := "Colorscheme Question"
	message := "Choose your favorite colorscheme"
	pagesize := 19
	options := make([]string, 0)
	for k, _ := range vim.ColorschemeMap {
		options = append(options, k)
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

// DashboardPlugin return bool according user choose
func DashboardPlugin() bool {
	message := "Do you want use dashboard-nvim a better StartScreenPlugin?"
	return cli.ConfirmTemplate(message)
}

// BufferLinePlugin return bool according user choose
func BufferLinePlugin() bool {
	message := "Do you want use vim-buffet as your bufferline?"
	return cli.ConfirmTemplate(message)
}

// SpacelinePlugin return bool according user choose
func SpacelinePlugin() bool {
	message := "Do you want use spaceline.vim a light and beautiful statusline?"
	return cli.ConfirmTemplate(message)
}

// ExplorerPlugin return the explorer plugin
func ExplorerPlugin() string {
	message := "What is your explorer plugin?"
	options := []string{"defx.nvim", "nerdtree", "coc-explorer"}
	return cli.SingleSelectTemplate(message, options)
}

// DatabasePlugin return bool according user choose
func DatabasePlugin() bool {
	message := "Do you need database plugins?"
	return cli.ConfirmTemplate(message)
}

// FuzzyFindPlugin return bool according user choose
func FuzzyFindPlugin() bool {
	message := "Do you want use fuzzy find plugin vim-clap?"
	return cli.ConfirmTemplate(message)
}

// EditorConfigPlugin return bool according user choose
func EditorConfigPlugin() bool {
	message := "Do you want use editorconfig to control program style(like indent,whitespace etc)"
	return cli.ConfirmTemplate(message)
}

// IndentLinePlugin return string according user choose
func IndentLinePlugin() string {
	message := "Choose  your favorite indentline plugin?"
	options := []string{"Yggdroot/indentLine", "nathanaelkane/vim-indent-guides"}
	return cli.SingleSelectTemplate(message, options)
}

// CommentPlugin return bool according user choose
func CommentPlugin() bool {
	message := "Do you want to use Caw.vim as comment plugin?"
	return cli.ConfirmTemplate(message)
}

// ViewSymbolsPlugin return bool according user choose
func ViewSymbolsPlugin() bool {
	message := "Do you want to use vista.vim to view tags and LSP symbols in sidebar"
	return cli.ConfirmTemplate(message)
}

// GentagsPlugin return bool according user choose
func GentagsPlugin() bool {
	message := "Do you want to use vim-gutentags to gen tags"
	return cli.ConfirmTemplate(message)
}

// QuickRunPlugin return bool according user choose
func QuickRunPlugin() bool {
	message := "Do you want to use vim-quickrun to fast run program in vim?"
	return cli.ConfirmTemplate(message)
}

// DataTypeFile return string slice according user choose
func DataTypeFile(r render.Render) []string {
	questionname := "Data filetype"
	message := "Which Data filetype you need?"
	pagesize := 10
	options := make([]string, 0)
	for k, _ := range NewDataFileMap(r) {
		options = append(options, k)
	}

	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

// EnhancePlugin return string slice according user choose
func EnhancePlugin(r render.Render) []string {
	questionname := "Enhance question"
	message := "Choose the enhance plugins that you need "
	pagesize := 10
	options := make([]string, 0)
	for k, _ := range NewEnhancePluginMap(r) {
		options = append(options, k)
	}

	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

// SandWichPlugin return bool according user choose
func SandWichPlugin() bool {
	message := "Do you want use vim-sandwich more useful than vim-surround?"
	return cli.ConfirmTemplate(message)
}

// VersionControlPlugin return string slice according user choose
func VersionControlPlugin(r render.Render) []string {
	questionname := "Version Control plugin"
	message := "Choose the version control plugins that you need"
	pagesize := 10

	options := make([]string, 0)
	for k, _ := range NewVersionPluginMap(r) {
		options = append(options, k)
	}

	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

// LanguageServerProtocol return string slice according user choose
func LanguageServerProtocol(r render.Render) []string {
	questionname := "LanguageQuestion"
	message := "What Languages do you write"
	pagesize := 19
	options := make([]string, 0)
	for k, _ := range NewLanguagePlugMap(r) {
		options = append(options, k)
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}
