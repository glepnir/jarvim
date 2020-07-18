// Package plugin provides ...
package plugin

const DeinContextFileType = `
[[plugins]]
repo = 'Shougo/context_filetype.vim'
`

const DeinMarkDown = `
[[plugins]]
repo = 'plasticboy/vim-markdown'
on_ft = 'markdown'
hook_add = '''
    let g:vim_markdown_folding_level = 1
    let g:vim_markdown_folding_style_pythonic = 1
    let g:vim_markdown_frontmatter = 1
    let g:vim_markdown_auto_insert_bullets = 1
    let g:vim_markdown_new_list_item_indent = 0
    let g:vim_markdown_conceal_code_blocks = 0
    let g:vim_markdown_conceal = 0
    let g:vim_markdown_strikethrough = 1
    let g:vim_markdown_edit_url_in = 'vsplit'
    let g:vim_markdown_fenced_languages = [
      \ 'c++=cpp',
      \ 'viml=vim',
      \ 'bash=sh',
      \ 'ini=dosini',
      \ 'js=javascript',
      \ 'json=javascript',
      \ 'jsx=javascriptreact',
      \ 'tsx=typescriptreact',
      \ 'docker=Dockerfile',
      \ 'makefile=make',
      \ 'py=python'
      \ ]
'''

[[plugins]]
repo = 'iamcco/markdown-preview.nvim'
on_ft = ['markdown', 'pandoc.markdown', 'rmd']
build = 'sh -c "cd app & yarn install"'
hook_source = '''
    let g:mkdp_auto_start = 0
'''
`

const DeinToml = `
[[plugins]]
repo = 'cespare/vim-toml'
on_ft = 'toml'
`
const DeinNginx = `
[[plugins]]
repo = 'chr4/nginx.vim'
on_ft = 'nginx'
`

const DeinJson = `
[[plugins]]
repo = 'kevinoid/vim-jsonc'
on_ft = 'json'
`

const DeinDockerFile = `
[[plugins]]
repo = 'ekalinin/Dockerfile.vim'
on_ft = [ 'Dockerfile', 'docker-compose' ]
`
