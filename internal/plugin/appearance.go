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
repo = 'glepnir/dashboard-nvim'
`
	// DeinStatusline plugin
	DeinStatusline = `
[[plugins]]
repo = 'glepnir/spaceline.vim'
hook_source = '''
	let g:spaceline_seperate_style= 'slant'
'''
`
	// DeinBufferLine plugin
	DeinBufferLine = `
[[plugins]]
repo = 'romgrk/barbar.nvim'
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
Plug 'glepnir/dashboard-nvim'
`
	//PlugBufferLine
	PlugBufferLine = `
Plug 'romgrk/barbar.nvim'
`
	//PlugStatusline
	PlugStatusline = `
Plug 'glepnir/spaceline.vim'
`
	// PlugStatuslineSetting
	PlugStatuslineSetting = `
let g:spaceline_seperate_style= 'slant'
`
)
