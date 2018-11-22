package main

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init().RegisterModule(&zapGen{ModuleBase: pgs.ModuleBase{}}).RegisterPostProcessor(pgsgo.GoFmt()).Render()
}
