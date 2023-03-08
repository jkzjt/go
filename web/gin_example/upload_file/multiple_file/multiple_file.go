package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const dst_dir = "d:/Users/Jake/go/web/gin_example/upload_file/files"

func main() {

	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		username := c.PostForm("username")

		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		files := form.File["files"]

		for i, file := range files {
			filename := filepath.Base(file.Filename)
			dst_file := dst_dir + "/" + filename
			if err := c.SaveUploadedFile(file, dst_file); err != nil {
				c.String(http.StatusBadRequest, "upload %dth file `%s` err: %s", i+1, file.Filename, err.Error())
				return
			}
		}

		c.String(http.StatusOK, "%s uploaded total %d files successfully", username, len(files))
	})

	r.Run(":80")
}
