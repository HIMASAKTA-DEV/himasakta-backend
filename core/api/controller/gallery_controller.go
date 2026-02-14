package controller

import (
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GalleryController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type galleryController struct {
	galleryService service.GalleryService
	s3Storage      storage.AwsS3
}

func NewGallery(galleryService service.GalleryService, s3Storage storage.AwsS3) GalleryController {
	return &galleryController{galleryService, s3Storage}
}

func (c *galleryController) Create(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		response.NewFailedWithCode(400, "image file is required", err).Send(ctx)
		return
	}

	// Validate file size (e.g., max 50MB)
	if file.Size > 50*1024*1024 {
		response.NewFailedWithCode(400, "file size exceeds 50MB limit", nil).Send(ctx)
		return
	}

	caption := ctx.PostForm("caption")
	category := ctx.PostForm("category")
	deptIdStr := ctx.PostForm("department_id")
	progIdStr := ctx.PostForm("progenda_id")

	var deptId, progId *uuid.UUID
	if deptIdStr != "" {
		id, err := uuid.Parse(deptIdStr)
		if err == nil {
			deptId = &id
		}
	}
	if progIdStr != "" {
		id, err := uuid.Parse(progIdStr)
		if err == nil {
			progId = &id
		}
	}

	// Upload to S3
	filename := fmt.Sprintf("%s-%s", uuid.New().String(), file.Filename)
	objectKey, err := c.s3Storage.UploadFile(filename, file, "gallery",
		"image/jpeg", "image/png", "image/jpg", "image/webp", "image/gif", "image/heif", "image/heic", "image/bmp", "image/tiff")
	if err != nil {
		response.NewFailed("failed to upload image to storage", err).Send(ctx)
		return
	}

	imageUrl := c.s3Storage.GetPublicLink(objectKey)

	req := dto.CreateGalleryRequest{
		ImageUrl:     imageUrl,
		Caption:      caption,
		Category:     category,
		DepartmentId: deptId,
		ProgendaId:   progId,
	}

	result, err := c.galleryService.Create(ctx.Request.Context(), req)
	if err != nil {
		// Attempt to cleanup uploaded file if db creation fails
		_ = c.s3Storage.DeleteFile(objectKey)
		response.NewFailed("failed create gallery", err).Send(ctx)
		return
	}

	response.NewSuccessCreated("success create gallery", result).Send(ctx)
}

func (c *galleryController) GetAll(ctx *gin.Context) {
	caption := ctx.Query("caption")
	galleries, metaRes, err := c.galleryService.GetAll(ctx.Request.Context(), meta.New(ctx), caption)
	if err != nil {
		response.NewFailed("failed get galleries", err).Send(ctx)
		return
	}

	response.NewSuccess("success get galleries", galleries, metaRes).Send(ctx)
}

func (c *galleryController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.galleryService.GetById(ctx.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed get gallery", err).Send(ctx)
		return
	}

	response.NewSuccess("success get gallery", result).Send(ctx)
}

func (c *galleryController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateGalleryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}

	result, err := c.galleryService.Update(ctx.Request.Context(), id, req)
	if err != nil {
		response.NewFailed("failed update gallery", err).Send(ctx)
		return
	}

	response.NewSuccess("success update gallery", result).Send(ctx)
}

func (c *galleryController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.galleryService.Delete(ctx.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed delete gallery", err).Send(ctx)
		return
	}

	response.NewSuccess("success delete gallery", nil).Send(ctx)
}
