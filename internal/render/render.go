// Package render provides ...
package render

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
)

type Render interface {
	GenerateCore(Leader, LocalLeaderKey string)
	GenerateGeneral()
	GenerateTheme()
	GenerateCacheTheme(colorschemes []string)
	GenerateColorscheme(colorschemes []string)
	GenerateDevIcons()
	GenerateStatusLine()
}

func NewRender(p Render) Render {
	var r Render = p
	return r
}

// RollBack will remove the config folder when
// some errors disappear
func RollBack(err error) {
	color.PrintError(err.Error())
	color.PrintWarn("Rolling back...")
	os.RemoveAll(vim.ConfPath)
	os.Exit(1)
}

func ParseTemplate(plugin string, plugintemplate string, w io.Writer, data interface{}) {
	tpl := template.Must(template.New("").Parse(plugintemplate))
	err := tpl.Execute(w, data)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess(fmt.Sprintf("Generate %v success", plugin))
}

func WriteTemplate(file string, plugin, plugintemplate string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		RollBack(err)
	}
	_, err = f.WriteString(plugintemplate)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess(fmt.Sprintf("Generate %v success", plugin))
}
