package web

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/infra/web/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreatePatientHandler(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	form := ctx.Request.Form

	email := form.Get("email")
	address := form.Get("address")
	insurance := form.Get("insurance")

	createPatientDto := dto.CreatePatientInputDto{
		Email:     &email,
		Name:      form.Get("name"),
		Phone:     form.Get("phone"),
		Address:   &address,
		Insurance: &insurance,
	}

	err = s.CreatePatientUsecase.Execute(ctx, createPatientDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) ListPatientHandler(ctx *gin.Context) {
	var req requests.GetByUuidRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request Uri"})
		return
	}

	response, err := s.GetUserUsecase.Execute(ctx, dto.GetUserInputDto{Uuid: req.Uuid})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		ctx.JSON(http.StatusNotFound, "User not found")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *Server) DeletePatientHandler(ctx *gin.Context) {
	var req requests.GetByUuidRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request Uri"})
		return
	}

	response, err := s.GetUserUsecase.Execute(ctx, dto.GetUserInputDto{Uuid: req.Uuid})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		ctx.JSON(http.StatusNotFound, "User not found")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *Server) UpdatePatientHandler(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	form := ctx.Request.Form

	value := ctx.Value("user_uuid")

	userUuid, ok := value.(string)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User uuid not found in context"})
		return
	}

	editUserDto := dto.UpdateUserInputDto{
		Uuid: userUuid,
	}
	if form.Get("name") != "" {
		editUserDto.Name = ptr(form.Get("name"))
	}

	err = s.UpdateUserUsecase.Execute(ctx, editUserDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
