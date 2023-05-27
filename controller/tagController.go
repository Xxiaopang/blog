package controller

import (
	"net/http"
	"project/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"project/service"
	"project/utils"
)

func FindAllTags(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		utils.Fail(c, nil, "查询失败")
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		utils.Fail(c, nil, "查询失败")
		return
	}
	articles, i := service.FindAllTags(limit, page)
	utils.Success(c, articles, i)
}

func CreateTags(c *gin.Context) {
	var tag vo.TagRequest
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "创建tag失败"+err.Error())
		zap.S().Info("创建tag失败" + err.Error())
		return
	}
	err = service.CreateTags(tag.Name)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	utils.Success(c, nil, "创建成功")
}

func DeleteTags(c *gin.Context) {

}
