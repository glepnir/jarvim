// Package plugin provides ...
package plugin

const DeinIndentLine = `
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
const DeinIndenGuides = `
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
