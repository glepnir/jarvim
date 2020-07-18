// Package render provides ...
package render

type Render interface {
	GenerateCore(Leader, LocalLeaderKey string)
	GenerateGeneral()
	GenerateTheme()
	GenerateCacheTheme(colorschemes []string)
	GenerateColorscheme(colorschemes []string)
	GenerateStatusLine()
}

func NewRender(p Render) Render {
	var r Render = p
	return r
}
