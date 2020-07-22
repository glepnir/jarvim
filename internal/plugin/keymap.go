// Package plugin provides ...
package plugin

const VimKeyMap = `
" See plugin/bufkill.vim
" use emacs keybind
nmap <C-x>k :BD<CR>
"Write buffer (save)
nnoremap <C-s> :<C-u>write<CR>
"yank to end
nnoremap Y y$
" Whitespace jump (see plugin/whitespace.vim)
nnoremap ]w :<C-u>WhitespaceNext<CR>
nnoremap [w :<C-u>WhitespacePrev<CR>
" Remove spaces at the end of lines
nnoremap <silent> <Space>cw :<C-u>silent! keeppatterns %substitute/\s\+$//e<CR>

" insert keymap like emacs
inoremap <C-w> <C-[>diwa
inoremap <C-h> <BS>
inoremap <C-d> <Del>
inoremap <C-k>  <ESC>d$a
inoremap <C-u> <C-G>u<C-U>
inoremap <C-b> <Left>
inoremap <C-f> <Right>
inoremap <C-a> <ESC>^i
inoremap <expr><C-e> pumvisible() ? "\<C-e>" : "\<End>"
"insert a newline
inoremap <C-O> <Esc>o
imap <C-S> <esc>:w<CR>
imap <C-Q> <esc>:wq<CR>

" command line alias
cnoremap <C-p> <Up>
cnoremap <C-b> <Left>
cnoremap <C-f> <Right>
cnoremap <C-a> <Home>
cnoremap <C-e> <End>
cnoremap <C-d> <Del>
cnoremap <C-h> <BS>
cnoremap <C-t> <C-R>=expand("%:p:h") . "/" <CR>

tnoremap <A-[> <C-\><C-n>

function! s:QuiteVim() abort
  if empty(expand('%:t'))
    execute ":q"
  else
    execute ":wq!"
  endif
endfunction
" Quiet
nnoremap <leader>qq :call <SID>QuiteVim()<CR>
nnoremap <leader>qw <esc>:q!<CR>

"switch window
nnoremap <C-h> <C-w>h
nnoremap <C-l> <C-w>l
nnoremap <C-j> <C-w>j
nnoremap <C-k> <C-w>k

" settings for resize splitted window
nmap <C-w>[ :vertical resize -3<CR>
nmap <C-w>] :vertical resize +3<CR>

" Session management shortcuts (see plugin/sessions.vim)
nmap <Leader>ss :<C-u>SessionSave<CR>
nmap <Leader>sl :<C-u>SessionLoad<CR>
`

const BuffetKeyMap = `
"--------------------------"
"     vim-buffet Keymap    "
"--------------------------"
nnoremap  ]b :<C-u>bp<CR>
nnoremap  [b :<C-u>bn<CR>
nnoremap <silent> <Leader>bc :Bonly<CR>
nnoremap <silent> <Leader>bx :Bw<CR>
nmap <leader>1 <Plug>BuffetSwitch(1)
nmap <leader>2 <Plug>BuffetSwitch(2)
nmap <leader>3 <Plug>BuffetSwitch(3)
nmap <leader>4 <Plug>BuffetSwitch(4)
nmap <leader>5 <Plug>BuffetSwitch(5)
nmap <leader>6 <Plug>BuffetSwitch(6)
nmap <leader>7 <Plug>BuffetSwitch(7)
nmap <leader>8 <Plug>BuffetSwitch(8)
nmap <leader>9 <Plug>BuffetSwitch(9)
nmap <leader>0 <Plug>BuffetSwitch(10)
`
const DefxKeyMap = `
nnoremap <silent> <Leader>e
  \ :<C-u>Defx -resume -toggle -buffer-name=tab` + "`tabpagenr()`<CR>"
const DefxFindKeyMap = `
nnoremap <silent> <Leader>F
  \ :<C-u>Defx -resume -buffer-name=tab` + "tabpagenr()` -search=`" + "expand('%:p')`<CR>"

const MarkdwonPreviewKeyMap = `
nnoremap <silent> <Leader>om  :<C-u>MarkdownPreview<CR>
`
const FloatermKeyMap = `
nnoremap <silent> <Leader>t :<C-u>FloatermToggle<CR>
`

const DataBaseKeyMap = `
nnoremap <silent> <Leader>od :DBUIToggle<CR>
`

const CocClapKeyMap = `
"--------------------------"
"     coc-clap Keymap      "
"--------------------------"
" Show all diagnostics
nnoremap <silent> <Leader>ce  :Clap coc_diagnostics<CR>
" Manage extensions
nnoremap <silent> <Leader>;   :Clap coc_extensions<CR>
" Show commands
nnoremap <silent> <Leader>,   :Clap coc_commands<CR>
" Search workspace symbols
nnoremap <silent> <Leader>cs  :Clap coc_symbols<CR>
nnoremap <silent> <Leader>cS  :Clap coc_services<CR>
nnoremap <silent> <leader>ct  :Clap coc_outline<CR>
`

