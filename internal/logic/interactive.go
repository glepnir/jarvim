// Package logic provides ...
package logic

import (
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/render/dein"
	"github.com/glepnir/jarvis/internal/render/vimplug"
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
	options := []string{
		"hardcoreplayers/oceanic-material",
		"drewtempelmeyer/palenight.vim",
		"gruvbox-community/gruvbox",
		"ayu-theme/ayu-vim",
		"NLKNguyen/papercolor-theme",
		"lifepillar/vim-gruvbox8",
		"lifepillar/vim-solarized8",
		"joshdick/onedark.vim",
		"arcticicestudio/nord-vim",
		"rakr/vim-one",
		"mhartington/oceanic-next",
		"patstockwell/vim-monokai-tasty",
		"dracula/vim",
		"chriskempson/base16-vim",
		"kristijanhusak/vim-hybrid-material",
		"kyoz/purify",
		"nanotech/jellybeans.vim",
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

func LspPlugin() string {
	message := "What is your  Lsp plugin?"
	options := []string{"coc.nvim", "nvim-lsp"}
	return cli.SingleSelectTemplate(message, options)
}

func ExplorerPlugin() string {
	message := "What is your explorer plugin?"
	options := []string{"defx.nvim", "nerdtree", "coc-explorer"}
	return cli.SingleSelectTemplate(message, options)
}

func Languages() []string {
	questionname := "LanguageQuestion"
	message := "What Languages do you write"
	pagesize := 19
	options := []string{
		"c-family",
		"R",
		"javascript",
		"typescript",
		"react",
		"vue",
		"go",
		"rust",
		"haskell",
		"php",
		"ruby",
		"scala",
		"shell",
		"lua",
		"python",
		"dockerfile",
		"json",
		"nginx",
		"toml",
		"html",
		"css",
		"less",
		"sass",
		"stylus",
		"sql",
		"dart",
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

func Enhance() []string {
	questionname := "Choose the enhance plugins"
	message := "What is your Leader Key?"
	pagesize := 19
	options := []string{
		"rhysd/accelerated-jk",
		"simnalamburt/vim-mundo",
		"easymotion/vim-easymotion",
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}
