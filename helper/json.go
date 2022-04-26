package helper

import (
	"encoding/json"
	"net/http"
	"thirthfamous/golang-restful-api-clean-architecture/model/web"

	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(statusCode int, ctx *gin.Context, response interface{}) {

	ctx.JSON(statusCode, web.WebResponse{
		Code:   statusCode,
		Status: StatusCode(statusCode),
		Data:   response,
	})
}

func StatusCode(httpStatus int) string {
	switch httpStatus {
	case http.StatusOK:
		return "OK"
	case http.StatusNotFound:
		return "NOT FOUND"
	case http.StatusBadRequest:
		return "BAD REQUEST"
	default:
		return "INTERNAL SERVER ERROR"
	}
}
