package v1

import (
	"boysave/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}

type ProfileRes struct {
	*entity.User
}

type SignUpReq struct {
	g.Meta     `path:"/user/signUp" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Password   string `v:"required|length:6,16"`
	RePassword string `v:"required|length:6,16|same:Password"`
	Nickname   string
	Email      string `v:"required|email"`
}

type SignUpRes struct {
}

type SignInReq struct {
	g.Meta   `path:"/user/signIn" method:"post" tags:"UserService" summary:"Sign in with user account"`
	Password string `v:"required|length:6,16"`
	Email    string `v:"required|email"`
}

type SignInRes struct {
}

type CheckNickNameReq struct {
	g.Meta   `path:"/user/checkNickName" method:"post" tags:"UserService" summary:"Check if the nickname is available"`
	Nickname string `v:"required|length:1,16"`
}

type CheckNickNameRes struct {
}

type IsSignInReq struct {
	g.Meta `path:"/user/isSignIn" method:"post" tags:"UserService" summary:"Check if the user is signed in"`
}

type IsSignInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type SignOutReq struct {
	g.Meta `path:"/user/signOut" method:"post" tags:"UserService" summary:"Sign out current user"`
}

type SignOutRes struct {
}
