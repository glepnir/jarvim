package plugin

const (
	DeinCoC = `
[[plugins]]
repo = 'neoclide/coc.nvim'
merged = 0
rev = 'release'
hook_add = '''
    let g:coc_snippet_next = '<TAB>'
    let g:coc_snippet_prev = '<S-TAB>'
    let g:coc_status_error_sign = '•'
    let g:coc_status_warning_sign = '•'
	{{ if index . 1}}
    let g:coc_global_extensions =[
		\ 'coc-snippets',
		\ 'coc-pairs',
		\ 'coc-json',
		\ 'coc-highlight',
		\ 'coc-git',
		\ 'coc-emoji',
		\ 'coc-lists',
		\ 'coc-stylelint',
		\ 'coc-yaml',
		\ 'coc-gitignore',
		\ 'coc-yank',
		\ 'coc-actions',
		\ 'coc-db',
		\ 'coc-spell-checker',
		\ 'coc-vimlsp',
		\ 'coc-explorer',
        \]
	{{else}}
    let g:coc_global_extensions =[
		\ 'coc-snippets',
		\ 'coc-pairs',
		\ 'coc-json',
		\ 'coc-highlight',
		\ 'coc-git',
		\ 'coc-emoji',
		\ 'coc-lists',
		\ 'coc-stylelint',
		\ 'coc-yaml',
		\ 'coc-gitignore',
		\ 'coc-yank',
		\ 'coc-actions',
		\ 'coc-db',
		\ 'coc-spell-checker',
		\ 'coc-vimlsp',
        \]
	{{end}}

    augroup coc_event
      autocmd!
      " Setup formatexpr specified filetype(s).
      autocmd FileType typescript,json setl formatexpr=CocAction('formatSelected')
      " Update signature help on jump placeholder
      autocmd User CocJumpPlaceholder call CocActionAsync('showSignatureHelp')
    augroup end

    " Highlight symbol under cursor on CursorHold
    autocmd CursorHold * silent call CocActionAsync('highlight')

    "Use tab for trigger completion with characters ahead and navigate
    inoremap <silent><expr> <TAB>
          \ pumvisible() ? "\<C-n>" :
          \ <SID>check_back_space() ? "\<TAB>" :
          \ coc#refresh()

    inoremap <expr><S-TAB> pumvisible() ? "\<C-p>" : "\<C-h>"
    inoremap <silent><expr> <cr> pumvisible() ? coc#_select_confirm() : "\<C-g>u\<CR>\<c-r>=coc#on_enter()\<CR>"

    function! s:check_back_space() abort
      let col = col('.') - 1
      return !col || getline('.')[col - 1]  =~# '\s'
    endfunction
'''

[[plugins]]
repo = 'honza/vim-snippets'
depends = 'coc.nvim'
if = 'has("python3")'
merged =  0
`
	PlugCoc = `
Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'honza/vim-snippets'
`

	PlugCocSetting = `
let g:coc_snippet_next = '<TAB>'
let g:coc_snippet_prev = '<S-TAB>'
let g:coc_status_error_sign = '•'
let g:coc_status_warning_sign = '•'

{{ if index . 1}}
let g:coc_global_extensions =[
	\ 'coc-snippets',
	\ 'coc-pairs',
	\ 'coc-json',
	\ 'coc-highlight',
	\ 'coc-git',
	\ 'coc-emoji',
	\ 'coc-lists',
	\ 'coc-stylelint',
	\ 'coc-yaml',
	\ 'coc-gitignore',
	\ 'coc-yank',
	\ 'coc-actions',
	\ 'coc-db',
	\ 'coc-spell-checker',
	\ 'coc-vimlsp',
	\ 'coc-explorer',
    \]
{{else}}
let g:coc_global_extensions =[
	\ 'coc-snippets',
	\ 'coc-pairs',
	\ 'coc-json',
	\ 'coc-highlight',
	\ 'coc-git',
	\ 'coc-emoji',
	\ 'coc-lists',
	\ 'coc-stylelint',
	\ 'coc-yaml',
	\ 'coc-gitignore',
	\ 'coc-yank',
	\ 'coc-actions',
	\ 'coc-db',
	\ 'coc-spell-checker',
	\ 'coc-vimlsp',
    \]
{{end}}


augroup coc_event
  autocmd!
  " Setup formatexpr specified filetype(s).
  autocmd FileType typescript,json setl formatexpr=CocAction('formatSelected')
  " Update signature help on jump placeholder
  autocmd User CocJumpPlaceholder call CocActionAsync('showSignatureHelp')
augroup end

" Highlight symbol under cursor on CursorHold
autocmd CursorHold * silent call CocActionAsync('highlight')

"Use tab for trigger completion with characters ahead and navigate
inoremap <silent><expr> <TAB>
      \ pumvisible() ? "\<C-n>" :
      \ <SID>check_back_space() ? "\<TAB>" :
      \ coc#refresh()

inoremap <expr><S-TAB> pumvisible() ? "\<C-p>" : "\<C-h>"
inoremap <silent><expr> <cr> pumvisible() ? coc#_select_confirm() : "\<C-g>u\<CR>\<c-r>=coc#on_enter()\<CR>"

function! s:check_back_space() abort
  let col = col('.') - 1
  return !col || getline('.')[col - 1]  =~# '\s'
endfunction
`
)

