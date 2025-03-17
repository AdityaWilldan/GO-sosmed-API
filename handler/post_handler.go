package handler

import (
	"GoSosmed/dto"
	"GoSosmed/errorhandler"
	"GoSosmed/helper"
	"GoSosmed/service"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	return &postHandler{
		service: service,
	}

}

func (h *postHandler) Create(ctx *gin.Context) {
	var post dto.PostRequest

	if err := ctx.ShouldBind(&post); err != nil {
		errorhandler.HandleError(ctx, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	homeDir, _ := os.UserHomeDir()
	uploadPath := filepath.Join(homeDir, "Documents/GoSosmed/public/picture")
	if post.Picture != nil {
		if err := os.MkdirAll(uploadPath, 0755); err != nil {
			errorhandler.HandleError(ctx, &errorhandler.InternalServerError{
				Message: err.Error(),
			})
			return
		}
		//RENAME GAMBAR
		ext := filepath.Ext(post.Picture.Filename)
		NewFileName := uuid.New().String() + ext

		//SAVE IMG DIRECTORY
		dst := filepath.Join("./public/picture", filepath.Base(NewFileName))
		ctx.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", ctx.Request.Host, NewFileName)
	}

	userID, _ := ctx.Get("userID")
	post.UserID = userID.(int)
	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Success Post Your Tweet",
	})
	ctx.JSON(http.StatusCreated, res)
}
