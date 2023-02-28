package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func FileController(c *gin.Context) {
	file, err := c.FormFile("raw")
	if err != nil {
		log.Fatal(err)
	}
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)

	filename := uuid.New().String()
	uploads := filepath.Join(dir, "uploads")
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		return
	}
	fullpath := path.Join("uploads", filename+filepath.Ext(file.Filename))
	fillErr := c.SaveUploadedFile(file, filepath.Join(dir, fullpath))
	if fillErr != nil {
		log.Fatal(fillErr)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
}
