package main

import (
	_ "boysave/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"boysave/internal/cmd"
	_ "boysave/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
