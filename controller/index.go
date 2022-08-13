package controller

import (
	"metrics/common"
	"metrics/utils/response"
	"metrics/utils/route"
)

func Index(c *route.RouteContext) *response.Response {
	return response.Resp().Json(common.Success)
}
