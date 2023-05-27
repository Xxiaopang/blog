package controller

import (
	"fmt"
	"os"
	"strconv"

	"project/service"
	"project/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Fail(c, nil, "form file error:"+err.Error())
		zap.S().Info("form file error:" + err.Error())
		return
	}
	dir, _ := os.Getwd()
	newFileName := uuid.New().String() + file.Filename
	dst := dir + "/upload/" + newFileName
	go func() {
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			utils.Fail(c, nil, "file save error:"+err.Error())
			zap.S().Info("file save error:" + err.Error())
			return
		}
		if err = service.CreateFile(file.Filename, newFileName); err != nil {
			utils.Fail(c, nil, "文件上传失败")
			return
		}
		//messageQueue <- file.Filename
	}()
	utils.Success(c, nil, "文件正在上传")
}

func DownloadFile(c *gin.Context) {
	fileID := c.Param("id")
	id, err := strconv.ParseUint(fileID, 10, 64)
	if err != nil {
		zap.S().Info("string to uint fail" + err.Error())
		utils.Fail(c, nil, "参数转换错误:"+err.Error())
		return
	}
	fileDto, err := service.FindFileByID(uint(id))

	if err != nil || fileDto.OldName == "" {
		//zap.S().Info("select fileName error:" + err.Error())
		utils.Fail(c, nil, "查询失败:")
		return
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileDto.OldName))
	c.Writer.Header().Set("Content-Type", "application/zip")

	dir, _ := os.Getwd()

	dst := dir + "/upload/" + fileDto.NewName
	c.File(dst)
	utils.Success(c, nil, "文件下载成功")
}

func FindAllFile(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		utils.Fail(c, nil, "数据转换失败")
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		utils.Fail(c, nil, "数据转换失败")
		return
	}
	list, i := service.FindAllFile(limit, page)
	utils.Success(c, list, i)
}
