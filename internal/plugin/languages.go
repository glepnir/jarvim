// Package plugin provides ...
package plugin

const (
	DeinCFamily = `
[[pluings]]
repo = 'jackguo380/vim-lsp-cxx-highlight'
on_ft = [ 'c','cpp' ]
hook_source = '''
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
'''
`
	DeinR = `
[[plugins]]
repo = 'jalvesaq/Nvim-R'
on_ft = 'R'
`

	DeinRExtension = "coc-R"

	DeinJavascript = `
[[plugins]]
repo = 'pangloss/vim-javascript'
on_ft = [ 'javascript', 'javascriptreact' ]
hook_source = '''
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
	DeinVueExtension = "coc-vetur"

	DeinGo = `
[[plugins]]
repo = 'hardcoreplayers/go-nvim'
on_ft = ['go','gomod']
hook_add = '''
	nnoremap gcg :GoAuToComment<CR>
'''
hook_source = '''
	call coc#config('languageserver', {
	\ 'golang': {
		\ "command": "gopls",
		\ "rootPatterns": ["go.mod"],
		\ "disableWorkspaceFolders": "true",
		\ "filetypes": ["go"]
		\ }
		\})
'''
`

	DeinRust = `
[[plugins]]
repo = 'rust-lang/rust.vim'
on_ft = 'rust'
`
	DeinRustExtension = "coc-rust-analyzer"

	DeinHaskell = `
[[plugins]]
repo = 'neovimhaskell/haskell-vim'
on_ft = 'haskell'
hook_source = '''
	call coc#config('languageserver', {
		\ 'haskell': {
		\ "command": "hie-wrapper",
		\ "rootPatterns": [".stack.yaml","cabal.config","package.yaml"],
		\ "filetypes": ["hs","lhs","haskell"],
		\ "initializationOptions":{},
		\ "settings":{
			\ "languageServerHaskell":{
			\ "hlintOn":"true",
			\ "maxNumberOfProblems":10,
			\ "completionSnippetsOn": "true"
		\ }
		\ }
		\ }
		\})
'''
`

	DeinPhp = `
[[plugins]]
repo = 'StanAngeloff/php.vim'
on_ft = 'php'
hook_source = '''
	"php lsp config
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
'''
`
	DeinRuby = `
[[plugins]]
repo = 'vim-ruby/vim-ruby'
on_ft = 'ruby'
`
	DeinRubyExtension = "coc-solargraph"

	DeinScala = `
[[plugins]]
repo = 'derekwyatt/vim-scala'
on_ft = 'scala'
`
	DeinScalaExtension = "coc-metals"

	DeinShell = `
[[plugins]]
repo = 'arzg/vim-sh'
on_ft = [ 'sh','zsh' ]
hook_source = '''
	call coc#config('languageserver', {
	\ 'bash': {
		\ "command": "bash-language-server",
		\ "args" : ["start"],
		\ "ignoredRootPaths": ["~"],
		\ "filetypes": ["sh"]
		\ }
		\})
'''
`
	DeinLua = `
[[plugins]]
repo = 'tbastos/vim-lua'
on_ft = 'lua'
hook_source = '''
	call coc#config ('languageserver', {
			\'lua': {
				\ "cwd": "full path of lua-language-server directory", (not sure this one is really necessary)
				\ "command": "full path to lua-language-server executable",
				\ "args": ["-E", "-e", "LANG=en", "[full path of lua-language-server directory]/main.lua"],
				\ "filetypes": ["lua"],
				\ "rootPatterns": [".git/"]
				\}
				\})
'''
`

	DeinPythonExtension = "coc-python"

	DeinPython = `
[[plugins]]
repo = 'vim-python/python-syntax'
on_ft = 'python'
hook_source = '''
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
repo = 'othree/html5.vim'
on_ft = 'html'
hook_source = '''
    let g:html5_event_handler_attributes_complete = 0
    let g:html5_rdfa_attributes_complete = 0
    let g:html5_microdata_attributes_complete = 0
    let g:html5_aria_attributes_complete = 0
'''
`
	DeinHtmlExtension = "coc-html"

	DeinCss = `
[[plugins]]
repo = 'hail2u/vim-css3-syntax'
on_ft = 'css'
`
	DeinCssExtension = "coc-css"

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
hook_source = '''
	call coc#add_extension('coc-flutter')
'''
`
	PlugCFamily = `
Plug 'jackguo380/vim-lsp-cxx-highlight',{'for' : [ 'c','cpp' ]}
`
	PlugCFamilyLsp = `
" ccls config
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
'''
`
	PlugR = `
Plug 'jalvesaq/Nvim-R' ,{'for': 'R'}
`
	PlugRLsp = `
call coc#add_extension('coc-R')
`
	PlugJavascript = `
Plug 'pangloss/vim-javascript', {'for': [ 'javascript', 'javascriptreact' ]}
Plug 'moll/vim-node', {'for': [ 'javascript', 'javascriptreact' ]}
`
	PlugJavascriptLsp = `
let g:javascript_plugin_jsdoc = 1
let g:javascript_plugin_flow = 1
"javascript lsp config
call coc#add_extension('coc-tsserver','coc-eslint','coc-prettier','coc-docthis')
`

	PlugTypescript = `
Plug  'HerringtonDarkholme/yats.vim' , {'for' : [ 'typescript', 'typescriptreact' ]}
`
	PlugTypescriptLsp = `
call coc#add_extension('coc-tsserver','coc-eslint', 'coc-prettier', 'coc-tslint-plugin' ,'coc-docthis')
`

	PlugReact = `
Plug 'MaxMEllon/vim-jsx-pretty', {'for': [ 'javascript', 'javascriptreact', 'typescriptreact' ]}
`
	PlugReactLsp = `
call coc#add_extension('coc-style-helper','coc-react-refactor')
`
	PlugVue = `
Plug 'posva/vim-vue' , {'for': 'vue'}
`

	PlugVueLsp = `
call coc#add_extension('coc-vetur')
`
	PlugGo = `
Plug  'hardcoreplayers/go-nvim' , {'for' : ['go','gomod']}
`
	PlugGoLsp = `
" Auto generate the comment for function or method
nnoremap gcg :GoAuToComment<CR>

call coc#config('languageserver', {
\ 'golang': {
	\ "command": "gopls",
	\ "rootPatterns": ["go.mod"],
	\ "disableWorkspaceFolders": "true",
	\ "filetypes": ["go"]
	\ }
	\})
