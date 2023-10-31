package helpers

import "github.com/gin-gonic/gin"

type response struct {
	Data interface{} `json:"data"`
}

type responseError struct {
	Message interface{} `json:"message"`
}

func RespondJSON(c *gin.Context, statusCode int, data interface{}, meta interface{}) {
	res := response{
		Data: data,
	}

	c.JSON(statusCode, res)
}

func RespondErrorJSON(c *gin.Context, statusCode int, err error) {
	res := responseError{
		Message: err.Error(),
	}

	c.JSON(statusCode, res)
}
