package pkg

import (
	"julianmindria/backendCoffee/config"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data,omitempty"`
	Meta        interface{} `json:"meta,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
	ctx.Abort()
	return
}

func NewRes(code int, data *config.Result) *Response {
	var respone = Response{
		Code:   code,
		Status: getStatus(code),
	}

	if respone.Code >= 400 {
		respone.Description = data.Data
	} else if data.Message != nil {
		respone.Description = data.Message
	} else {
		respone.Data = data.Data
	}

	if data.Meta != nil {
		respone.Meta = data.Meta
	}

	return &respone
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
		break
	case 201:
		desc = "Created"
		break
	case 400:
		desc = "Bad Request"
		break
	case 401:
		desc = "Unauthorized"
		break
	case 403:
		desc = "Forbidden"
		break
	case 404:
		desc = "Not Found"
		break
	case 500:
		desc = "Internal Server Error"
		break
	case 501:
		desc = "Bad Gateway"
		break
	case 304:
		desc = "Not Modified"
		break
	default:
		desc = ""
	}

	return desc
}
