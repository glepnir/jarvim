// Package plugin provides ...
package plugin

const (
	DeinVimagt = `
[[plugins]]
repo = 'jreybert/vimagit'
on_cmd = 'Magit'
hook_source = '''
	autocmd User VimagitEnterCommit startinsert
'''
`
	DeinCommita = `
[[plugins]]
repo = 'rhysd/committia.vim'
on_path = [ 'COMMIT_EDITMSG', 'MERGE_MSG' ]
hook_source = '''
	let g:committia_min_window_width = 7
'''
`

	DeinFugiTive = `
[[plugins]]
repo = 'tpope/vim-fugitive'
on_cmd = [ 'G', 'Git', 'Gfetch', 'Gpush', 'Glog', 'Gclog', 'Gdiffsplit' ]
hook_source = '''
    augroup user_fugitive_plugin
      autocmd!
      autocmd FileType fugitiveblame normal A
    augroup END
'''
`

	DeinGina = `
[[plugins]]
repo = 'lambdalisue/gina.vim'
on_cmd = 'Gina'
`

	PlugVimagit = `
Plug 'jreybert/vimagit'
`

	PlugVimagitSetting = `
autocmd User VimagitEnterCommit startinsert
`

	PlugCommita = `
Plug 'rhysd/committia.vim'
`

	PlugFugTive = `
Plug 'tpope/vim-fugitive',{'on': [ 'G', 'Git', 'Gfetch', 'Gpush', 'Glog', 'Gclog', 'Gdiffsplit' ]}
`

	PlugFugTiveSetting = `
augroup user_fugitive_plugin
  autocmd!
  autocmd FileType fugitiveblame normal A
augroup END
`

	PlugGina = `
Plug 'lambdalisue/gina.vim',{'on': 'Gina'}
`

	PlugCommitaSetting = `
let g:committia_min_window_width = 7
`
)
