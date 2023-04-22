package model

type UserCreateInput struct {
	Nickname   string `v:"required|length:1,20#请输入昵称|昵称长度应当在:min到:max之间"`
	Password   string `v:"required|length:6,20#请输入密码|密码长度应当在:min到:max之间"`
	Email      string `v:"required|email#请输入邮箱|邮箱格式不正确"`
	RePassword string `v:"required|length:6,20|same:password#请再次输入密码|密码长度应当在:min到:max之间|两次输入的密码不一致"`
}

type UserUpdateInput struct {
	Nickname string `v:"required|length:1,20#请输入昵称|昵称长度应当在:min到:max之间"`
}

type UserLoginInput struct {
	Email    string `v:"required|email#请输入邮箱|邮箱格式不正确"`
	Password string `v:"required|length:6,20#请输入密码|密码长度应当在:min到:max之间"`
}
