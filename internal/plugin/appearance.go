// Package plugin provides ...
package plugin

const DeinDevicons = `
[[plugins]]
repo = 'ryanoasis/vim-devicons'
`
const DeinColorscheme = `
{{range .}}
[[plugins]]
repo = '{{.}}'
{{end}}
`
const DeinStatusline = `
[[plugins]]
repo = 'hardcoreplayers/spaceline.vim'
hook_source = '''
	let g:spaceline_seperate_style= 'slant
'''
`

const DeinBufferLine = `
[[plugins]]
repo = 'hardcoreplayers/vim-buffet'
on_event = ['BufReadPre','BufNewFile']
`
