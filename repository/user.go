package repository

import (
	"errors"
	. "ri/driver"
	. "ri/model"
	"xorm.io/builder"
)

// QueryUser Select user by email and hash password
func QueryUser(email string, hashpassword string) (User, error) {
	user := User{}
	row, err := Engine.Cols("USERUIN", "NAME").Table("USERS").Where("EMAIL = ? AND PASSWORD = ?", email, hashpassword).Get(&user)
	if err != nil {
		return user, err
	} else if !row {
		return user, errors.New("Wrong username or password")
	}
	return user, nil
}

func QueryUserByUin(uin string) (email string, err error) {
	email = ""
	row, err := Engine.Cols("EMAIL").Table("USERS").Where("USERUIN = ?", uin).Get(&email)
	if err != nil {
		return "", err
	} else if !row {
		return "", errors.New("User not exist")
	}
	return email, nil
}

// QueryGroups Select groups info that use in
func QueryGroups(useruin string) ([]Group, error) {
	groups := make([]Group, 0)
	err := Engine.Cols("GROUPUIN", "NAME").Table("GROUPS").In("GROUPUIN", builder.Select("GROUPUIN").From("GROUPMEMBERSHIP").Where(builder.Eq{"useruin": useruin})).Find(&groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// QueryAllGroups Select all system groups
func QueryAllGroups() ([]Group, error) {
	groups := make([]Group, 0)
	err := Engine.Cols("GROUPUIN", "NAME").Table("GROUPS").Find(&groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
