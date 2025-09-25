package controller

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	disk := c.Param("disk")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	// Ambil ekstensi file asli
	ext := filepath.Ext(file.Filename)
	// Generate nama baru pakai MD5 hash dari timestamp + nama asli biar unik
	hash := md5.New()
	hash.Write([]byte(file.Filename + time.Now().String()))
	hashed := hex.EncodeToString(hash.Sum(nil)) + ext

	basePath, _ := os.Getwd()
	path := basePath + "/storage/" + disk
	filename := filepath.Join(path, hashed)
	// filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": true, "data": hashed, "message": "Success"})
}

func Encode(c *gin.Context) {
	filename := c.Param("filename")
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(filepath.FromSlash("storage/public/" + filename))
	if err != nil {
		// log.Fatal(err)
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	// fmt.Println(base64Encoding)
	c.JSON(201, gin.H{"status": true, "data": base64Encoding, "message": nil})
}

func FileStream(c *gin.Context) {
	filename := c.Param("filename")
	basePath, _ := os.Getwd()
	c.File(basePath + "/storage/public/" + filename)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
