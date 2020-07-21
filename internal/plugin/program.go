// Package plugin provides ...
package plugin

const (
	DeinEditorConfig = `
[[plugins]]
repo = 'editorconfig/editorconfig-vim'
`

	DeinIndentLine = `
[[plugins]]
repo = 'Yggdroot/indentLine'
on_event = 'BufReadPre'
hook_source = '''
    let g:indentLine_enabled = 1
    let g:indentLine_char='┆'
    let g:indentLine_fileTypeExclude = ['defx', 'denite','startify','tagbar','vista_kind','vista','coc-explorer']
    let g:indentLine_concealcursor = 'niv'
    let g:indentLine_showFirstIndentLevel =1
'''
`
	DeinIndenGuides = `
[[plugins]]
repo = 'nathanaelkane/vim-indent-guides'
on_event = 'FileType'
hook_post_source: IndentGuidesEnable
hook_add = '''
    let g:indent_guides_default_mapping = 0
    let g:indent_guides_tab_guides = 0
    let g:indent_guides_color_change_percent = 3
    let g:indent_guides_guide_size = 1
    let g:indent_guides_exclude_filetypes = [
      \ 'help', 'denite', 'denite-filter', 'startify',
      \ 'vista', 'vista_kind', 'tagbar', 'nerdtree',
      \ 'lsp-hover', 'clap_input'
      \ ]
'''
`

	DeinCaw = `
[[plugins]]
repo = 'tyru/caw.vim'
depends = 'context_filetype.vim'
on_map = { nx = '<Plug>' }
`

	DeinVista = `
[[plugins]]
repo = 'liuchengxu/vista.vim'
on_cmd = ['Vista', 'Vista!', 'Vista!!']
hook_source = '''
    let g:vista#renderer#enable_icon = 1
    let g:vista_disable_statusline = 1
    let g:vista_default_executive = 'ctags'
    let g:vista_echo_cursor_strategy = 'floating_win'
    let g:vista_vimwiki_executive = 'markdown'
    let g:vista_executive_for = {
      \ 'vimwiki': 'markdown',
      \ 'pandoc': 'markdown',
      \ 'markdown': 'toc',
      \ 'yaml': 'coc',
      \ 'typescript': 'coc',
      \ 'typescriptreact': 'coc',
      \ }
'''
`

	DeinGuTenTags = `
[[plugins]]
repo = 'ludovicchabant/vim-gutentags'
if = 'executable("ctags")'
on_event = ['BufReadPost', 'BufWritePost']
hook_source = '''
    let g:gutentags_cache_dir = $DATA_PATH . '/tags'
    let g:gutentags_project_root = ['.root', '.git', '.svn', '.hg', '.project','go.mod','/usr/local']
    let g:gutentags_generate_on_write = 1
    let g:gutentags_generate_on_missing = 1
    let g:gutentags_generate_on_new = 0
    let g:gutentags_exclude_filetypes = [ 'defx', 'denite', 'vista', 'magit' ]
    let g:gutentags_ctags_extra_args = ['--output-format=e-ctags']
    let g:gutentags_ctags_exclude = ['*.json', '*.js', '*.ts', '*.jsx', '*.css', '*.less', '*.sass', '*.go', '*.dart', 'node_modules', 'dist', 'vendor']
'''
`

	DeinQuickRun = `
[[plugins]]
repo = 'thinca/vim-quickrun'
on_cmd = 'QuickRun'
hook_source = '''
    let g:quickrun_config = {
      \   "_" : {
        \       "outputter" : "message",
        \   },
        \}
    let g:quickrun_no_default_key_mappings = 1
'''
`

	DeinEmmet = `
[[plugins]]
repo = 'mattn/emmet-vim'
on_event = 'InsertEnter'
on_ft = [ 'html','css','javascript','javascriptreact','vue','typescript','typescriptreact' ]
hook_source = '''
    let g:user_emmet_complete_tag = 0
    let g:user_emmet_install_global = 0
    let g:user_emmet_install_command = 0
    let g:user_emmet_mode = 'i'
'''
`
	PlugIndentLine = `
Plug 'Yggdroot/indentLine'
`
	PlugIndentGuides = `
Plug 'nathanaelkane/vim-indent-guides'
`

	PlugIndentLineSetting = `
let g:indentLine_enabled = 1
let g:indentLine_char='┆'
let g:indentLine_fileTypeExclude = ['defx', 'denite','startify','tagbar','vista_kind','vista','coc-explorer']
let g:indentLine_concealcursor = 'niv'
let g:indentLine_showFirstIndentLevel =1
`
	PlugIndentGuidesSetting = `
let g:indent_guides_default_mapping = 0
let g:indent_guides_tab_guides = 0
let g:indent_guides_color_change_percent = 3
let g:indent_guides_guide_size = 1
let g:indent_guides_exclude_filetypes = [
  \ 'help', 'denite', 'denite-filter', 'startify',
  \ 'vista', 'vista_kind', 'tagbar', 'nerdtree',
  \ 'lsp-hover', 'clap_input'
  \ ]
`
	PlugCawvim = `
Plug 'tyru/caw.vim'
`
	PlugVista = `
Plug 'liuchengxu/vista.vim',{'on': ['Vista', 'Vista!', 'Vista!!']}
`
	PlugVistaSetting = `
let g:vista#renderer#enable_icon = 1
let g:vista_disable_statusline = 1
let g:vista_default_executive = 'ctags'
let g:vista_echo_cursor_strategy = 'floating_win'
let g:vista_vimwiki_executive = 'markdown'
let g:vista_executive_for = {
  \ 'vimwiki': 'markdown',
  \ 'pandoc': 'markdown',
  \ 'markdown': 'toc',
  \ 'yaml': 'coc',
  \ 'typescript': 'coc',
  \ 'typescriptreact': 'coc',
  \ }
`
	PlugGuTenTags = `
if executable('ctags')
	Plug 'ludovicchabant/vim-gutentags'
endif
`
	PlugGuTenTagsSetting = `
let g:gutentags_cache_dir = $DATA_PATH . '/tags'
let g:gutentags_project_root = ['.root', '.git', '.svn', '.hg', '.project','go.mod','/usr/local']
let g:gutentags_generate_on_write = 1
let g:gutentags_generate_on_missing = 1
let g:gutentags_generate_on_new = 0
let g:gutentags_exclude_filetypes = [ 'defx', 'denite', 'vista', 'magit' ]
let g:gutentags_ctags_extra_args = ['--output-format=e-ctags']
let g:gutentags_ctags_exclude = ['*.json', '*.js', '*.ts', '*.jsx', '*.css', '*.less', '*.sass', '*.go', '*.dart', 'node_modules', 'dist', 'vendor']
`
	PlugQuickRun = `
Plug 'thinca/vim-quickrun',{'on': 'QuickRun'}
`
	PlugQuickRunSetting = `
let g:quickrun_config = {
  \   "_" : {
    \       "outputter" : "message",
    \   },
    \}
let g:quickrun_no_default_key_mappings = 1
`
	PlugEditorConfig = `
Plug 'editorconfig/editorconfig-vim'
`

	PlugEmmet = `
Plug 'mattn/emmet-vim' ,{'for': [ 'html','css','javascript','javascriptreact','vue','typescript','typescriptreact' ]}
`

	PlugEmmetSetting = `
let g:user_emmet_complete_tag = 0
let g:user_emmet_install_global = 0
let g:user_emmet_install_command = 0
let g:user_emmet_mode = 'i'
`
)
