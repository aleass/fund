package service

import (
	"bytes"
	"encoding/json"
	"fund/common"
	"fund/model"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
阶段收益
*/

var (
	dataFormat = []byte("datas:[")
)

type FundData struct {
}

func (f *FundData) GetData() {
	common.Logger.Info("执行 阶段收益任务")

	refer := [][2]string{
		{"Referer", "https://fund.eastmoney.com/ZQ_jzzzl.html"},
	}
	res, err := common.HttpRequest(http.MethodGet, common.DataUrl, nil, refer)
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}

	if len(res) == 0 {
		return
	}

	if index := bytes.Index(res, dataFormat); index != -1 {
		res = res[index+len(dataFormat)-1:]
		if index2 := bytes.Index(res, []byte("]]")); index2 != -1 {
			res = res[:index2+2]
			f.extract(res)
		}
	}
}

func (f *FundData) extract(raw []byte) {
	var bufferEarnings []model.DfFundEarnings
	var updateEarnings []model.DfFundEarnings
	common.FuncDb.Model(&model.DfFundEarnings{}).Find(&bufferEarnings)
	var earningsMap = make(map[string]int64, len(bufferEarnings))
	for _, v := range bufferEarnings {
		earningsMap[v.Code] = v.Id
	}

	bufferEarnings = bufferEarnings[:0]

	var earList [][]string
	json.Unmarshal(raw, &earList)
	now := time.Now()

	var df []model.DfFundList
	common.FuncDb.Model(&model.DfFundList{}).Select("id,code").Where("Inc_date is null").Find(&df)
	var codeMap = make(map[string]int64, len(df))
	for _, fund := range df {
		codeMap[fund.Code] = fund.Id
	}

	updateBuff := strings.Builder{}
	for _, val := range earList {
		if val[3] == "" {
			continue
		}
		var earnings = model.DfFundEarnings{
			Code:            val[0],
			LastUpdateTime:  now.Unix(),
			UpdatedAt:       &now,
			NavPerUnit:      common.Int642Float64(val[3]),
			DailyGrowthRate: "0",
			CumulativeNav:   "0",
			Past1Month:      "0",
			Past1Week:       "0",
			Past1Year:       "0",
			Past2Years:      "0",
			Past3Months:     "0",
			Past3Years:      "0",
			Past6Months:     "0",
			SinceInception:  "0",
			ThisYear:        "0",

			Date: now.Format("2006-01-02"),
		}

		if id, ok := earningsMap[earnings.Code]; ok {
			earnings.Id = id
			updateEarnings = append(updateEarnings, earnings)
			continue
		}
		earnings.AddTime = now.Unix()
		earnings.CreatedAt = &now
		bufferEarnings = append(bufferEarnings, earnings)

	}
	if len(bufferEarnings) > 0 {
		var err = common.FuncDb.CreateInBatches(&bufferEarnings, 1).Error
		if err != nil {
			log.Println(err.Error())
			common.Logger.Error(err.Error())
		}
	}
	if len(updateEarnings) > 0 {
		for _, earning := range updateEarnings {
			var err = common.FuncDb.Updates(&earning).Error
			if err != nil {
				log.Println(err.Error())
				common.Logger.Error(err.Error())
			}
		}

	}

	if updateBuff.Len() != 0 {
		var err = common.FuncDb.Exec(updateBuff.String()).Error
		if err != nil {
			log.Println(err.Error())
			common.Logger.Error(err.Error())
		}
	}

}
