// Package render provides ...
package render

import (
	"fmt"
	"os"
	"text/template"

	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
)

// Render interface
type Render interface {
	GenerateInit()
	GenerateCore(Leader, LocalLeaderKey string, leaderkeymap map[string]string)
	GeneratePlugMan()
	GenerateGeneral()
	GenerateAutoloadFunc()
	GenerateTheme()
	GenerateCacheTheme(usercolors []string, colorschememap map[string]string)
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
	GenerateDataTypeFile(datafile []string, datafilemap map[string]string)
	GenerateEnhanceplugin(plugins []string, enhancepluginmap map[string]string)
	GenerateSandWich(sandwich bool)
	GenerateTextObj()
	GenerateVersionControl(userversion []string, versionmap map[string]string)
	GeneratePluginFolder()
	GenerateLanguagePlugin(UserLanguages []string, LanguagesPluginMap map[string]string)
	GenerateCocJson()
	GenerateVimMap()
	GenerateInstallScripts()
}

// RollBack will remove the config folder when
// some errors disappear
func RollBack(err error) {
	color.PrintError(err.Error())
	color.PrintWarn("Rolling back...")
	os.RemoveAll(vim.ConfPath)
	os.Exit(0)
}

// ParseTemplate
func ParseTemplate(targetfile, name string, plugintemplate string, data interface{}) {
	f, err := os.OpenFile(targetfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		RollBack(err)
	}
	tpl := template.Must(template.New("").Parse(plugintemplate))
	err = tpl.Execute(f, data)
	if err != nil {
		RollBack(err)
	}
	color.PrintSuccess(fmt.Sprintf("Generate %v success", name))
}

// WriteTemplate
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

// WithConfirm
func WithConfirm(confirm bool, targetfile, name, plugintemplate string) {
	if confirm {
		WriteTemplate(targetfile, name, plugintemplate)
		return
	}
	color.PrintWarn(fmt.Sprintf("Skip generate %v cofnig", name))
}
