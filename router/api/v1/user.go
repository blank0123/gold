package v1

import (
	"fmt"
	"github.com/kainhuck/gold/models"
	"github.com/kainhuck/gold/pkg/app"
	"github.com/kainhuck/gold/pkg/config"
	"github.com/kainhuck/gold/pkg/e"
	"github.com/kainhuck/gold/pkg/log"
	"github.com/kainhuck/gold/pkg/util"
	"github.com/kainhuck/gold/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 1. 用户注册
func Register(ctx *gin.Context) {
	var user user_service.UserRegister
	appG := app.Gin{C: ctx}
	// 1. 用户输入用户名，密码，确认密码
	httpCode, eCode := app.BindAndValid(ctx, &user)
	if eCode != e.SUCCESS{
		appG.Response(httpCode, eCode, nil)
		return
	}

	// 2. 判断用户名是否已经被注册
	if user.CheckUsername() {
		appG.Response(http.StatusOK, e.ERROR_EXIST_USERNAME, nil)
		return
	}

	// 3. 密码和确认密码是否一样
	if !user.CheckSamePassword() {
		appG.Response(http.StatusOK, e.ERROR_DIFF_PASSWORD, nil)
		return
	}

	// 4. OK -> database
	ok := user.Save()
	if !ok {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 2. 用户登录
func Login(ctx *gin.Context) {
	var user user_service.UserLogin
	appG := app.Gin{ctx}

	// 1. 获取用户名密码
	httpCode, eCode := app.BindAndValid(ctx, &user)
	if eCode != e.SUCCESS{
		appG.Response(httpCode, eCode, nil)
		return
	}

	// 2. 查找用户名是否存在
	if !user.IsExistUsername() {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USERNAME, nil)
		return
	}

	// 3. 判断密码是否正确
	if !user.CorrectPassword() {
		appG.Response(http.StatusOK, e.ERROR_INCORRECT_PASSWORD, nil)
		return
	}

	// 4. 生成token，返回
	token, err := util.GenerateToken(user.Username, user.Password)
	if err != nil {
		log.SugarLogger.Errorf("生成token失败, err: %v", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// 3. 用户修改密码
func EditPassword(ctx *gin.Context) {
	var user user_service.UserEditPassword
	appG := app.Gin{ctx}

	// 1. 输入用户名，密码，新密码
	httpCode, eCode := app.BindAndValid(ctx, &user)
	fmt.Println(httpCode, eCode, user)
	if eCode != e.SUCCESS{
		appG.Response(httpCode, eCode, nil)
		return
	}

	// 2. 验证用户名是否存在
	if !user.IsExistUsername() {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USERNAME, nil)
		return
	}

	// 3. 验证密码是否正确
	if !user.CorrectPassword() {
		appG.Response(http.StatusOK, e.ERROR_INCORRECT_PASSWORD, nil)
		return
	}

	// 4. 修改密码
	if !user.EditPassword() {
		appG.Response(http.StatusOK, e.ERROR_EDIT_PASSWORD, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 4. 管理员获取所有用户
func GetUsers(ctx *gin.Context) {
	appG := app.Gin{ctx}
	// 1. 获取传来的page
	num := util.GetPage(ctx)
	user := &user_service.UserGets{
		PageNum:  num,
		PageSize: config.Collection.App.PageSize,
	}

	users, err := user.Get()
	if err != nil {
		log.SugarLogger.Errorf("获取用户失败, err: %v", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, users)

}

// 5. 管理员获取指定用户
func GetUser(ctx *gin.Context) {
	appG := app.Gin{ctx}
	// 1. 获取用户ID
	id := ctx.Param("id")

	// 2. 判断用户id是否存在
	exist, err := models.IsExistID(id)
	if err != nil {
		log.SugarLogger.Errorf("通过id判断用户是否存在失败, err: %v", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USERID, nil)
		return
	}

	// 3. 获取用户信息
	user, err := models.GetUserByID(id)
	if err != nil {
		log.SugarLogger.Errorf("通过id获取用户信息失败, err: %v", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}

// 6. 管理员删除指定用户
func DeleteUser(ctx *gin.Context) {
	appG := app.Gin{ctx}
	// 1. 获取用户的id
	id := ctx.Param("id")

	// 2. 判断是否存在该用户
	exist, err := models.IsExistID(id)
	if err != nil {
		log.SugarLogger.Errorf("通过id判断用户是否存在失败, err: %v", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USERID, nil)
		return
	}

	// 3. 数据库中删除该用户
	err = models.DeleteUserByID(id)

	if err != nil {
		log.SugarLogger.Errorf("通过id删除用户失败, err: %v", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
