// Package render provides ...
package render

import (
	"log"
	"os"
	"text/template"

	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
	"github.com/glepnir/jarvis/pkg/util"
)

var (
	templates = map[string]string{
		"dein":   "./../../template/neovim-dein",
		"plugin": "./../../template/neovim-plug"}

	vimConf = os.Getenv("HOME") + "/.config/nvim"
)

func GenerateTemplate() {
	if util.Exist(vimConf) {
		os.Rename(vimConf, vimConf+"-bak")
		color.PrintWarn("Back up your vim config to nvim-bak folder")
	}
	util.CopyDir(templates[vim.PluginManage], vimConf)
}

func GenerateLeaderkey(LeaderKey, LocalLeaderKey string) {
	keymap := map[string]string{
		"Space":        "\\<Space>",
		"Comma(,)":     ",",
		"Semicolon(;)": ";",
	}
	f, err := os.Create(os.ExpandEnv("$HOME/.config/nvim/core.vim"))
	if err != nil {
		log.Fatal(err)
		return
	}
	tpl := template.Must(template.ParseFiles(templates[vim.PluginManage] + "/core/core.vim"))
	tpl.Execute(f, []string{keymap[LeaderKey], keymap[LocalLeaderKey]})
}