const CocKeyMap = `
"--------------------------"
"     Coc Keymap           "
"--------------------------"
" Remap for do codeAction of selected region
function! s:cocActionsOpenFromSelected(type) abort
    execute 'CocCommand actions.open ' . a:type
endfunction
xmap <silent> <Leader>a :<C-u>execute 'CocCommand actions.open ' . visualmode()<CR>
nmap <silent> <Leader>a :<C-u>set operatorfunc=<SID>cocActionsOpenFromSelected<CR>g@
" Do default action for next item.
nmap <silent> [a  :<C-u>CocNext<CR>
" Do default action for previous item.
nmap <silent> ]a  :<C-u>CocPrev<CR>
" Use [e and ]e for navigate diagnostics
nmap <silent> ]e <Plug>(coc-diagnostic-prev)
nmap <silent> [e <Plug>(coc-diagnostic-next)
" Remap for rename current word
nmap <Leader>cn <Plug>(coc-rename)
" Remap for format selected region
vmap <Leader>cf  <Plug>(coc-format-selected)
nmap <Leader>cf  <Plug>(coc-format-selected)
" Fix autofix problem of current line
nmap <Leader>cF  <Plug>(coc-fix-current)
" Remap keys for gotos
nmap <silent> gd :<C-u>call initself#definition_other_window()<CR>
nmap <silent> gy <Plug>(coc-type-definition)
nmap <silent> <Leader>ci <Plug>(coc-implementation)
nmap <silent> gr <Plug>(coc-references)
" Use K for show documentation in float window
nnoremap <silent> K :call CocActionAsync('doHover')<CR>
" use <c-space> for trigger completion.
inoremap <silent><expr> <c-space> coc#refresh()
nmap ]g <Plug>(coc-git-prevchunk)
nmap [g <Plug>(coc-git-nextchunk)
" show chunk diff at current position
nmap <Leader>gi <Plug>(coc-git-chunkinfo)
" show commit contains current position
nmap <Leader>gm <Plug>(coc-git-commit)
" float window scroll
nnoremap <expr><C-f> coc#util#has_float() ? coc#util#float_scroll(1) : "\<C-f>"
nnoremap <expr><C-b> coc#util#has_float() ? coc#util#float_scroll(0) : "\<C-b>"
" Use <TAB> for selections ranges.
" NOTE: Requires 'textDocument/selectionRange' support from the language server.
" coc-tsserver, coc-python are the examples of servers that support it.
nmap <silent> <TAB> <Plug>(coc-range-select)
xmap <silent> <TAB> <Plug>(coc-range-select)
" Add :OR command for organize imports of the current buffer.
command! -nargs=0 OR  :call CocAction('runCommand', 'editor.action.organizeImport')
nnoremap <silent> <Leader>co :<C-u>OR<CR>
" multiple cursors
nmap <silent><M-s> <Plug>(coc-cursors-position)
nmap <expr> <silent><M-d> initself#select_current_word()
xmap <silent><M-d> <Plug>(coc-cursors-range)
nmap <silent><M-c>  <Plug>(coc-cursors-operator)

" Use :Format for format current buffer
command! -nargs=0 Format :call CocAction('format')

nnoremap  <Leader>fz :<C-u>CocSearch -w<Space>
" Introduce function text object
" NOTE: Requires 'textDocument.documentSymbol' support from the language server.
xmap if <Plug>(coc-funcobj-i)
xmap af <Plug>(coc-funcobj-a)
omap if <Plug>(coc-funcobj-i)
omap af <Plug>(coc-funcobj-a)
nmap gcj :execute 'CocCommand docthis.documentThis'<CR>
`

const ClapKeyMap = `
"--------------------------"
"     vim-clap Keymap      "
"--------------------------"
nnoremap <silent> <Leader>tc :<C-u>Clap colors<CR>
nnoremap <silent> <Leader>bb :<C-u>Clap buffers<CR>
nnoremap <silent> <Leader>fa :<C-u>Clap grep2<CR>
nnoremap <silent> <Leader>fb :<C-u>Clap marks<CR>
"like emacs counsel-find-file
nnoremap <silent> <C-x><C-f> :<C-u>Clap filer<CR>
nnoremap <silent> <Leader>ff :<C-u>Clap files ++finder=rg --ignore --hidden --files<cr>
nnoremap <silent> <Leader>fg :<C-u>Clap gfiles<CR>
nnoremap <silent> <Leader>fw :<C-u>Clap grep ++query=<cword><cr>
nnoremap <silent> <Leader>fh :<C-u>Clap history<CR>
nnoremap <silent> <Leader>fW :<C-u>Clap windows<CR>
nnoremap <silent> <Leader>fl :<C-u>Clap loclist<CR>
nnoremap <silent> <Leader>fu :<C-u>Clap git_diff_files<CR>
nnoremap <silent> <Leader>fv :<C-u>Clap grep ++query=@visual<CR>
nnoremap <silent> <Leader>oc :<C-u>Clap personalconf<CR>
`
const CawKeyMap = `
function! InitCaw() abort
	if ! (&l:modifiable && &buftype ==# '')
		silent! nunmap <buffer> gc
		silent! xunmap <buffer> gc
		silent! nunmap <buffer> gcc
		silent! xunmap <buffer> gcc
	else
		nmap <buffer> gc <Plug>(caw:prefix)
		xmap <buffer> gc <Plug>(caw:prefix)
		nmap <buffer> gcc <Plug>(caw:hatpos:toggle)
		xmap <buffer> gcc <Plug>(caw:hatpos:toggle)
	endif
endfunction
autocmd FileType * call InitCaw()
call InitCaw()
`

