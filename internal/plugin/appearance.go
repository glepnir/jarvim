// Package plugin provides ...
package plugin

const (
	// DeinDevicons plugin
	DeinDevicons = `
[[plugins]]
repo = 'ryanoasis/vim-devicons'
`
	DeinColorscheme = `
{{range .}}
[[plugins]]
repo = '{{.}}'
{{end}}
`

	// DeinDashboard plugin
	DeinDashboard = `
[[plugins]]
repo = 'hardcoreplayers/dashboard-nvim'
`
	// DeinStatusline plugin
	DeinStatusline = `
[[plugins]]
repo = 'hardcoreplayers/spaceline.vim'
hook_source = '''
	let g:spaceline_seperate_style= 'slant'
'''
`
	// DeinBufferLine plugin
	DeinBufferLine = `
[[plugins]]
repo = 'hardcoreplayers/vim-buffet'
on_event = ['BufReadPre','BufNewFile']
`
	// PlugColorscheme
	PlugColorscheme = `
{{range .}}
Plug '{{.}}'
{{end}}
`
	// PlugDevicons
	PlugDevicons = `
Plug 'ryanoasis/vim-devicons'
`
	// PlugDashboard
	PlugDashboard = `
Plug 'hardcoreplayers/dashboard-nvim'
`
	//PlugBufferLine
	PlugBufferLine = `
Plug 'hardcoreplayers/vim-buffet'
`
	//PlugStatusline
	PlugStatusline = `
Plug 'hardcoreplayers/spaceline.vim'
`
	// PlugStatuslineSetting
	PlugStatuslineSetting = `
let g:spaceline_seperate_style= 'slant'
`
)
