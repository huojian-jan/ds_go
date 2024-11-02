package controller

import (
	"blog/common"
	"blog/model"
	"blog/response"
	"github.com/gin-gonic/gin"
)

func CreatePost(cxt *gin.Context) {
	DB := common.GetDB()
	var requestPost model.Post
	if err := cxt.Bind(&requestPost); err != nil {
		response.Fail(cxt, gin.H{"msg": "绑定失败"}, "")
		return
	}

	//获取用户信息
	user, _ := cxt.Get("user")
	var category model.Category
	DB.Where("name = ?", requestPost.CategoryName).First(&category)
	if category.ID == 0 {
		response.Fail(cxt, gin.H{"msg": "分类不存在"}, "")
		return
	}

	post := model.Post{
		UserId:       user.(model.User).ID,
		CategoryId:   category.ID,
		Content:      requestPost.Content,
		CategoryName: category.Name,
		HeadImg:      requestPost.HeadImg,
		Title:        requestPost.Title,
	}
	DB.AutoMigrate(&model.Post{})
	if err := DB.Create(&post).Error; err != nil {
		response.Fail(cxt, gin.H{"msg": err.Error()}, "")
		return
	}
	response.Success(cxt, gin.H{"msg": "发布成功"}, "")
}

func DeletePost(cxt *gin.Context) {

}

func UpdatePost(cxt *gin.Context) {

}

func GetPost(cxt *gin.Context) {

}

func GetPosts(cxt *gin.Context) {

}
