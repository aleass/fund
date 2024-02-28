package service

import (
	"bytes"
	"fmt"
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
	earningsFormat = []byte("var rankData = {datas:[")
)

type FundEarnings struct {
}

func (f *FundEarnings) GetData() {
	common.Logger.Info("执行 阶段收益任务")

	refer := [][2]string{
		{"Referer", "http://fund.eastmoney.com/data/fundranking.html"},
	}
	res, err := common.HttpRequest(http.MethodPost, common.EarningsUrl, nil, refer)
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}

	if len(res) == 0 {
		return
	}

	if index := bytes.Index(res, earningsFormat); index != -1 {
		res = res[index+len(earningsFormat)+1:]
		if index2 := bytes.IndexByte(res, ']'); index2 != -1 {
			f.extract(res[:index2-1])
		}
	}
	//购买的
	f.GetPurData()
}

func (f *FundEarnings) extract(data []byte) {
	var bufferEarnings []model.DfFundEarnings
	var updateEarnings []model.DfFundEarnings
	common.FuncDb.Model(&model.DfFundEarnings{}).Select("code,id").Find(&bufferEarnings)
	var earningsMap = make(map[string]int64, len(bufferEarnings))
	for _, v := range bufferEarnings {
		earningsMap[v.Code] = v.Id
	}

	bufferEarnings = bufferEarnings[:0]

	earList := bytes.Split(data, []byte(`","`))
	now := time.Now()
	nowDate := now.Format("2006-01-02 15:04:05")

	var df []model.DfFundList
	common.FuncDb.Model(&model.DfFundList{}).Select("id,code").Where("Inc_date is null").Find(&df)
	var codeMap = make(map[string]int64, len(df))
	for _, fund := range df {
		codeMap[fund.Code] = fund.Id
	}

	updateBuff := strings.Builder{}
	sql := "UPDATE `df_fund_list` SET `Inc_date`='%s',`date`='%s' WHERE `id` = %d;"

	for _, v := range earList {
		val := bytes.Split(v, []byte(","))
		var earnings = model.DfFundEarnings{
			Name:            string(val[1]),
			Code:            string(val[0]),
			Date:            string(val[3]),
			LastUpdateTime:  now.Unix(),
			UpdatedAt:       &now,
			DailyGrowthRate: common.Int642Float64(string(val[6])),
			CumulativeNav:   common.DefaultVal(string(val[5])),
			NavPerUnit:      common.Int642Float64(string(val[4])),
			Past1Month:      common.DefaultVal(string(val[8])),
			Past1Week:       common.DefaultVal(string(val[7])),
			Past1Year:       common.DefaultVal(string(val[11])),
			Past2Years:      common.DefaultVal(string(val[12])),
			Past3Months:     common.DefaultVal(string(val[9])),
			Past3Years:      common.DefaultVal(string(val[13])),
			Past6Months:     common.DefaultVal(string(val[10])),
			Past5Years:      "",
			SinceInception:  common.DefaultVal(string(val[15])),
			ThisYear:        common.DefaultVal(string(val[14])),
		}

		if earnings.Date == "" {
			earnings.Date = "0001-01-01"
		}

		//成立日
		if id, ok := codeMap[earnings.Code]; ok {
			updateBuff.WriteString(fmt.Sprintf(sql, string(val[16]), nowDate, id))
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
		common.FuncDb.CreateInBatches(&bufferEarnings, 100)
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

// 获取已购买的收益
func (f *FundEarnings) GetPurData() {
	var purchases []model.DfFuncPurchase
	var err = common.FuncDb.Model(&model.DfFuncPurchase{}).Where("purchase_amount > 0 and deleted_at is null").Find(&purchases).Error
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}
	if len(purchases) == 0 {
		return
	}
	now := time.Now()
	for _, purchase := range purchases {
		var code = purchase.Code
		refer := [][2]string{
			{"Referer", fmt.Sprintf("https://fundf10.eastmoney.com/jdzf_%s.html", code)},
		}
		res, err := common.HttpRequest(http.MethodPost, fmt.Sprintf(common.CodeEarningsUrl, code), nil, refer)
		if err != nil {
			log.Println(err.Error())
			common.Logger.Error(err.Error())
			return
		}

		var (
			ThisYear       = []byte(`今年来`)
			Past1Week      = []byte(`近1周`)
			Past1Month     = []byte(`近1月`)
			Past3Months    = []byte(`近3月`)
			Past6Months    = []byte(`近6月`)
			Past1Year      = []byte(`近1年`)
			Past2Years     = []byte(`近2年`)
			Past3Years     = []byte(`近3年`)
			Past5Years     = []byte(`近5年`)
			SinceInception = []byte(`成立来`)
			dataHead       = []byte(`</li><li class='tor red bold'>`)
			dataTail       = []byte(`</li>`)

			timerand = [][]byte{
				ThisYear,
				Past1Week,
				Past1Month,
				Past3Months,
				Past6Months,
				Past1Year,
				Past2Years,
				Past3Years,
				Past5Years,
				SinceInception,
			}
		)

		var ok bool
		var earnings = model.DfFundEarnings{
			LastUpdateTime: now.Unix(),
			UpdatedAt:      &now,
		}
		for _, times := range timerand {
			if index := bytes.Index(res, times); index != 1 {
				res = res[index:]
				tail := bytes.Index(res, dataHead)
				if tail == len(times) {
					ok = true
					res = res[tail+len(dataHead):]
					tail = bytes.Index(res, dataTail) - 1
					switch string(times) {
					case string(ThisYear):
						earnings.ThisYear = string(res[:tail])
					case string(Past1Week):
						earnings.Past1Week = string(res[:tail])
					case string(Past1Month):
						earnings.Past1Month = string(res[:tail])
					case string(Past3Months):
						earnings.Past3Months = string(res[:tail])
					case string(Past6Months):
						earnings.Past6Months = string(res[:tail])
					case string(Past1Year):
						earnings.Past1Year = string(res[:tail])
					case string(Past2Years):
						earnings.Past2Years = string(res[:tail])
					case string(Past3Years):
						earnings.Past3Years = string(res[:tail])
					case string(Past5Years):
						earnings.Past5Years = string(res[:tail])
					case string(SinceInception):
						earnings.SinceInception = string(res[:tail])
					}
				}
			}
		}
		if ok {
			common.FuncDb.Where("code = ?", purchase.Code).Updates(earnings)
		}
	}

}
