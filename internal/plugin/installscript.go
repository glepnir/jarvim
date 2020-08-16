// Package plugin provides ...
package plugin

const (
	DeinMakeFile = `
SHELL = /bin/bash
vim := $(if $(shell which nvim),nvim,$(shell which vim))
vim_version := '${shell $(vim) --version}'
XDG_CACHE_HOME ?= $(HOME)/.cache

default: install

install:
	@mkdir -vp "$(XDG_CACHE_HOME)/vim/"{backup,session,swap,tags,undo}; \
	$(vim)  -V1 -es -i NONE -N -u core/dein.vim -c "try | call dein#update() | call dein#recache_runtimepath() | finally | echomsg '' | qall! | endtry"
	@rm -rf "$(XDG_CACHE_HOME)/vim/dein/cache_vim";\
	@rm -rf "$(XDG_CACHE_HOME)/vim/dein/state_nvim.vim";\
	@rm -rf "$(XDG_CACHE_HOME)/vim/dein/.cache/";\
	$(vim) -u init.vim -c 'call dein#recache_runtimepath()|q'

upgrade:
	$(vim) -V1 -es -i NONE -N -u config/init.vim -c "try | call dein#clear_state() | call dein#update() | finally | qall! | endtry"
`
	DeinInstallShell = `
#!/usr/bin/env bash

# Colors
ESC_SEQ="\x1b["
COL_RESET=$ESC_SEQ"39;49;00m"
COL_RED=$ESC_SEQ"31;01m"
COL_GREEN=$ESC_SEQ"32;01m"
COL_YELLOW=$ESC_SEQ"33;01m"
COL_BLUE=$ESC_SEQ"34;01m"
COL_MAGENTA=$ESC_SEQ"35;01m"
COL_CYAN=$ESC_SEQ"36;01m"

function ok() {
    echo -e "$COL_GREEN[ok]$COL_CYAN "$1
}

function running() {
    echo -e "$COL_YELLOW ⇒ $COL_RESET "$1": "
}

function action() {
    echo -e "\n$COL_YELLOW[action]:$COL_MAGENTA\n ⇒ $1..."
}

function warn() {
    echo -e "$COL_YELLOW[warning]$COL_RESET "$1
}

function error() {
    echo -e "$COL_RED[error]$COL_RESET "$1
}

_try_pyenv() {
	local name='' src=''
	if hash pyenv 2>/dev/null; then
		echo '===> pyenv found, searching virtualenvs…'
		for name in 'neovim' 'neovim3' 'nvim'; do
			src="$(pyenv prefix "${name}" 2>/dev/null)"
			if [ -d "${src}" ]; then
				error "===> pyenv virtualenv found '${name}'"
				# Symlink virtualenv for easy access
				ln -fhs "${src}" "${__venv}"
				return 0
			fi
		done
		warn "===> skipping pyenv. manual virtualenv isn't found"
		warn
		warn "Press Ctrl+C and use pyenv to create one yourself (name it 'neovim')"
		warn "and run ${0} again. Or press Enter to continue and try 'python3'."
		read -r
	else
		warn '===> pyenv not found, skipping'
	fi
	return 1
}

_try_python() {
	if ! hash python3 2>/dev/null; then
		warn '===> python3 not found, skipping'
		return 1
	fi
	ok "===> python3 found"
	[ -d "${__venv}" ] || python3 -m venv "${__venv}"
}

Install_Pynvim() {
	# Concat a base path for vim cache and virtual environment
	local __cache="${XDG_CACHE_HOME:-$HOME/.cache}/vim"
	local __venv="${__cache}/venv"
	mkdir -p "${__cache}"

	if [ -d "${__venv}/neovim3" ]; then
		error -n '===> ERROR: Python 2 has ended its life, only one virtualenv is '
		warn 'created now.'
		warn "Delete '${__venv}' (or backup!) first, and then run ${0} again."
	elif _try_pyenv || _try_python; then
		# Install Python 3 requirements
		"${__venv}/bin/pip" install -U pynvim PyYAML Send2Trash
		ok '===> success'
	else
		error '===> ERROR: unable to setup python3 virtualenv.'
		warn -e '\nConsider using pyenv with its virtualenv plugin:'
		warn '- https://github.com/pyenv/pyenv'
		warn '- https://github.com/pyenv/pyenv-virtualenv'
	fi
}

action "Checking node and yarn..."

node --version | grep "v" &> /dev/null
if [ $? != 0 ]; then
  error "Node not installed"
  warn "Please install node use this script 'curl -sL install-node.now.sh/lts | bash' "
  exit 1;
fi

yarn --version | grep "v" &> /dev/null
if [ $? == 0 ]; then
  error "yarn not installed"
  warn "Please install yarn use this script 'curl --compressed -o- -L https://yarnpkg.com/install.sh | bash' "
  exit 1;
fi

ok "===> check pass"

action "Install tools"

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac


if [ "$(uname)" == "Darwin" ]; then
  running "Found you use mac"
  brew install bat
  brew install ripgrep
  brew install --HEAD universal-ctags/universal-ctags/universal-ctags
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
  running "Found you use Linux"
  if [ -x "$(command -v apk)" ];then sudo apk add bat ripgrep ; fi
  if [ -x "$(command -v pkg)" ];then sudo pkg install bat ripgrep ; fi
  if [ -x "$(command -v pacman)" ];then sudo pacman -S bat ripgrep ; fi
  if [ -x "$(command -v apt-get)" ]; then sudo apt-get install bat ripgrep ; fi
  if [ -x "$(command -v dnf)" ]; then sudo dnf install bat ripgrep ; fi
  if [ -x "$(command -v nix-env)" ]; then sudo nix-env -i bat ripgrep ; fi
  if [ -x "$(command -v zypper)" ]; then sudo zypper install bat ripgrep ; fi
  if [ -x "$(command -v yum)" ]; then sudo yum install bat ripgrep ; fi
fi

action "Install pynvim..."
Install_Pynvim

action "Install plugins"
cd $HOME/.config/nvim/
make
running "Clean up..."
nvim -u init.vim -c 'call dein#recache_runtimepath()|q'

ok "\n
Congratulations thinkvim install success!!!\n
Install your favorite font on here https://www.nerdfonts.com/font-downloads\n
If you use linux,you need install ctags with janson support.\n
Install the Lsp for your languages.\n
Thanks for you love this neovim config."
`

	PlugMakefile = `
SHELL = /bin/bash
vim := $(if $(shell which nvim),nvim,$(shell which vim))
vim_version := '${shell $(vim) --version}'
XDG_CACHE_HOME ?= $(HOME)/.cache

default: install

install:
	@mkdir -vp "$(XDG_CACHE_HOME)/vim/"{backup,session,swap,tags,undo}; \
	$(vim) +'PlugInstall --sync' +qa

upgrade:
	$(vim) +'PlugUpdate --sync' +qa
`

	PlugInstallShell = `
#!/usr/bin/env bash

# Colors
ESC_SEQ="\x1b["
COL_RESET=$ESC_SEQ"39;49;00m"
COL_RED=$ESC_SEQ"31;01m"
COL_GREEN=$ESC_SEQ"32;01m"
COL_YELLOW=$ESC_SEQ"33;01m"
COL_BLUE=$ESC_SEQ"34;01m"
COL_MAGENTA=$ESC_SEQ"35;01m"
COL_CYAN=$ESC_SEQ"36;01m"

function ok() {
    echo -e "$COL_GREEN[ok]$COL_CYAN "$1
}

function running() {
    echo -e "$COL_YELLOW ⇒ $COL_RESET "$1": "
}

function action() {
    echo -e "\n$COL_YELLOW[action]:$COL_MAGENTA\n ⇒ $1..."
}

function warn() {
    echo -e "$COL_YELLOW[warning]$COL_RESET "$1
}

function error() {
    echo -e "$COL_RED[error]$COL_RESET "$1
}

_try_pyenv() {
	local name='' src=''
	if hash pyenv 2>/dev/null; then
		echo '===> pyenv found, searching virtualenvs…'
		for name in 'neovim' 'neovim3' 'nvim'; do
			src="$(pyenv prefix "${name}" 2>/dev/null)"
			if [ -d "${src}" ]; then
				error "===> pyenv virtualenv found '${name}'"
				# Symlink virtualenv for easy access
				ln -fhs "${src}" "${__venv}"
				return 0
			fi
		done
		warn "===> skipping pyenv. manual virtualenv isn't found"
		warn
		warn "Press Ctrl+C and use pyenv to create one yourself (name it 'neovim')"
		warn "and run ${0} again. Or press Enter to continue and try 'python3'."
		read -r
	else
		warn '===> pyenv not found, skipping'
	fi
	return 1
}

_try_python() {
	if ! hash python3 2>/dev/null; then
		warn '===> python3 not found, skipping'
		return 1
	fi
	ok "===> python3 found"
	[ -d "${__venv}" ] || python3 -m venv "${__venv}"
}

Install_Pynvim() {
	# Concat a base path for vim cache and virtual environment
	local __cache="${XDG_CACHE_HOME:-$HOME/.cache}/vim"
	local __venv="${__cache}/venv"
	mkdir -p "${__cache}"

	if [ -d "${__venv}/neovim3" ]; then
		error -n '===> ERROR: Python 2 has ended its life, only one virtualenv is '
		warn 'created now.'
		warn "Delete '${__venv}' (or backup!) first, and then run ${0} again."
	elif _try_pyenv || _try_python; then
		# Install Python 3 requirements
		"${__venv}/bin/pip" install -U pynvim PyYAML Send2Trash
		ok '===> success'
	else
		error '===> ERROR: unable to setup python3 virtualenv.'
		warn -e '\nConsider using pyenv with its virtualenv plugin:'
		warn '- https://github.com/pyenv/pyenv'
		warn '- https://github.com/pyenv/pyenv-virtualenv'
	fi
}

action "Checking node and yarn..."

node --version | grep "v" &> /dev/null
if [ $? != 0 ]; then
  error "Node not installed"
  warn "Please install node use this script 'curl -sL install-node.now.sh/lts | bash' "
  exit 1;
fi

yarn --version | grep "v" &> /dev/null
if [ $? == 0 ]; then
  error "yarn not installed"
  warn "Please install yarn use this script 'curl --compressed -o- -L https://yarnpkg.com/install.sh | bash' "
  exit 1;
fi

ok "===> check pass"

action "Install tools"

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac


if [ "$(uname)" == "Darwin" ]; then
  running "Found you use mac"
  brew install bat
  brew install ripgrep
  brew install --HEAD universal-ctags/universal-ctags/universal-ctags
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
  running "Found you use Linux"
  if [ -x "$(command -v apk)" ];then sudo apk add bat ripgrep ; fi
  if [ -x "$(command -v pkg)" ];then sudo pkg install bat ripgrep ; fi
  if [ -x "$(command -v pacman)" ];then sudo pacman -S bat ripgrep ; fi
  if [ -x "$(command -v apt-get)" ]; then sudo apt-get install bat ripgrep ; fi
  if [ -x "$(command -v dnf)" ]; then sudo dnf install bat ripgrep ; fi
  if [ -x "$(command -v nix-env)" ]; then sudo nix-env -i bat ripgrep ; fi
  if [ -x "$(command -v zypper)" ]; then sudo zypper install bat ripgrep ; fi
  if [ -x "$(command -v yum)" ]; then sudo yum install bat ripgrep ; fi
fi

action "Install pynvim..."
Install_Pynvim

action "Install plugins"
cd $HOME/.config/nvim/
make
running "Clean up..."

ok "\n
Congratulations thinkvim install success!!!\n
Install your favorite font on here https://www.nerdfonts.com/font-downloads\n
If you use linux,you need install ctags with janson support.\n
Install the Lsp for your languages.\n
Thanks for you love this neovim config."
`
)
