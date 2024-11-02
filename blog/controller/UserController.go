package controller

import (
	"blog/common"
	"blog/dto"
	"blog/model"
	"blog/response"
	"blog/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(cxt *gin.Context) {
	DB := common.GetDB()
	var resquestUser = model.User{}
	cxt.Bind(&resquestUser)

	//获取参数
	name := resquestUser.Name
	telephone := resquestUser.Telephone
	password := resquestUser.Password

	//数据验证
	if len(telephone) != 11 {
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位")
		return
	}
	if len(password) < 8 {
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "密码不能少于8位")
	}
	if len(name) == 0 {
		name = util.RandString(10)
	}

	if IsTelephoneExists(DB, telephone) {
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(cxt, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	if err := DB.Create(&newUser).Error; err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{
			"msg": "用户注册失败",
		})
		return
	}
	//返回结果
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(cxt, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Print("token generate error:%v", err.Error())
		return
	}
	response.Success(cxt, gin.H{"token": token}, "注册成功")
}

func LogIn(cxt *gin.Context) {
	DB := common.GetDB()
	var resquestUser = model.User{}
	cxt.Bind(&resquestUser)

	//获取参数
	telephone := resquestUser.Telephone
	password := resquestUser.Password

	if len(telephone) != 11 {
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位")
		return
	}
	if len(password) < 8 {
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "密码不能少于8位")
		return
	}

	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(cxt, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(cxt, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(resquestUser)
	if err != nil {
		response.Response(cxt, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		return
	}
	response.Success(cxt, gin.H{"token": token}, "登陆成功")
}

func Info(cxt *gin.Context) {
	user, _ := cxt.Get("user")
	cxt.JSON(http.StatusOK, gin.H{
		"code": 200,
		//TODO(下面转dto的方式没有看懂，需要后面再想想看)
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}

func IsTelephoneExists(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
