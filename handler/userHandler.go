package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"ri/enum"
	"ri/model"
	. "ri/repository"
	. "ri/utils"
)

var userInfo model.UserInfo

// AllGroups System groups
// @Summary System groups
// @Description Get all system groups
// @Tags user
// @Produce application/json
// @Success 200 {object} model.ResultCont
// @Failure 404 {object} model.ResultCont
// @Router /api/user/groups [get]
func AllGroups(c *gin.Context) {
	result := model.NewResult(c)
	groups, err := QueryAllGroups()
	if err != nil {
		result.Faild(http.StatusNotFound, err.Error(), enum.Error)
	}
	result.Success(groups)
}

// Login login interface
// @Summary login with login info
// @Description verify login account by email, password and group
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param user body model.LoginInfo true "user login info"
// @Success 200 {object} model.ResultCont
// @Failure 500 {object} model.ResultCont
// @Failure 404 {object} model.ResultCont
// @Failure 403 {object} model.ResultCont
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	loginInfo := model.LoginInfo{}
	err := c.ShouldBindBodyWith(&loginInfo, binding.JSON)
	result := model.NewResult(c)
	if err != nil {
		result.Faild(http.StatusInternalServerError, err.Error(), enum.Page)
		return
	}
	hashPassword := HashPassword(loginInfo.Password)
	user, err := QueryUser(loginInfo.Email, hashPassword)
	if err != nil {
		result.Faild(http.StatusNotFound, err.Error(), enum.Error)
		return
	}
	user.Email = loginInfo.Email
	groups, err := QueryGroups(user.UserUin)
	if err != nil {
		result.Faild(http.StatusInternalServerError, err.Error(), enum.Error)
	}
	inGroup := InGroup(loginInfo.GroupUin, groups)
	if !inGroup {
		result.Faild(http.StatusForbidden, "Permission denied, you are not allown with this group.", 1)
		return
	}
	token, err := GenerateToken(model.HmacUser{UserUin: user.UserUin, Username: user.Name}, loginInfo.AutoLogin)
	if err != nil {
		result.Faild(http.StatusInternalServerError, err.Error(), 9)
		return
	}
	loginResponse := model.LoginResponse{
		Token: token,
	}

	group := model.Group{}
	for i := range groups {
		if groups[i].GroupUin == loginInfo.GroupUin {
			group.GroupUin = groups[i].GroupUin
			group.Name = groups[i].Name
		}
	}
	userInfo.SetUserInfo(user, group, groups)

	result.Success(loginResponse)
}

// InGroup If selected group in the visited group list
func InGroup(groupUin string, groups []model.Group) bool {
	for i := range groups {
		if groups[i].GroupUin == groupUin {
			return true
		}
	}
	return false
}

// CurrentUser get current user
// @Summary get current user
// @Description get current user info, include name & groups
// @Tags user
// @Produce application/json
// @Success 200 {object} model.ResultCont
// @Failure 401 {object} model.ResultCont
// @Router /api/user/current [get]
func CurrentUser(c *gin.Context) {
	result := model.NewResult(c)
	info := userInfo.UserInfo()
	if info.UserUin != "" {
		result.Success(info)
		return
	}
	result.Faild(http.StatusUnauthorized, "please login first", enum.Error)
}

// OutLogin logout
// @Summary logout current account
// @Description logout and return nil
// @Tags user
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.ResultCont
// @Router /api/user/outLogin [post]
func OutLogin(c *gin.Context) {
	result := model.NewResult(c)
	user := model.User{}
	groups := []model.Group{}
	group := model.Group{}
	userInfo.SetUserInfo(user, group, groups)
	result.Success(userInfo)
}
