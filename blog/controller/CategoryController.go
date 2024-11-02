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
		response.Response(cxt, http.StatusInternalServerError, 500, nil, "数据入库失败")
		return
	}
	response.Response(cxt, http.StatusOK, 200, nil, fmt.Sprintf("分类%s创建成功", name))
}

func Delete(cxt *gin.Context) {

}

func Update(cxt *gin.Context) {

}

func GetCategoryById(cxt *gin.Context) {
	//DB := common.GetDB()
	var requestCategory model.Category
	cxt.Bind(&requestCategory)
	var id1 = cxt.Param("ID")
	var id2 = requestCategory.ID

	fmt.Println(id1, id2)

}
