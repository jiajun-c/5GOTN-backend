package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Train(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed"})
		return
	}

	filename := filepath.Base(file.Filename)
	fmt.Println("Upload the file", filename)
	// 将文件保存到 "./uploaded-files" 目录下
	err = c.SaveUploadedFile(file, fmt.Sprintf("./%s", filename))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
