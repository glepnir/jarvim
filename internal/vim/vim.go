// Package vim provides ...
package vim

import (
	"os"

	"github.com/glepnir/jarvis/internal/plugin"
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/render/dein"
)

var (
	ConfPath             = os.ExpandEnv("$HOME/.config/xnvim")
	CachePath            = os.ExpandEnv("$HOME/.cache/nvim/")
	ConfCore             = ConfPath + "/core/"
	ConfAutoload         = ConfPath + "/autoload/"
	ConfModules          = ConfPath + "/modules/"
	ConfPlugin           = ConfPath + "/plugin/"
	PluginManage         string
	Leaderkey            string
	LocalLeaderKey       string
	Colorscheme          []string
	StartScreenPlugin    bool
	StatusLine           bool
	BufferLine           bool
	Explorer             string
	Database             bool
	Fuzzyfind            bool
	EditorConfig         bool
	IndentPlugin         string
	CommentPlugin        bool
	OutLinePlugin        bool
	TagsPlugin           bool
	QuickRun             bool
	DataTypeFile         []string
	EnhancePlugins       []string
	SandwichPlugin       bool
	VersionControlPlugin []string
	UserLanguages        []string

	LeaderKeyMap = map[string]string{
		"Space":        "\\<Space>",
		"Comma(,)":     ",",
		"Semicolon(;)": ";",
	}

	ColorschemeMap = map[string]string{
		"hardcoreplayers/oceanic-material":   "oceanic_materail",
		"drewtempelmeyer/palenight.vim":      "palenight",
		"gruvbox-community/gruvbox":          "gruvbox",
		"ayu-theme/ayu-vim":                  "ayu",
		"NLKNguyen/papercolor-theme":         "PaperColor",
		"lifepillar/vim-gruvbox8":            "gruvbox8",
		"lifepillar/vim-solarized8":          "solarized8",
		"joshdick/onedark.vim":               "onedark",
		"arcticicestudio/nord-vim":           "nord",
		"rakr/vim-one":                       "one",
		"mhartington/oceanic-next":           "OceanicNext",
		"dracula/vim":                        "dracula",
		"chriskempson/base16-vim":            "base16-default-dark",
		"kristijanhusak/vim-hybrid-material": "hybrid_material",
		"nanotech/jellybeans.vim":            "jellybeans",
	}
	VersionControlMap = map[string]string{
		"jreybert/vimagit":     plugin.DeinVimagt,
		"tpope/vim-fugitive":   plugin.DeinFugiTive,
		"lambdalisue/gina.vim": plugin.DeinGina,
	}
	LanguagesPluginMap = map[string]string{
		"C-family":   plugin.DeinCFamily,
		"R":          plugin.DeinR,
		"Javascript": plugin.DeinJavascript,
		"Typescript": plugin.DeinTypescript,
		"Dart":       plugin.DeinDart,
		"React":      plugin.DeinReact,
		"Vue":        plugin.DeinVue,
		"Go":         plugin.DeinGo,
		"Rust":       plugin.DeinRust,
		"Haskell":    plugin.DeinHaskell,
		"Php":        plugin.DeinPhp,
		"Ruby":       plugin.DeinRuby,
		"Scala":      plugin.DeinScala,
		"Shell":      plugin.DeinShell,
		"Lua":        plugin.DeinLua,
		"Python":     plugin.DeinPython,
		"Html":       plugin.DeinHtml,
		"Css":        plugin.DeinCss,
		"Less":       plugin.DeinLess,
		"Sass scss":  plugin.DeinSass,
		"Stylus":     plugin.DeinStylus,
	}
)

func NewDataFileMap(r render.Render) map[string]string {
	_, ok := r.(*dein.Dein)
	if ok {
		return map[string]string{
			"MarkDown":   plugin.DeinMarkDown,
			"Toml":       plugin.DeinToml,
			"Nginx":      plugin.DeinNginx,
			"Json":       plugin.DeinJson,
			"Dockerfile": plugin.DeinDockerFile,
		}
	}
	return map[string]string{
		"MarkDown":   plugin.PlugMarkDown,
		"Toml":       plugin.PlugToml,
		"Nginx":      plugin.PlugNginx,
		"Json":       plugin.PlugJson,
		"Dockerfile": plugin.PlugDockerFile,
	}

}

func NewEnhancePluginMap(r render.Render) map[string]string {
	_, ok := r.(*dein.Dein)
	if ok {
		return map[string]string{
			"accelerated-jk accelerate up-down moving (j and k mapping)": plugin.DeinFastJK,
			"vim-mundo  vim undo tree":                                   plugin.DeinMundo,
			"vim-easymotion fast jump":                                   plugin.DeinEasyMotion,
			"rainbow  rainbow parentheses":                               plugin.DeinRainbow,
			"vim-floterm  vim terminal float":                            plugin.DeinFloaterm,
		}
	}
	return map[string]string{
		"accelerated-jk accelerate up-down moving (j and k mapping)": plugin.PlugFastJK,
		"vim-mundo  vim undo tree":                                   plugin.PlugMundo,
		"vim-easymotion fast jump":                                   plugin.PlugEasyMotion,
		"rainbow  rainbow parentheses":                               plugin.PlugRainbow,
		"vim-floterm  vim terminal float":                            plugin.PlugFloaterm,
	}
}
