// Package plugin provides ...
package plugin

const (
	// DeinDein
	DeinDein = `
[[plugins]]
repo = 'Shougo/dein.vim'
`
	//DeinFastJK
	DeinFastJK = `
[[plugins]]
repo = 'rhysd/accelerated-jk'
on_map = {n = '<Plug>'}
hook_add = '''
  nmap <silent>j <Plug>(accelerated_jk_gj)
  nmap <silent>k <Plug>(accelerated_jk_gk)
'''
`
	// DeinMundo
	DeinMundo = `
[[plugins]]
repo = 'simnalamburt/vim-mundo'
on_cmd = 'MundoToggle'
`
	// DeinEasyMotion
	DeinEasyMotion = `
[[plugins]]
repo = 'easymotion/vim-easymotion'
on_map = { n = '<Plug>' }
hook_source = '''
    let g:EasyMotion_do_mapping = 0
    let g:EasyMotion_prompt = 'Jump to → '
    let g:EasyMotion_keys = 'fjdkswbeoavn'
    let g:EasyMotion_smartcase = 1
    let g:EasyMotion_use_smartsign_us = 1
'''
`
	// DeinRainbow
	DeinRainbow = `
[[plugins]]
repo = 'luochen1990/rainbow'
on_ft = [
      'html',
      'css',
      'javascript',
      'javascriptreact',
      'go',
      'python',
      'lua',
      'rust',
      'vim',
      'less',
      'stylus',
      'sass',
      'scss',
      'json',
      'ruby',
      'toml',
    ]
hook_source = '''
	let g:rainbow_active = 1
'''
`
	// DeinFloaterm
	DeinFloaterm = `
repo = 'voldikss/vim-floaterm'
on_cmd = ['FloatermNew', 'FloatermToggle', 'FloatermPrev', 'FloatermNext', 'FloatermSend']
hook_source= '''
    let g:floaterm_position = 'center'
    let g:floaterm_wintype = 'floating'

    " Set floaterm window's background to black
    hi Floaterm guibg=black
    " Set floating window border line color to cyan, and background to orange
    hi FloatermBorder guibg=none guifg=cyan
'''
`
	// PlugFastJK
	PlugFastJK = `
Plug 'rhysd/accelerated-jk'
`
	// PlugMundo
	PlugMundo = `
Plug 'simnalamburt/vim-mundo'
`
	// PlugEasyMotion
	PlugEasyMotion = `
Plug 'easymotion/vim-easymotion'
`
	// PlugRainbow
	PlugRainbow = `
Plug 'luochen1990/rainbow',{'for': [
      'html',
      'css',
      'javascript',
      'javascriptreact',
      'go',
      'python',
      'lua',
      'rust',
      'vim',
      'less',
      'stylus',
      'sass',
      'scss',
      'json',
      'ruby',
      'toml',
    ]}
`

	// PlugFloaterm
	PlugFloaterm = `
Plug 'voldikss/vim-floaterm'
`
	// PlugRainbowSetting
	PlugRainbowSetting = `
" Rainbow
let g:rainbow_active = 1
`
	// PlugFloatermSetting
	PlugFloatermSetting = `
"Floaterm
let g:floaterm_position = 'center'
let g:floaterm_wintype = 'floating'

" Set floaterm window's background to black
hi Floaterm guibg=black
" Set floating window border line color to cyan, and background to orange
hi FloatermBorder guibg=none guifg=cyan
`
	// PlugEasyMotionSetting
	PlugEasyMotionSetting = `
" Easymotion
let g:EasyMotion_do_mapping = 0
let g:EasyMotion_prompt = 'Jump to → '
let g:EasyMotion_keys = 'fjdkswbeoavn'
let g:EasyMotion_smartcase = 1
let g:EasyMotion_use_smartsign_us = 1
`
	// PlugFastJKSetting
	PlugFastJKSetting = `
"accelerated-jk
nmap <silent>j <Plug>(accelerated_jk_gj)
nmap <silent>k <Plug>(accelerated_jk_gk)
`
)
