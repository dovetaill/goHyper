package main

import (
	_ "boysave/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"boysave/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
