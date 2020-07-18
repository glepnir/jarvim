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
	GenerateInit()
	GenerateCore(Leader, LocalLeaderKey string)
	GenerateGeneral()
	GenerateTheme()
	GenerateCacheTheme(colorschemes []string)
	GenerateColorscheme(colorschemes []string)
	GenerateDevIcons()
	GenerateDashboard(dashboard bool)
	GenerateBufferLine(bufferline bool)
	GenerateStatusLine(statusline bool)
	GenerateExplorer(explorer string)
	GenerateDatabase(database bool)
	GenerateFuzzyFind(fuzzfind bool)
	GenerateEditorConfig(editorconfig bool)
	GenerateIndentLine(indentplugin string)
	GenerateComment(comment bool)
	GenerateOutLine(outline bool)
	GenerateTags(tagsplugin bool)
	GenerateQuickRun(quickrun bool)
	GenerateDataTypeFile(datafile []string)
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

func ParseTemplate(name string, plugintemplate string, w io.Writer, data interface{}) {
	tpl := template.Must(template.New("").Parse(plugintemplate))
	err := tpl.Execute(w, data)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess(fmt.Sprintf("Generate %v success", name))
}

func WriteTemplate(targetfile string, name, plugintemplate string) {
	f, err := os.OpenFile(targetfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		RollBack(err)
	}
	_, err = f.WriteString(plugintemplate)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess(fmt.Sprintf("Generate %v success", name))
}
