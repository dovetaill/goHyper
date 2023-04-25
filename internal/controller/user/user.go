package user

import (
	v1 "boysave/api/hello/v1"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignUpRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello SignIn!")
	return
}
