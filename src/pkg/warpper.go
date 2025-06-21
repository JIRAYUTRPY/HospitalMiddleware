package pkg

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal server error")
)

func WrapperOK() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.JSON(http.StatusOK, map[string]any{
			"status": true,
		})
	}
}

func WrapperError() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, map[string]any{
				"reason": err.Error(),
			})
		}
	}
}

// func WrapperError(c *gin.Context, err error) {
// 	response := gin.H{
// 		"status": false,
// 		"error":  err.Error(),
// 	}
// 	switch err {
// 	case ErrNotFound:
// 		response["error"] = "Not Found"
// 		c.JSON(http.StatusNotFound, response)
// 		c.Abort()
// 	case ErrBadRequest:
// 		response["error"] = "Bad Request"
// 		c.JSON(http.StatusBadRequest, response)
// 		c.Abort()
// 	case ErrInternal:
// 		response["error"] = "Internal Server Error"
// 		c.JSON(http.StatusInternalServerError, response)
// 		c.Abort()
// 	default:
// 		c.JSON(http.StatusInternalServerError, response)
// 		c.Abort()
// 	}
// }
