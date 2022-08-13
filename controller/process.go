package controller

import (
	"errors"
	"metrics/common"
	"metrics/model/core"
	"metrics/model/service"
	logs "metrics/utils/log"
	"metrics/utils/response"
	"metrics/utils/route"
	"strconv"
	"time"
)

// AutoCreateProcess 自动添加所有运行的进程
func AutoCreateProcess(c *route.RouteContext) *response.Response {
	err := service.AutoCreateProcess()
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}
	return response.Resp().Json(common.Success)
}

// FetchAllRunningProcess 获取所有正在运行的进程
func FetchAllRunningProcess(c *route.RouteContext) *response.Response {
	res, err := service.FetchAllRunningProcess()
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}
	common.Success.Data = res
	return response.Resp().Json(common.Success)
}

// FetchProcess 获取所有监控进程
func FetchProcess(c *route.RouteContext) *response.Response {
	res, err := service.FetchAllMonitorProcess()
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}
	common.Success.Data = res
	return response.Resp().Json(common.Success)
}

// UpdateProcess 更新监控进程
func UpdateProcess(c *route.RouteContext) *response.Response {
	var data core.Process = core.Process{}
	err := c.BindJSON(&data)
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}

	data.UpdateTime = int(time.Now().Unix())
	err = service.UpdateMonitorProcess(data)
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}
	return response.Resp().Json(common.Success)
}

// DeleteProcess 删除监控进程
func DeleteProcess(c *route.RouteContext) *response.Response {
	idStr := c.Query("id")
	if idStr == "" {
		logs.Error(errors.New("参数错误"))
		return response.Resp().Json(common.ParamInvalid)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}
	err = service.DeleteMonitorProcessByID(id)
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}
	return response.Resp().Json(common.Success)
}

// AddProcess 增加监控进程
func AddProcess(c *route.RouteContext) *response.Response {
	var data core.Process = core.Process{}
	err := c.BindJSON(&data)
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}

	data.UpdateTime = int(time.Now().Unix())
	err = service.CreateMonitorProcess(data)
	if err != nil {
		logs.Error(err)
		return response.Resp().Json(common.Failure)
	}
	return response.Resp().Json(common.Success)
}
