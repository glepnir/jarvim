// Package main provides ...
package main

import (
	"fmt"

	"github.com/glepnir/jarvis/internal/render"
)

func main() {
	render.GenerateTemplate()
	fmt.Println(render.NeoVimConf)
}
