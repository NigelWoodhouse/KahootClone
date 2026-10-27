package controller

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quiz.com/quiz/internal/service"
)

type ImageController struct {
	imageService *service.ImageService
}

func Image(imageService *service.ImageService) ImageController {
	return ImageController{
		imageService: imageService,
	}
}

func (c ImageController) UploadImage(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		fileHeader, err = ctx.FormFile("file")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Image file is required under 'image' or 'file' form field",
			})
		}
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer file.Close()

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" || contentType == "application/octet-stream" {
		ext := strings.ToLower(fileHeader.Filename)
		if strings.HasSuffix(ext, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(ext, ".jpg") || strings.HasSuffix(ext, ".jpeg") {
			contentType = "image/jpeg"
		} else if strings.HasSuffix(ext, ".gif") {
			contentType = "image/gif"
		} else if strings.HasSuffix(ext, ".webp") {
			contentType = "image/webp"
		} else if strings.HasSuffix(ext, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "image/png"
		}
	}

	id, err := c.imageService.UploadImage(fileHeader.Filename, contentType, file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":  id.Hex(),
		"url": fmt.Sprintf("/api/images/%s", id.Hex()),
	})
}

func (c ImageController) GetImage(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid image ID")
	}

	stream, contentType, length, err := c.imageService.DownloadImage(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Image not found")
	}

	ctx.Set("Content-Type", contentType)
	ctx.Set("Cache-Control", "public, max-age=31536000")
	if length > 0 {
		return ctx.SendStream(stream, int(length))
	}
	return ctx.SendStream(stream)
}

func (c ImageController) DeleteImage(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid image ID")
	}

	if err := c.imageService.DeleteImage(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
