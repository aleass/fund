package service

import (
	cron "github.com/robfig/cron/v3"
)

var (
	List    = FundList{}
	Star    = FundStar{}
	BuySell = FundBuySell{}
	//Task     = daysPastTimeRank{}
	Purchase = FundPurchase{}
	Earnings = FundEarnings{}
	//EarningsRank = FundEaringsRank{}
)

// 定时
func InitTask() {

	//初始化更新
	println("start task")
	//触发定时
	c := cron.New()
	//阶段收益 1点
	_, err := c.AddFunc("0 23 * * 2-6", Earnings.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//净值
	_, err = c.AddFunc("15 23 * * 2-6", Earnings.CumulativeNav)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//已买基金盈利情况
	_, err = c.AddFunc("30 23 * * 2-6", Purchase.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//交易日 0点
	_, err = c.AddFunc("0 0 * * 2-6", TradeDay)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//交易日 0点
	_, err = c.AddFunc("0 1 * * 2-6", Earnings.GetPurData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	////基金购买情况 0点
	//_, err = c.AddFunc("10 0 * * 2-6", BuySell.GetData)
	//if err != nil {
	//	panic("cron err :" + err.Error())
	//}

	////基金购买情况 2点
	//_, err = c.AddFunc("0 2 * * 2-6", Task.Send)
	//if err != nil {
	//	panic("cron err :" + err.Error())
	//}

	//收益排行 0点
	//_, err = c.AddFunc("0 0 * * 2-6", EarningsRank.GetData)
	//if err != nil {
	//	panic("cron err :" + err.Error())
	//}

	//基金评级任务 每个月
	_, err = c.AddFunc("0 4 * */1 *", Star.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//基金列表 每个月
	_, err = c.AddFunc("0 3 * */1 *", List.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}
	c.Run()
}
