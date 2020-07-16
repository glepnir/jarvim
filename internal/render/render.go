// Package render provides ...
package render

import (
	"os"
	"text/template"

	"github.com/glepnir/jarvis/internal/plugin"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
	"github.com/glepnir/jarvis/pkg/util"
)

var (
	templates = map[string]string{
		"dein":   "./../../template/neovim-dein",
		"plugin": "./../../template/neovim-plug"}
)

func GenerateTemplate() {
	if util.Exist(vim.ConfPath) {
		os.Rename(vim.ConfPath, vim.ConfPath+"-bak")
		color.PrintWarn("Back up your vim config to nvim-bak folder")
	}
	util.CopyDir(templates[vim.PluginManage], vim.ConfPath)
}

func GenerateLeaderkey(LeaderKey, LocalLeaderKey string) {
	keymap := map[string]string{
		"Space":        "\\<Space>",
		"Comma(,)":     ",",
		"Semicolon(;)": ";",
	}
	f, err := os.Create(os.ExpandEnv("$HOME/.config/nvim/core.vim"))
	if err != nil {
		color.PrintError(err.Error())
		return
	}
	tpl := template.Must(template.ParseFiles(templates[vim.PluginManage] + "/core/core.vim"))
	tpl.Execute(f, []string{keymap[LeaderKey], keymap[LocalLeaderKey]})
}

func GenerateGeneral() {
	if !util.Exist(vim.ConfPath + "/core") {
		os.Mkdir(vim.ConfPath+"/core", 0700)
	}
	f, err := os.OpenFile(vim.ConfPath+"/core/general.vim", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		color.PrintError(err.Error())
		os.Exit(1)
	}
	_, err = f.WriteString(plugin.General)
	if err != nil {
		color.PrintError(err.Error())
		os.Exit(1)
	}
}
