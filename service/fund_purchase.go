package service

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"
	"weather/common"
	"weather/model"
)

type FundPurchase struct {
}

func (f *FundPurchase) GetData() {
	common.Logger.Info("执行 购买基金收益查询")
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
	//构建数据
	var (
		codes        []string
		pruchaseData = map[string]*model.DfFuncPurchase{}
	)
	for i, purchase := range purchases {
		codes = append(codes, purchase.Code)
		pruchaseData[purchase.Code] = &purchases[i]
	}

	//查询当前
	var fundList []model.DfFundEarnings
	err = common.FuncDb.Model(&model.DfFundEarnings{}).Where("code in (?)", codes).Where("deleted_at is null").Find(&fundList).Error
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return
	}
	currdate := time.Now().Format("2006-01-01")
	var msg strings.Builder
	for _, fund := range fundList {
		purchase, ok := pruchaseData[fund.Code]
		if !ok {
			continue
		}
		distributionDate := check(fund.Code, currdate)
		earing := purchase.Distribution + fund.NavPerUnit*purchase.HoldingQuantity - purchase.PurchaseAmount
		msg.WriteString(fmt.Sprintf("\n\n%s \n\t价值:%.2f \n\t盈利:%.2f%s",
			fund.Name, fund.NavPerUnit*purchase.HoldingQuantity, earing, distributionDate))
	}

	common.Send(msg.String()[2:], "mine")
}

// 检查是否有分红
func check(code string, currdate string) string {
	var format = []byte(`<table class='w782 comm cfxq'><thead><tr><th class='first'>年份</th><th>权益登记日</th><th>除息日</th><th>每份分红</th><th class='last'>分红发放日</th></tr></thead><tbody><tr><td>`)
	var format2 = []byte("</td><td>")
	url := `https://fundf10.eastmoney.com/fhsp_%s.html`
	url = fmt.Sprintf(url, code)

	res, err := common.HttpRequest(common.GetType, url, nil, nil)
	if err != nil {
		log.Println(err.Error())
		common.Logger.Error(err.Error())
		return ""
	}
	index := bytes.Index(res, format)
	if index == -1 {
		return ""
	}
	res = res[index+len(format):]
	index = bytes.Index(res, format2)
	if index == -1 {
		return ""
	}
	res = res[index+len(format2):]
	index = bytes.Index(res, format2)
	if index == -1 {
		return ""
	}
	res = res[:index]
	if len(res) == 0 {
		return ""
	}
	date := string(res)
	if date < currdate {
		return ""
	}

	return "\n\t分红:" + date
}
