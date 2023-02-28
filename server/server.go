package server

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github/shawncos/file-sync/server/controller"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	router.POST("/api/v1/files", controller.FileController)
	router.GET("/api/v1/qrcodes", controller.QrcodeController)
	router.GET("/api/v1/uploads/:path", controller.UploadsController)
	router.GET("/api/v1/addresses", controller.AddressController)
	router.POST("/api/v1/texts", controller.TextController)
	router.StaticFS("/static", http.FS(staticFiles))
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
		} else {
			c.Status(http.StatusNotFound)
		}
	})
	router.Run(":" + "27149")
}
