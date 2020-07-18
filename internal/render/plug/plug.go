// Package plug provides ...
package plug

import "github.com/glepnir/jarvis/internal/render"

type Plug struct{}

var _ render.Render = (*Plug)(nil)

func (p *Plug) GenerateCore(Leaderkey, LocalLeaderKey string) {

}

func (p *Plug) GenerateGeneral() {

}

func (p *Plug) GenerateTheme() {

}

func (p *Plug) GenerateCacheTheme(colorschemes []string) {

}

func (p *Plug) GenerateColorscheme(colorschemes []string) {

}

func (p *Plug) GenerateStatusLine() {

}
