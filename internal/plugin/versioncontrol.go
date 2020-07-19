// Package plugin provides ...
package plugin

const DeinVimagt = `
[[plugins]]
repo = 'jreybert/vimagit'
on_cmd = 'Magit'
hook_source = '''
	autocmd User VimagitEnterCommit startinsert
'''
`
const DeinCommita = `
[[plugins]]
repo = 'rhysd/committia.vim'
on_path = [ 'COMMIT_EDITMSG', 'MERGE_MSG' ]
hook_source = '''
	let g:committia_min_window_width = 7
'''
`

const DeinFugiTive = `
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

const DeinGina = `
[[plugins]]
repo = 'lambdalisue/gina.vim'
on_cmd = 'Gina'
`
