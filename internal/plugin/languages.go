// Package plugin provides ...
package plugin

const (
	DeinCFamily = `
[[pluings]]
repo = 'jackguo380/vim-lsp-cxx-highlight'
on_ft = [ 'c','cpp' ]
hook_add = '''
	" ccls config
	if dein#tap('vim-lsp-cxx-highlight')
		call coc#config('languageserver', {
		\ 'ccls': {
			\ "command": "ccls",
			\ "rootPatterns": [".ccls", "compile_commands.json", ".git/", ".hg/"],
			\ "filetypes": ["c","cpp","objc","objcpp"],
			\ "initializationOptions": {
			\ "cache":{
				\ "directory": "/tmp/ccls"
			\ }
			\ }
			\ }
			\})
	endif
'''
`
	DeinR = `
[[plugins]]
repo = 'jalvesaq/Nvim-R'
on_ft = 'R'
hook_add = '''
	"R lsp config
	if dein#tap('Nvim-R')
		call coc#add_extension('coc-R')
	endif
'''
`
	DeinJavascript = `
[[plugins]]
repo = 'pangloss/vim-javascript'
on_ft = [ 'javascript', 'javascriptreact' ]
hook_add = '''
	let g:javascript_plugin_jsdoc = 1
    let g:javascript_plugin_flow = 1
	"javascript lsp config
	if dein#tap('vim-javascript')
		call coc#add_extension('coc-tsserver','coc-eslint','coc-prettier','coc-docthis')
	endif
'''

[[plugins]]
repo = 'moll/vim-node'
on_ft = [ 'javascript', 'javascriptreact' ]
`

	DeinTypescript = `
[[plugin]]
repo = 'HerringtonDarkholme/yats.vim'
on_ft = [ 'typescript', 'typescriptreact' ]
hook_add = '''
	if dein#tap('yats.vim')
		call coc#add_extension('coc-tsserver','coc-eslint', 'coc-prettier', 'coc-tslint-plugin' ,'coc-docthis')
	endif
'''
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
hook_add = '''
	if dein#tap('go-nvim')
		call coc#config('languageserver', {
		\ 'golang': {
			\ "command": "gopls",
			\ "rootPatterns": ["go.mod"],
			\ "disableWorkspaceFolders": "true",
			\ "filetypes": ["go"]
			\ }
			\})
	endif
'''
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
hook_add = '''
	"php lsp config
	if dein#tap('php.vim')
		call coc#config('languageserver', {
		\ 'intelephense': {
			\ "command": "intelephense",
			\ "args": ["--stdio"],
			\ "filetypes": ["php"],
			\ "initializationOptions": {
				\ "storagePath": "/tmp/intelephense"
			\ }
			\ }
			\})
	endif
'''
`
	DeinRuby = `
[[plugins]]
repo = 'vim-ruby/vim-ruby'
on_ft = 'ruby'
hook_add = '''
	"Ruby lsp config
	if dein#tap('vim-ruby')
		call coc#add_extension('coc-solargraph')
	endif
'''
`

	DeinScala = `
[[plugins]]
repo = 'derekwyatt/vim-scala'
on_ft = 'scala'
hook_add = '''
	"Scala lsp config
	if dein#tap('vim-scala')
		call coc#add_extension('coc-metals')
	endif
'''
`
	DeinShell = `
[[plugins]]
repo = 'arzg/vim-sh'
on_ft = [ 'sh','zsh' ]
hook_add = '''
	"shell lsp config
	if dein#tap('vim-sh')
		call coc#config('languageserver', {
		\ 'bash': {
			\ "command": "bash-language-server",
			\ "args" : ["start"],
			\ "ignoredRootPaths": ["~"],
			\ "filetypes": ["sh"]
			\ }
			\})
	endif
'''
`
	DeinLua = `
[[plugins]]
repo = 'tbastos/vim-lua'
on_ft = 'lua'
hook_add = '''
	if dein#tap(vim-lua)
		call coc#config ('languageserver', {
				\'lua': {
					\ "cwd": "full path of lua-language-server directory", (not sure this one is really necessary)
					\ "command": "full path to lua-language-server executable",
					\ "args": ["-E", "-e", "LANG=en", "[full path of lua-language-server directory]/main.lua"],
					\ "filetypes": ["lua"],
					\ "rootPatterns": [".git/"]
					\}
					\})
	endif
'''
`

	DeinPython = `
[[plugins]]
repo = 'vim-python/python-syntax'
on_ft = 'python'
hook_add = '''
	"python lsp config
	if dein#tap('python-syntax')
		let g:python_highlight_all = 1
		call coc#add_extension('coc-python')
	endif
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
	if dein#tap(html5.vim)
		call coc#add_extension('coc-html')
	endif
'''
`

	DeinCss = `
[[plugins]]
repo = 'hail2u/vim-css3-syntax'
on_ft = 'css'
hook_add = '''
	call coc#add_extension('coc-css')
'''
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
hook_add = '''
	call coc#add_extension('coc-flutter')
'''
`
)
