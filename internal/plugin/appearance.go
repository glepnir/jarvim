// Package plugin provides ...
package plugin

const (
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

	DeinDashboard = `
[[plugins]]
repo = 'hardcoreplayers/dashboard-nvim'
`

	DeinStatusline = `
[[plugins]]
repo = 'hardcoreplayers/spaceline.vim'
hook_source = '''
	let g:spaceline_seperate_style= 'slant'
'''
`

	DeinBufferLine = `
[[plugins]]
repo = 'hardcoreplayers/vim-buffet'
on_event = ['BufReadPre','BufNewFile']
`

	PlugColorscheme = `
{{range .}}
Plug '{{.}}'
{{end}}
`

	PlugDevicons = `
Plug 'ryanoasis/vim-devicons'
`
	PlugDashboard = `
Plug 'hardcoreplayers/dashboard-nvim'
`
	PlugBufferLine = `
Plug 'hardcoreplayers/vim-buffet'
`
	PlugStatusline = `
Plug 'hardcoreplayers/spaceline.vim'
`
	PlugStatuslineSetting = `
let g:spaceline_seperate_style= 'slant'
`
)
