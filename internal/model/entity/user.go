package entity

type User struct {
	Id            uint   `json:"id" description:"User ID"`
	Password      string `json:"password" description:"User password"`
	Nickname      string `json:"nickname" description:"User nickname"`
	Email         string `json:"email" description:"User email"`
	Avatar        string `json:"avatar" description:"User avatar"`
	Role          string `json:"role" description:"User role"`
	LastLoginIp   string `json:"last_login_ip" description:"User last login ip"`
	LastLoginTime string `json:"last_login_time" description:"User last login time"`
	CreateTime    string `json:"create_time" description:"User create time"`
	UpdateTime    string `json:"update_time" description:"User update time"`
}
