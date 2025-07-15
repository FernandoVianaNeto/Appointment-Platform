package web

import (
	"appointment-platform-backend-backend/internal/application/exceptions"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, status int, response any) {
	c.JSON(status, response)
}

func NoContentResponse(c *gin.Context, status int) {
	c.Status(status)

}
func ErrorResponse(c *gin.Context, err error) {
	status, msg := errorTreatment(c, err)

	c.JSON(status, msg)
}

func errorTreatment(ctx context.Context, err error) (int, interface{}) {
	switch e := err.(type) {
	case exceptions.ApplicationError:
		return e.Code(), e.Message(ctx)
	default:
		return http.StatusInternalServerError, gin.H{
			"message": "unexpected error",
		}
	}
}