const CocJson = `
{
  "suggest.triggerAfterInsertEnter": false,
  "suggest.timeout": 500,
  "suggest.noselect": false,
  "suggest.detailField": "abbr",
  "suggest.snippetIndicator": "🌟",
  "suggest.triggerCompletionWait": 100,
  "suggest.echodocSupport": true,
  "suggest.completionItemKindLabels": {
    "keyword": "\uf1de",
    "variable": "\ue79b",
    "value": "\uf89f",
    "operator": "\u03a8",
    "function": "\u0192",
    "reference": "\ufa46",
    "constant": "\uf8fe",
    "method": "\uf09a",
    "struct": "\ufb44",
    "class": "\uf0e8",
    "interface": "\uf417",
    "text": "\ue612",
    "enum": "\uf435",
    "enumMember": "\uf02b",
    "module": "\uf40d",
    "color": "\ue22b",
    "property": "\ue624",
    "field": "\uf9be",
    "unit": "\uf475",
    "event": "\ufacd",
    "file": "\uf15c",
    "folder": "\uf114",
    "snippet": "\ue60b",
    "typeParameter": "\uf728",
    "default": "\uf29c"
  },
  //diagnostic
  "diagnostic.signOffset": 9999999,
  "diagnostic.errorSign": "●",
  "diagnostic.warningSign": "●",
  "diagnostic.infoSign": "●",
  "diagnostic.displayByAle": false,
  "diagnostic.refreshOnInsertMode": true,
  //git
  "git.enableGutters": true,
  "git.branchCharacter": "\uf418",
  "git.addGBlameToBufferVar": true,
  "git.addGBlameToVirtualText": true,
  "git.virtualTextPrefix": " ❯❯❯ ",
  "git.addedSign.hlGroup": "GitGutterAdd",
  "git.changedSign.hlGroup": "GitGutterChange",
  "git.removedSign.hlGroup": "GitGutterDelete",
  "git.topRemovedSign.hlGroup": "GitGutterDelete",
  "git.changeRemovedSign.hlGroup": "GitGutterChangeDelete",
  "git.addedSign.text": "▋",
  "git.changedSign.text": "▋",
  "git.removedSign.text": "▋",
  "git.topRemovedSign.text": "▔",
  "git.changeRemovedSign.text": "▋",
  //Snippet
  "coc.preferences.snippetStatusText": "Ⓢ ",
  //CocList
  "list.source.files.defaultOptions": ["--auto-preview"],
  "list.source.outline.defaultOptions": ["--auto-preview"],
  //prettier
  "coc.preferences.formatOnSaveFiletypes": [
    "html",
    "css",
    "markdown",
    "javascript",
    "javascriptreact",
    "typescript",
    "typescriptreact",
    "go",
    "json"
  ],
  "prettier.statusItemText": "ⓟ ",
  "prettier.eslintIntegration": true,
  "prettier.tslintIntegration": false,
  "prettier.stylelintIntegration": true,
  "prettier.disableSuccessMessage": true,
  //emmet
  "emmet.includeLanguages": {
    "vue-html": "html",
    "javascript": "javascriptreact",
    "typescriptreact": "typescriptreact"
  },
  //imselect
  "imselect.enableStatusItem": false,
  //multiple curssor
  "cursors.nextKey": "<C-n>",
  "cursors.previousKey": "<C-p>",
  "cursors.cancelKey": "<esc>",
  //sessions
  "session.directory": "~/.cache/vim/session",
  // coc-explorer
  "explorer.icon.expanded": "▾",
  "explorer.icon.collapsed": "▸",
  "explorer.width": 30,
  "explorer.icon.enableNerdfont": true,
  "explorer.quitOnOpen": true,
  "explorer.floating.width": -20,
  "explorer.floating.height": -10,
  "explorer.openAction.strategy": "sourceWindow",
  // eslint
  "eslint.filetypes": [
    "javascript",
    "javascriptreact",
    "typescript",
    "typescriptreact"
  ],
  "eslint.autoFix": true,
  "eslint.autoFixOnSave": true,
  // coc-actions
  "coc-actions.hideCursor": false,
  // Cspell
  "cSpell.fixSpellingWithRenameProvider": true,
  "cSpell.showStatus": false,
  "cSpell.enabledLanguageIds": [
    "asciidoc",
    "c",
    "cpp",
    "csharp",
    "css",
    "git-commit",
    "gitcommit",
    "go",
    "haskell",
    "html",
    "java",
    "javascript",
    "javascriptreact",
    "json",
    "jsonc",
    "latex",
    "less",
    "markdown",
    "php",
    "plaintext",
    "python",
    "pug",
    "restructuredtext",
    "rust",
    "scala",
    "scss",
    "text",
    "typescript",
    "typescriptreact",
    "yaml",
    "yml",
    "vim"
  ]
}
`
