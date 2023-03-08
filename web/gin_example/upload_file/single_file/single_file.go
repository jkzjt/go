package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const dst_dir = "d:/Users/Jake/go/web/gin_example/upload_file/files"

func main() {
	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8Mib
	r.POST("/upload", func(c *gin.Context) {
		username := c.PostForm("username")

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get file err: %s", err.Error())
			return
		}

		filename := filepath.Base(file.Filename)
		dst_file := dst_dir + "/" + filename
		if err := c.SaveUploadedFile(file, dst_file); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "%s uploaded file `%s` to `%s` diretory successfully", username, file.Filename, dst_dir)
	})

	r.Run(":80")
}
