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
	{"拉取0,1,2数据,发送已买基金盈利情况", GetPurDataEarningsAll},
}

func GetPurDataEarningsAll() {
	Earnings.GetData()
	Earnings.CumulativeNav()
	Earnings.GetPurData()
	Purchase.GetData()
}
