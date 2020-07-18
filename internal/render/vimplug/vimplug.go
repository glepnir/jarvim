// Package plug provides ...
package vimplug

import "github.com/glepnir/jarvis/internal/render"

type VimPlug struct{}

var _ render.Render = (*VimPlug)(nil)

func (v *VimPlug) GenerateCore(Leaderkey, LocalLeaderKey string) {

}

func (v *VimPlug) GenerateGeneral() {

}

func (v *VimPlug) GenerateTheme() {

}

func (v *VimPlug) GenerateCacheTheme(colorschemes []string) {

}

func (v *VimPlug) GenerateColorscheme(colorschemes []string) {

}

func (v *VimPlug) GenerateDevIcons() {

}

func (v *VimPlug) GenerateBufferLine() {

}

func (v *VimPlug) GenerateStatusLine() {

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
