package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type User struct {
	g.Meta        `orm:"table:users, do:true"`
	Id            interface{}
	Password      interface{}
	Nickname      interface{}
	Email         interface{}
	Avatar        interface{}
	Role          interface{}
	LastLoginIp   interface{}
	LastLoginTime interface{}
	CreateTime    *gtime.Time
	UpdateTime    *gtime.Time
}
