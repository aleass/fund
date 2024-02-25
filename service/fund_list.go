package service

import (
	"bytes"
	"encoding/json"
	"fund/common"
	"fund/model"
	"log"
	"net/http"
	"time"
	"unsafe"
)

var fundListFormat = []byte("var r = ")

type FundList struct {
	data [][]string
}

func (f *FundList) GetData() {
	common.Logger.Info("执行 基金列表")

	raw, err := common.HttpRequest(http.MethodGet, common.FundListUrl, nil, nil)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	index := bytes.Index(raw, fundListFormat)
	if index == -1 {
		return
	}
	raw = raw[index+len(fundListFormat) : len(raw)-1]
	json.Unmarshal(raw, &f.data)
	f.extract()
}

type memequal struct {
	Code       string `gorm:"column:code"        desc:"代码"`
	Name       string `gorm:"column:name"        desc:"名字"`
	Pinyin     string `gorm:"column:pinyin"      desc:"拼音"`
	AbbrPinyin string `gorm:"column:abbr_pinyin" desc:"拼音简写"`
	Type       string `gorm:"column:type"        desc:"基金类型"`
}

func (f *FundList) extract() {
	//检查新增的基金
	var df []model.DfFundList
	common.FuncDb.Model(&model.DfFundList{}).Find(&df)
	var codeMap = make(map[string]*model.DfFundList, len(df))
	for i, fund := range df {
		codeMap[fund.Code] = &df[i]
	}
	var newFund = make([]model.DfFundList, 0, 100)
	var updateFund = make([]model.DfFundList, 0, 100)
	now := time.Now()
	for _, fund := range f.data {
		temp := model.DfFundList{
			AbbrPinyin:     fund[1],
			Code:           fund[0],
			Name:           fund[2],
			Pinyin:         fund[4],
			Type:           fund[3],
			LastUpdateTime: now.Unix(),
			UpdatedAt:      &now,
		}
		if _fund, ok := codeMap[fund[0]]; ok {
			e1 := *(*memequal)(unsafe.Pointer(&temp))
			e2 := *(*memequal)(unsafe.Pointer(_fund))
			if e2 == e1 {
				continue
			}
			temp.Id = _fund.Id
			updateFund = append(updateFund, temp)
			continue
		}
		temp.AddTime = now.Unix()
		temp.CreatedAt = &now
		newFund = append(newFund, temp)
	}

	//新增基金
	if len(newFund) > 0 {
		db := common.FuncDb.CreateInBatches(newFund, 100)
		if err := db.Error; err != nil {
			log.Println(err.Error())
			common.Logger.Error(err.Error())
			return
		}
	}
	//更新基金
	if len(updateFund) > 0 {
		for _, fund := range updateFund {
			if err := common.FuncDb.Updates(fund).Error; err != nil {
				log.Println(err.Error())
				common.Logger.Error(err.Error())
				return
			}
		}

	}
	//清空
	f.data = f.data[:0]
}
