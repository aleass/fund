package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weather/common"
	"weather/model"
)

type TradeDayResp struct {
	Data struct {
		Klines []string `json:"klines"`
	} `json:"data"`
}

// 更新平均-收益差值数据
func UpGainData() {
	sql := "UPDATE `fund`.`df_fund_earnings_rank` SET `gain` = total_rate -kind_avg_rate where gain != 0"
	err := common.FuncDb.Exec(sql).Error
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}
}

// 更新交易日
func TradeDay() {
	common.Logger.Info("执行 获取交易日")
	var db model.TradeDay
	err := common.FuncDb.Model(&model.TradeDay{}).Order("date desc").Limit(1).Find(&db).Error
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}

	var beg int64 = 0
	if db.Date != 0 {
		beg = db.Date
	}
	url := fmt.Sprintf(common.TradeDayUrl, beg)
	raw, err := common.HttpRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}
	var resp TradeDayResp
	json.Unmarshal(raw, &resp)

	_date := make([]byte, 8)
	var tradeList = make([]model.TradeDay, 0, len(resp.Data.Klines))
	for _, date := range resp.Data.Klines {
		i := copy(_date, date[:4])
		i += copy(_date[i:], date[5:7])
		copy(_date[i:], date[8:])
		var dateInt = common.Str2Int64(string(_date))
		if db.Date == dateInt {
			continue
		}
		tradeList = append(tradeList, model.TradeDay{dateInt})
	}
	err = common.FuncDb.CreateInBatches(tradeList, 10000).Error
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
	}

	common.Logger.Info("结束 获取交易日")
}
