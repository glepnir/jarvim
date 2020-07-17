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

func BackUpUserConf() {
	if util.Exist(vim.ConfPath) {
		os.Rename(vim.ConfPath, vim.ConfPath+"-bak")
		color.PrintWarn("Back up your vim config to nvim-bak folder")
	}
	// Ensure our folder exists
	os.MkdirAll(vim.ConfPath+"/core", 0700)
}

func GenerateCore(LeaderKey, LocalLeaderKey string) {
	keymap := map[string]string{
		"Space":        "\\<Space>",
		"Comma(,)":     ",",
		"Semicolon(;)": ";",
	}
	f, err := os.Create(vim.ConfPath + "/core/core.vim")
	if err != nil {
		color.PrintError(err.Error())
		return
	}
	tpl := template.Must(template.New("").Parse(plugin.Core))
	err = tpl.Execute(f, []string{keymap[LeaderKey], keymap[LocalLeaderKey]})
	if err != nil {
		color.PrintError(err.Error())
		os.Exit(1)
	}
	color.PrintSuccess("Generate Core.vim Success")
}

func GenerateGeneral() {
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
	color.PrintSuccess("Generate general.vim success")
}
