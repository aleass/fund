package main

import (
	"weather/common"
	"weather/service"
)

func main() {
	//初始化配置
	common.InitConfig()

	//任务启动
	service.InitTask()
}
