package handler

import (
	"fmt"

	backup "github.com/dignelidxdx/HackthonGo/internal/backup"
	"github.com/dignelidxdx/HackthonGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type BackUpHandler struct {
	service backup.BackUpService
}

func NewBackUp(s backup.BackUpService) *BackUpHandler {
	return &BackUpHandler{service: s}
}

func (backup *BackUpHandler) SaveFiles() gin.HandlerFunc {

	return func(context *gin.Context) {

		var request web.RequestBackUp
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(400, web.NewResponse(400, "", fmt.Sprintf("There was a error %v", err)))
		} else {

			err := backup.service.SaveElementToDB(request.NameFile)
			if err != nil {
				context.JSON(400, web.NewResponse(400, "", fmt.Sprintf("There was a error %v", err)))
			} else {
				context.JSON(200, web.NewResponse(200, "Se cargo item con exito", ""))
			}
		}

	}
}
