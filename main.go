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
				println("执行:阶段收益")
				service.Earnings.GetData()
			case "2": //净值
				println("执行:净值")
				service.Data.GetData()
			case "3": //已买基金
				println("执行:已买基金")
				service.Purchase.GetData()
			case "q":
				return
			}
			return
		}
	}

	//任务启动
	service.InitTask()
}
