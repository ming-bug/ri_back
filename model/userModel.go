package model

// User basic user info
type User struct {
	UserUin string `json:"userUin" xorm:"'USERUIN'" example:"1"`
	Name    string `json:"name" xorm:"'NAME'" example:"Shiming"`
	Email   string `json:"email" xorm:"'EMAIL'" example:"shiming.xue@hgc.com.hk"`
}

// LoginInfo login params
type LoginInfo struct {
	Email     string `json:"email" example:"abc@abc.com"`
	Password  string `json:"password" example:"p@ssword1"`
	GroupUin  string `json:"groupUin" example:"2"`
	AutoLogin bool   `json:"autoLogin" example:"true"`
}

// Group user group
type Group struct {
	GroupUin string `json:"groupUin" xorm:"'GROUPUIN'" example:"1024"`
	Name     string `json:"name" xorm:"'NAME'" example:"DEFAULT GROUP"`
}

// UserInfo user info
type UserInfo struct {
	User
	CurrentGroup Group   `json:"currentGroup"`
	Groups       []Group `json:"groups"`
}

func (ui *UserInfo) UserInfo() *UserInfo {
	return ui
}

func (ui *UserInfo) SetUserInfo(user User, currentGroup Group, groups []Group) {
	ui.User = user
	ui.CurrentGroup = currentGroup
	ui.Groups = groups
}
