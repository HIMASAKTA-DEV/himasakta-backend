package config

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/storage"
	"github.com/gin-gonic/gin"
)

func NewRouter(server *gin.Engine, s3 storage.AwsS3) *gin.Engine {
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Route Not Found",
		})
	})

	server.MaxMultipartMemory = 30 * 1024 * 1024
	server.Use(customRecovery())
	server.Use(middleware.SecurityMiddleware())
	server.Use(middleware.CORSMiddleware())

	server.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong 123",
		})
	})
	// server.POST("/api/v1/uploads", func(ctx *gin.Context) {
	// 	file, err := ctx.FormFile("file")
	// 	if err != nil {
	// 		response.NewFailed("failed to get file", err).SendWithAbort(ctx)
	// 		return
	// 	}

	// 	filename := fmt.Sprintf("assets-%s.%s", ulid.Make(), utils.GetExtensions(file.Filename))

	// 	// Map the folders based on extensions or just one folder
	// 	folderName := "others"
	// 	ext := strings.ToLower(utils.GetExtensions(file.Filename))
	// 	if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "webp" {
	// 		folderName = "images"
	// 	}

	// 	objectKey, err := s3.UploadFile(filename, file, folderName)
	// 	if err != nil {
	// 		response.NewFailed("failed upload to s3", err).SendWithAbort(ctx)
	// 		return
	// 	}

	// 	fileURL := s3.GetPublicLink(objectKey)

	// 	response.NewSuccess("success upload image", gin.H{
	// 		"url":  fileURL,
	// 		"path": objectKey,
	// 	}).Send(ctx)
	// })

	server.Static("/api/static", "./public/uploads")
	return server
}

func customRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var wrappedErr error
				if e, ok := err.(error); ok {
					wrappedErr = e
				} else {
					wrappedErr = fmt.Errorf("%v", err)
				}

				fmt.Println(mylog.ColorizePanic(fmt.Sprintf("\n[Recovery] Panic occurred: %v\n", err)))
				stack := debug.Stack()
				coloredStack := mylog.ColorizePanic(string(stack))

				fmt.Fprintln(os.Stderr, coloredStack)
				response.NewFailed("server panic occured", wrappedErr).
					SendWithAbort(ctx)
			}
		}()

		ctx.Next()
	}
}
