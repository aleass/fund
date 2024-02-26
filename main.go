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
		for {

			println(`请输入:
1.阶段收益
2.净值
3.已买基金
4.已买基金阶段收益
q.退出
`)
			command, _ = reader.ReadString('\n')
			//清楚无效字节
			command = strings.ReplaceAll(command, "\r", "")
			command = strings.ReplaceAll(command, "\n", "")
			command = strings.TrimSpace(command)

			//任务启动
			switch command {
			case "1": //阶段收益
				println("执行:1.阶段收益")
				service.Earnings.GetData()

				println("执行完成:阶段收益")
			case "2": //净值
				println("执行:2.净值")
				service.Data.GetData()
				println("执行完成:净值")
			case "3": //已买基金
				println("执行:3.已买基金")
				service.Purchase.GetData()
				println("执行完成:已买基金")
			case "4": //已买基金阶段收益
				println("执行:4.已买基金阶段收益")
				service.Earnings.GetPurData()
				println("执行完成:已买基金阶段收益")
			case "q":
				return
			}
			println("\r\r")
		}
	}

	//任务启动
	service.InitTask()
}
