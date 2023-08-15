package middleware

import (
	"julianmindria/backendCoffee/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("image_file")
	if err != nil {
		if err.Error() == "http: no such file" {
			ctx.Set("image", "")
			ctx.Next()
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing file"})
		return
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	result, err := pkg.Cloudinary(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	ctx.Set("image", result)
	ctx.Next()
}