`

	PlugRust = `
Plug 'rust-lang/rust.vim' ,{'for' : 'rust'}
`
	PlugRustLsp = `
call coc#add_extension('coc-rust-analyzer')
`
	PlugHaskell = `
Plug 'neovimhaskell/haskell-vim' , {'for': 'haskell'}
`

	PlugHaskellLsp = `
call coc#config('languageserver', {
    \ 'haskell': {
      \ "command": "hie-wrapper",
      \ "rootPatterns": [".stack.yaml","cabal.config","package.yaml"],
      \ "filetypes": ["hs","lhs","haskell"],
      \ "initializationOptions":{},
      \ "settings":{
        \ "languageServerHaskell":{
          \ "hlintOn":"true",
          \ "maxNumberOfProblems":10,
          \ "completionSnippetsOn": "true"
      \ }
      \ }
      \ }
      \})
`

	PlugPhp = `
Plug 'StanAngeloff/php.vim' , {'for': 'php'}
`
	PlugPhpLsp = `
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
`
	PlugRuby = `
Plug 'vim-ruby/vim-ruby' ,{'for' : 'ruby'}
`
	PlugRubyLsp = `
call coc#add_extension('coc-solargraph')
`

	PlugScala = `
Plug 'derekwyatt/vim-scala',{'for': 'scala'}
`
	PlugScalaLsp = `
call coc#add_extension('coc-metals')
`
	PlugShell = `
Plug 'arzg/vim-sh', {'for': [ 'sh','zsh' ]}
`
	PlugShellLsp = `
call coc#config('languageserver', {
\ 'bash': {
	\ "command": "bash-language-server",
	\ "args" : ["start"],
	\ "ignoredRootPaths": ["~"],
	\ "filetypes": ["sh"]
	\ }
	\})
`
	PlugLua = `
Plug 'tbastos/vim-lua',{'for': 'lua'}
`
	PlugLuaLsp = `
call coc#config ('languageserver', {
		\'lua': {
			\ "cwd": "full path of lua-language-server directory", (not sure this one is really necessary)
			\ "command": "full path to lua-language-server executable",
			\ "args": ["-E", "-e", "LANG=en", "[full path of lua-language-server directory]/main.lua"],
			\ "filetypes": ["lua"],
			\ "rootPatterns": [".git/"]
			\}
			\})
`

	PlugPython = `
Plug 'vim-python/python-syntax' ,{'for': 'python'}
Plug 'Vimjas/vim-python-pep8-indent',{'for': 'python'}
Plug 'vim-scripts/python_match.vim',{'for': 'python'}
Plug 'raimon49/requirements.txt.vim',{'for': 'python'}
`
	PlugPythonLsp = `
let g:python_highlight_all = 1
call coc#add_extension('coc-python')
`

	PlugHtml = `
Plug 'othree/html5.vim', {'for': 'html'}
`

	PlugHtmlLsp = `
let g:html5_event_handler_attributes_complete = 0
let g:html5_rdfa_attributes_complete = 0
let g:html5_microdata_attributes_complete = 0
let g:html5_aria_attributes_complete = 0

call coc#add_extension('coc-html')
`

	PlugCss = `
Plug 'hail2u/vim-css3-syntax' ,{'for': 'css'}
`
	PlugCssLsp = `
call coc#add_extension('coc-css')
`
	PlugLess = `
Plug 'groenewege/vim-less' ,{'for' : 'less'}
`
	PlugSass = `
Plug 'cakebaker/scss-syntax.vim' ,{'for' : [ 'scss', 'sass' ]}
`
	PlugStylus = `
Plug 'iloginow/vim-stylus' , {'for' : 'stylus'}
`

	PlugDart = `
Plug  'dart-lang/dart-vim-plugin' ,{'for': 'dart'}
`

	PlugDartLsp = `
call coc#add_extension('coc-flutter')
`
)

var JsTsExtensions = []string{
	"coc-tsserver",
	"coc-eslint",
	"coc-prettier",
	"coc-tslint-plugin",
	"coc-docthis",
}

var ReactExtensions = []string{
	"coc-style-helper",
	"coc-react-refactor",
}
