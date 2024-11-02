package controller

import (
	"blog/common"
	"blog/model"
	"blog/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(cxt *gin.Context) {
	DB := common.GetDB()
	var requestCategory model.Category
	cxt.Bind(&requestCategory)
	name := requestCategory.Name

	//验证数据合法性
	//TODO(研究一下，是不是有现成的库，可以做参数验证)
	if len(name) == 0 {
		response.Response(cxt, http.StatusUnprocessableEntity, 400, nil, "分类名称不能为空")
		return
	}
	//检查是否存在
	var existedCategory model.Category
	DB.Where("name = ?", name).First(&existedCategory)
	if existedCategory.ID != 0 {
		response.Response(cxt, http.StatusUnprocessableEntity, 400, nil, fmt.Sprintf("分类%s已存在", name))
		return
	}

	newCategory := model.Category{
		Name:        name,
		Instruction: requestCategory.Instruction,
	}
	DB.AutoMigrate(&model.Category{})
	if err := DB.Create(&newCategory).Error; err != nil {
		response.Response(cxt, http.StatusInternalServerError, 500, gin.H{"msg": err.Error()}, "数据入库失败")
		return
	}
	response.Response(cxt, http.StatusOK, 200, nil, fmt.Sprintf("分类%s创建成功", name))
}

func DeleteCategory(cxt *gin.Context) {
	DB := common.GetDB()
	var userCategory model.Category
	cxt.Bind(&userCategory)

	if len(userCategory.Name) == 0 {
		response.Response(cxt, http.StatusUnprocessableEntity, 400, nil, "分类名称不能为空")
		return
	}

	var dbCategory model.Category
	DB.Where("name = ?", userCategory.Name).First(&dbCategory)
	if dbCategory.ID == 0 {
		response.Response(cxt, http.StatusUnprocessableEntity, 400, nil, fmt.Sprintf("分类名称%s不存在", userCategory.Name))
		return
	}

	if err := DB.Delete(&dbCategory).Error; err != nil {
		response.Response(cxt, http.StatusUnprocessableEntity, 400, nil, fmt.Sprintf("%s 删除失败", userCategory.Name))
	}
	response.Success(cxt, nil, fmt.Sprintf("%s 删除成功", dbCategory.Name))
}

func Update(cxt *gin.Context) {
	DB := common.GetDB()
	var requestCategory model.Category
	cxt.Bind(&requestCategory)
	if len(requestCategory.Name) == 0 {
		response.Fail(cxt, nil, "分类名称不能为空")
		return
	}
	var dbCategory model.Category
	DB.Where("name = ?", requestCategory.Name).First(&dbCategory)
	if dbCategory.ID == 0 {
		response.Fail(cxt, nil, fmt.Sprintf("分类%s不存在", requestCategory.Name))
		return
	}
	dbCategory.Instruction = requestCategory.Instruction
	if err := DB.Save(&dbCategory).Error; err != nil {
		response.Fail(cxt, nil, err.Error())
		return
	}
	response.Success(cxt, gin.H{"msg": "更新成功", "data": dbCategory}, "更新成功")
}

func GetCategories(cxt *gin.Context) {
	DB := common.GetDB()
	categories := []model.Category{}
	//DB.Not("deleted_at = ?", "").Find(&users)
	DB.Find(&categories)
	response.Success(cxt, gin.H{"data": categories}, "")
}
