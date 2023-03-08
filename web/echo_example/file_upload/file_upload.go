package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

const dstDir = "D:/Users/Jake/go/web/echo_example/file_upload/files/"

func main() {
	e := echo.New()

	e.POST("/upload", func(ctx echo.Context) error {
		username := ctx.FormValue("username")

		file, err := ctx.FormFile("avatar")
		if err != nil {
			return err
		}
		log.Printf("file#name >>> %s", file.Filename)

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dstFile := dstDir + file.Filename
		dst, err := os.Create(dstFile)
		if err != nil {
			return err
		}
		defer dst.Close()
		log.Printf("dst#name >>> %s", dst.Name())
		if absPath, err := filepath.Abs(dst.Name()); err != nil {
			return err
		} else {
			log.Printf("dst#abspath >>> %s", absPath)
		}

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, echo.Map{
			"method":  ctx.Request().Method,
			"route":   ctx.Path(),
			"message": username + " uploaded a file  successfully",
		})
	})

	e.Logger.Fatal(e.Start(":80"))
}
