// Package logic provides ...
package logic

import (
	"github.com/glepnir/jarvis/internal/plugin"
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/render/dein"
	"github.com/glepnir/jarvis/internal/render/vimplug"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/cli"
)

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
func LeaderKey() string {
	message := "What is your Leader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

func LocalLeaderKey() string {
	message := "What is your LocalLeader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

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

func DashboardPlugin() bool {
	message := "Do you want use dashboard-nvim a better StartScreenPlugin?"
	return cli.ConfirmTemplate(message)
}

func BufferLinePlugin() bool {
	message := "Do you want use vim-buffet as your bufferline?"
	return cli.ConfirmTemplate(message)
}

func SpacelinePlugin() bool {
	message := "Do you want use spaceline.vim a light and beautiful statusline?"
	return cli.ConfirmTemplate(message)
}

func ExplorerPlugin() string {
	message := "What is your explorer plugin?"
	options := []string{"defx.nvim", "nerdtree", "coc-explorer"}
	return cli.SingleSelectTemplate(message, options)
}

func DatabasePlugin() bool {
	message := "Do you need database plugins?"
	return cli.ConfirmTemplate(message)
}

func FuzzyFindPlugin() bool {
	message := "Do you want use fuzzy find plugin vim-clap?"
	return cli.ConfirmTemplate(message)
}

func EditorConfigPlugin() bool {
	message := "Do you want use editorconfig to control program style(like indent,whitespace etc)"
	return cli.ConfirmTemplate(message)
}

func IndentLinePlugin() string {
	message := "Choose  your favorite indentline plugin?"
	options := []string{"Yggdroot/indentLine", "nathanaelkane/vim-indent-guides"}
	return cli.SingleSelectTemplate(message, options)
}

func CommentPlugin() bool {
	message := "Do you want to use Caw.vim as comment plugin?"
	return cli.ConfirmTemplate(message)
}

func ViewSymbolsPlugin() bool {
	message := "Do you want to use vista.vim to view tags and LSP symbols in sidebar"
	return cli.ConfirmTemplate(message)
}

func GentagsPlugin() bool {
	message := "Do you want to use vim-gutentags to gen tags"
	return cli.ConfirmTemplate(message)
}

func QuickRunPlugin() bool {
	message := "Do you want to use vim-quickrun to fast run program in vim?"
	return cli.ConfirmTemplate(message)
}

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

func SandWichPlugin() bool {
	message := "Do you want use vim-sandwich more useful than vim-surround?"
	return cli.ConfirmTemplate(message)
}

func VersionControlPlugin() []string {
	questionname := "Version Control plugin"
	message := "Choose the version control plugins that you need"
	pagesize := 10

	options := make([]string, 0)
	for k, _ := range vim.VersionControlMap {
		options = append(options, k)
	}

	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

func LanguageServerProtocol() []string {
	questionname := "LanguageQuestion"
	message := "What Languages do you write"
	pagesize := 19
	options := make([]string, 0)
	for k, _ := range vim.LanguagesPluginMap {
		options = append(options, k)
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}
