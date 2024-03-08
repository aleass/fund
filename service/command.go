package service

type Commons struct {
	Title string
	Func  func()
}

func GetTaskList() []Commons {
	return commandList
}

func GetTaskByIndex(index int) *Commons {
	if index < 0 || index-1 > len(commandList) {
		return nil
	}
	return &commandList[index]
}

var commandList = []Commons{
	{"基金阶段收益", Earnings.GetData},
	{"基金净值", Earnings.CumulativeNav},
	{"已买基金阶段收益", Earnings.GetPurData},
	{"发送:已买基金盈利情况", Purchase.GetData},
	{"更新:基金阶段收益 基金净值 发送:已买基金盈利情况", Earnings.GetPurData},
}
