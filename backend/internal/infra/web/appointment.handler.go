package web

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/infra/web/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateAppointmentHandler(ctx *gin.Context) {
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

	if form.Get("start_date") == "" ||
		form.Get("end_date") == "" ||
		form.Get("patient_uuid") == "" ||
		form.Get("insurance") == "" ||
		form.Get("technician") == "" ||
		form.Get("location") == "" ||
		form.Get("procedure") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	createAppointmentDto := dto.CreateAppointmentInputDto{
		UserUuid:    userUuid,
		StartDate:   form.Get("start_date"),
		EndDate:     form.Get("end_date"),
		PatientUuid: form.Get("patient_uuid"),
		Insurance:   form.Get("insurance"),
		Technician:  form.Get("technician"),
		Location:    form.Get("location"),
		Procedure:   form.Get("procedure"),
	}

	err = s.CreateAppointmentUsecase.Execute(ctx, createAppointmentDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) ListAppointmentsHandler(ctx *gin.Context) {
	var queryParams requests.ListAppointmentRequest

	value := ctx.Value("user_uuid")

	userUuid, ok := value.(string)

	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found in context"})
		return
	}

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
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

	response, err := s.ListAppointmentUsecase.Execute(ctx, dto.ListAppointmentInputDto{
		UserUuid:    userUuid,
		Page:        page,
		SearchInput: &queryParams.SearchTerm,
		FilterType:  &queryParams.FilterType,
		Date:        &queryParams.Date,
	})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *Server) EditAppointmentHandler(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	form := ctx.Request.Form

	if form.Get("uuid") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing appointment uuid"})
		return
	}

	editAppointmentDto := dto.EditAppointmentInputDto{
		Uuid: form.Get("uuid"),
	}

	if form.Get("status") != "" {
		editAppointmentDto.Status = ptr(form.Get("status"))
	}

	if form.Get("start_date") != "" {
		editAppointmentDto.StartDate = ptr(form.Get("start_date"))
	}

	if form.Get("end_date") != "" {
		editAppointmentDto.EndDate = ptr(form.Get("end_date"))
	}

	if form.Get("procedure") != "" {
		editAppointmentDto.Procedure = ptr(form.Get("procedure"))
	}

	err = s.EditAppointmentUsecase.Execute(ctx, editAppointmentDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) DeleteAppointmentHandler(ctx *gin.Context) {
	var req requests.DeleteAppointmentRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request Uri"})
		return
	}

	err := s.DeleteAppointmentUsecase.Execute(ctx, dto.DeleteAppointmentInputDto{Uuid: &req.Uuid})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
