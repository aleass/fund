package service

import (
	"bytes"
	"fmt"
	"fund/common"
	"fund/model"
	"log"
	"net/http"
	"time"
)

// 获取已购买的收益
func (f *FundEarnings) GetPurData() {
	var purchases []model.DfFuncPurchase
	var err = common.FuncDb.Model(&model.DfFuncPurchase{}).Where("holding_quantity > 0 and deleted_at is null").Find(&purchases).Error
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
