// Package plugin provides ...
package plugin

const (
	DeinCFamily = `
[[pluings]]
repo = 'jackguo380/vim-lsp-cxx-highlight'
on_ft = [ 'c','cpp' ]
`
	DeinR = `
[[plugins]]
repo = 'jalvesaq/Nvim-R'
on_ft = 'R'
`
	DeinJavascript = `
[[plugins]]
repo = 'pangloss/vim-javascript'
on_ft = [ 'javascript', 'javascriptreact' ]
hook_add = '''
	let g:javascript_plugin_jsdoc = 1
    let g:javascript_plugin_flow = 1
'''

[[plugins]]
repo = 'moll/vim-node'
on_ft = [ 'javascript', 'javascriptreact' ]
`

	DeinTypescript = `
[[plugin]]
repo = 'HerringtonDarkholme/yats.vim'
on_ft = [ 'typescript', 'typescriptreact' ]
`

	DeinReact = `
[[plugins]]
repo = 'MaxMEllon/vim-jsx-pretty'
on_ft = [ 'javascript', 'javascriptreact', 'typescriptreact' ]
depends = 'vim-javascript'
hook_add = '''
	let g:vim_jsx_pretty_colorful_config = 1
'''
`
	DeinVue = `
[[plugins]]
repo = 'posva/vim-vue'
on_ft = 'vue'
`
	DeinGo = `
[[plugins]]
repo = 'hardcoreplayers/go-nvim'
on_ft = ['go','gomod']
`

	DeinRust = `
[[plugins]]
repo = 'rust-lang/rust.vim'
on_ft = 'rust'
`
	DeinHaskell = `
[[plugins]]
repo = 'neovimhaskell/haskell-vim'
on_ft = 'haskell'
`

	DeinPhp = `
[[plugins]]
repo = 'StanAngeloff/php.vim'
on_ft = 'php'
`
	DeinRuby = `
[[plugins]]
repo = 'vim-ruby/vim-ruby'
on_ft = 'ruby'
`

	DeinScala = `
[[plugins]]
repo = 'derekwyatt/vim-scala'
on_ft = 'scala'
`
	DeinShell = `
[[plugins]]
repo = 'arzg/vim-sh'
on_ft = [ 'sh','zsh' ]
`
	DeinLua = `
[[plugins]]
repo = 'tbastos/vim-lua'
on_ft = 'lua'
`

	DeinPython = `
[[plugins]]
repo = 'vim-python/python-syntax'
on_ft = 'python'
hook_add = '''
	let g:python_highlight_all = 1
'''

[[plugins]]
repo = 'Vimjas/vim-python-pep8-indent'
on_ft = 'python'

[[plugins]]
repo = 'vim-scripts/python_match.vim'
on_ft = 'python'

[[plugins]]
repo = 'raimon49/requirements.txt.vim'
on_ft = 'requirements'
`

	DeinHtml = `
[[plugins]]
repo = othree/html5.vim
on_ft = 'html'
hook_add = '''
    let g:html5_event_handler_attributes_complete = 0
    let g:html5_rdfa_attributes_complete = 0
    let g:html5_microdata_attributes_complete = 0
    let g:html5_aria_attributes_complete = 0
'''
`

	DeinCss = `
[[plugins]]
repo = 'hail2u/vim-css3-syntax'
on_ft = 'css'
`
	DeinLess = `
repo = 'groenewege/vim-less'
on_ft = 'less'
`
	DeinSass = `
[[plugins]]
repo = 'cakebaker/scss-syntax.vim'
on_ft = [ 'scss', 'sass' ]
`
	DeinStylus = `
[[plugins]]
repo = 'iloginow/vim-stylus'
on_ft = 'stylus'
`

	DeinDart = `
[[plugins]]
repo = 'dart-lang/dart-vim-plugin'
on_ft = 'dart'
`
)
