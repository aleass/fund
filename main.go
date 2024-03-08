package main

import (
	"bufio"
	"fund/common"
	"fund/service"
	"os"
	"strings"
)

func main() {
	//初始化配置
	common.InitConfig()

	//命令启动
	if len(os.Args) > 1 && os.Args[1] == "help" {
		var command string
		reader := bufio.NewReader(os.Stdin)
		println(`请输入:`)
		for {
			for i, v := range service.GetTaskList() {
				println(i, v.Title)
			}
			command, _ = reader.ReadString('\n')
			//清楚无效字节
			command = strings.ReplaceAll(command, "\r", "")
			command = strings.ReplaceAll(command, "\n", "")
			command = strings.TrimSpace(command)

			//获取任务
			task := service.GetTaskByIndex(common.Str2Int(command))
			if task == nil {
				continue
			}

			println("执行开始:", command, task.Title)
			task.Func()
			println("执行完成:", command, task.Title)
			println("\r\r")
		}
	}

	//任务启动
	service.InitTask()
}
