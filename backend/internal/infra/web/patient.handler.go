package web

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/infra/web/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreatePatientHandler(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	value := ctx.Value("user_uuid")

	userUuid, ok := value.(string)

	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found in context"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	form := ctx.Request.Form

	if form.Get("name") == "" || form.Get("phone") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body. Missing params"})
		return
	}

	email := form.Get("email")
	address := form.Get("address")
	insurance := form.Get("insurance")

	createPatientDto := dto.CreatePatientInputDto{
		UserUuid:  userUuid,
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
	var queryParams requests.ListPatientRequest

	if err := ctx.ShouldBindUri(&queryParams); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request Uri"})
		return
	}

	page := 1

	if queryParams.Page != "" {
		pageInt, err := strconv.Atoi(queryParams.Page)

		if err == nil {
			page = pageInt
		}
	}

	response, err := s.ListPatientUsecase.Execute(ctx, dto.ListPatientInputDto{
		Page:        page,
		SearchInput: &queryParams.SearchTerm,
		FilterType:  &queryParams.FilterType,
	})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *Server) DeletePatientHandler(ctx *gin.Context) {
	var req requests.DeletePatientRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request Uri"})
		return
	}

	err := s.DeletePatientUsecase.Execute(ctx, dto.DeletePatientInputDto{Uuid: &req.Uuid})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) EditPatientHandler(ctx *gin.Context) {
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

	editPatientDto := dto.EditPatientInputDto{
		Uuid: userUuid,
	}

	if form.Get("name") != "" {
		editPatientDto.Name = ptr(form.Get("name"))
	}

	if form.Get("phone") != "" {
		editPatientDto.Phone = ptr(form.Get("phone"))
	}

	if form.Get("email") != "" {
		editPatientDto.Email = ptr(form.Get("email"))
	}

	if form.Get("address") != "" {
		editPatientDto.Address = ptr(form.Get("address"))
	}

	err = s.EditPatientUsecase.Execute(ctx, editPatientDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
