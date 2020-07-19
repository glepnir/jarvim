// Package plug provides ...
package vimplug

import (
	"os"

	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/color"
	"github.com/glepnir/jarvis/pkg/util"
)

type VimPlug struct{}

var _ render.Render = (*VimPlug)(nil)

func (v *VimPlug) GenerateInit() {
}

func (v *VimPlug) GenerateCore(Leaderkey, LocalLeaderKey string, leaderkeymap map[string]string) {

}

func (v *VimPlug) GeneratePlugMan() {}

func (v *VimPlug) GenerateGeneral() {

}

func (v *VimPlug) GenerateAutoloadFunc() {

}

func (v *VimPlug) GenerateTheme() {

}

func (v *VimPlug) GenerateCacheTheme(usercolors []string, colorschememap map[string]string) {

}

func (v *VimPlug) GenerateColorscheme(colorschemes []string) {

}

func (v *VimPlug) GenerateDevIcons() {

}

func (v *VimPlug) GenerateDashboard(dashboard bool) {}

func (v *VimPlug) GenerateBufferLine(bufferline bool) {

}

func (v *VimPlug) GenerateStatusLine(statusline bool) {

}

func (v *VimPlug) GenerateExplorer(explorer string) {

}

func (v *VimPlug) GenerateDatabase(database bool) {

}

func (v *VimPlug) GenerateFuzzyFind(fuzzyfind bool) {

}

func (v *VimPlug) GenerateIndentLine(indentplugin string) {

}

func (v *VimPlug) GenerateComment(comment bool) {

}

func (v *VimPlug) GenerateOutLine(outline bool) {

}

func (v *VimPlug) GenerateTags(tagsplugin bool) {

}

func (v *VimPlug) GenerateQuickRun(quickrun bool) {

}

func (v *VimPlug) GenerateEditorConfig(editorconfig bool) {
}

func (v *VimPlug) GenerateDataTypeFile(datafile []string, datafilemap map[string]string) {
}

func (v *VimPlug) GenerateEnhanceplugin(plugins []string, enhancepluginmap map[string]string) {}

func (v *VimPlug) GenerateSandWich(sandiwch bool) {

}

func (v *VimPlug) GenerateTextObj() {

}

func (v *VimPlug) GenerateVersionControl(userversion []string, versionmap map[string]string) {}

func (v *VimPlug) GeneratePluginFolder() {
	err := util.CopyDir("./../../plugin", vim.ConfPlugin)
	if err != nil {
		color.PrintError("Copy plugin folder to your vim config path failed")
		os.Exit(0)
	}
	color.PrintSuccess("Generate plugin folder success")
}
