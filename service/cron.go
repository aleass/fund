package service

import (
	cron "github.com/robfig/cron/v3"
)

// 定时
func InitTask() {
	list := FundList{}
	star := FundStar{}
	earnings := FundEarnings{}
	//earningsRank := FundEaringsRank{}
	buySell := FundBuySell{}
	task := daysPastTimeRank{}
	purchase := FundPurchase{}
	data := FundData{}
	//初始化更新
	initFund := func() {
		println("start")
		list.GetData()
		data.GetData()
	}
	initFund()

	//触发定时
	c := cron.New()
	//阶段收益 1点
	_, err := c.AddFunc("0 11 * * 2-6", earnings.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//净值
	_, err = c.AddFunc("15 11 * * 2-6", data.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//已买基金
	_, err = c.AddFunc("30 11 * * 2-6", purchase.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//交易日 0点
	_, err = c.AddFunc("0 0 * * 2-6", TradeDay)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//基金购买情况 0点
	_, err = c.AddFunc("10 0 * * 2-6", buySell.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//基金购买情况 2点
	_, err = c.AddFunc("0 2 * * 2-6", task.Send)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//收益排行 0点
	//_, err = c.AddFunc("0 0 * * 2-6", earningsRank.GetData)
	//if err != nil {
	//	panic("cron err :" + err.Error())
	//}

	//基金评级任务 每个月
	_, err = c.AddFunc("0 0 * */1 *", star.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}

	//基金列表 每个月
	_, err = c.AddFunc("0 0 * */1 *", list.GetData)
	if err != nil {
		panic("cron err :" + err.Error())
	}
	c.Run()
}