const CommittiaKeyMap = `
let g:committia_hooks = {}
function! g:committia_hooks.edit_open(info)
  imap <buffer><C-d> <Plug>(committia-scroll-diff-down-half)
  imap <buffer><C-u> <Plug>(committia-scroll-diff-up-half)
  setlocal winminheight=1 winheight=1
  resize 10
  startinsert
endfunction
`

const QuickRunKeyMap = `
nnoremap <silent> <Leader>cr :QuickRun<CR>
`

const VistaKeyMap = `
nnoremap <silent> <Leader>i :<C-u>Vista!!<CR>
`

const EasyMotionKeyMap = `
nmap gsj <Plug>(easymotion-w)
nmap gsk <Plug>(easymotion-b)
nmap gsf <Plug>(easymotion-overwin-f)
`

const MundoKeyMap = `
nnoremap <silent> <Leader>m :MundoToggle<CR>
`

const SandWichKeyMap = `
nmap <silent> sa <Plug>(operator-sandwich-add)
xmap <silent> sa <Plug>(operator-sandwich-add)
omap <silent> sa <Plug>(operator-sandwich-g@)
nmap <silent> sd <Plug>(operator-sandwich-delete)<Plug>(operator-sandwich-release-count)<Plug>(textobj-sandwich-query-a)
xmap <silent> sd <Plug>(operator-sandwich-delete)
nmap <silent> sr <Plug>(operator-sandwich-replace)<Plug>(operator-sandwich-release-count)<Plug>(textobj-sandwich-query-a)  xmap <silent> sr <Plug>(operator-sandwich-replace)
nmap <silent> sdb <Plug>(operator-sandwich-delete)<Plug>(operator-sandwich-release-count)<Plug>(textobj-sandwich-auto-a)
nmap <silent> srb <Plug>(operator-sandwich-replace)<Plug>(operator-sandwich-release-count)<Plug>(textobj-sandwich-auto-a)
omap ib <Plug>(textobj-sandwich-auto-i)
xmap ib <Plug>(textobj-sandwich-auto-i)
omap ab <Plug>(textobj-sandwich-auto-a)
xmap ab <Plug>(textobj-sandwich-auto-a)
omap is <Plug>(textobj-sandwich-query-i)
xmap is <Plug>(textobj-sandwich-query-i)
omap as <Plug>(textobj-sandwich-query-a)
xmap as <Plug>(textobj-sandwich-query-a)
`

const NiceBlockKeyMap = `
silent! xmap I  <Plug>(niceblock-I)
silent! xmap gI <Plug>(niceblock-gI)
silent! xmap A  <Plug>(niceblock-A)
`

const VimExpandRegionKeyMap = `
xmap v <Plug>(expand_region_expand)
xmap V <Plug>(expand_region_shrink)
`
const DsfKeyMap = `
nmap dsf <Plug>DsfDelete
nmap csf <Plug>DsfChange
`

const SplitJoinKeyMap = `
let g:splitjoin_join_mapping = ''
let g:splitjoin_split_mapping = ''
nmap sj :SplitjoinJoin<CR>
nmap sk :SplitjoinSplit<CR>
`

const OperatorReplaceKeyMap = `
xmap p <Plug>(operator-replace)
`

const MultiBlockKeyMap = `
omap <silent> ab <Plug>(textobj-multiblock-a)
omap <silent> ib <Plug>(textobj-multiblock-i)
xmap <silent> ab <Plug>(textobj-multiblock-a)
xmap <silent> ib <Plug>(textobj-multiblock-i)
`

const TextObjFunctionKeyMap = `
omap <silent> af <Plug>(textobj-function-a)
omap <silent> if <Plug>(textobj-function-i)
xmap <silent> af <Plug>(textobj-function-a)
xmap <silent> if <Plug>(textobj-function-i)
`

const VimagitKeyMap = `
nnoremap <silent> <Leader>gg :Magit<CR>
`

const FugiTiveKeyMap = `
nnoremap <silent> <Leader>ga :Git add %:p<CR>
nnoremap <silent> <Leader>gd :Gdiffsplit<CR>
nnoremap <silent> <Leader>gc :Git commit<CR>
nnoremap <silent> <Leader>gb :Git blame<CR>
nnoremap <silent> <Leader>gf :Gfetch<CR>
nnoremap <silent> <Leader>gs :Git<CR>
nnoremap <silent> <Leader>gp :Gpush<CR>
`

const NerdTreeKeyMap = `
nnoremap <silent><leader>e :NERDTreeToggle <CR>
nnoremap <silent><leader>F :NERDTreeFind <CR>
`

const CocExplorerKeyMap = `
nmap <leader>e :CocCommand explorer<CR>
`
