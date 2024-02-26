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
	var err = common.FuncDb.Model(&model.DfFundEarnings{}).Select("id,code").Find(&bufferEarnings).Error
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}

	var earningsMap = make(map[string]int64, len(bufferEarnings))
	for _, v := range bufferEarnings {
		earningsMap[v.Code] = v.Id
	}

	bufferEarnings = bufferEarnings[:0]

	var earList [][]string
	json.Unmarshal(raw, &earList)
	now := time.Now()

	var df []model.DfFundList
	err = common.FuncDb.Model(&model.DfFundList{}).Select("name,code").Where("deleted_at is null").Find(&df).Error
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}

	var codeMap = make(map[string]string, len(df))
	for _, fund := range df {
		codeMap[fund.Code] = fund.Name
	}

	updateBuff := strings.Builder{}
	for _, val := range earList {
		if val[3] == "" {
			continue
		}
		var earnings = model.DfFundEarnings{
			Name:           codeMap[val[0]],
			LastUpdateTime: now.Unix(),
			UpdatedAt:      &now,
			NavPerUnit:     common.Int642Float64(val[3]),
			Date:           now.Format("2006-01-02"),
		}

		if id, ok := earningsMap[earnings.Code]; ok {
			earnings.Id = id
			earnings.Code = val[0]
			updateEarnings = append(updateEarnings, earnings)
			continue
		}
		earnings.AddTime = now.Unix()
		earnings.CreatedAt = &now
		earnings.DailyGrowthRate = "0"
		earnings.CumulativeNav = "0"
		earnings.Past1Month = "0"
		earnings.Past1Week = "0"
		earnings.Past1Year = "0"
		earnings.Past2Years = "0"
		earnings.Past3Months = "0"
		earnings.Past3Years = "0"
		earnings.Past6Months = "0"
		earnings.SinceInception = "0"
		earnings.ThisYear = "0"
		bufferEarnings = append(bufferEarnings, earnings)

	}
	if len(bufferEarnings) > 0 {
		var err = common.FuncDb.CreateInBatches(&bufferEarnings, 1).Error
		if err != nil {
			log.Println(err.Error())
			common.Logger.Error(err.Error())
			return
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
