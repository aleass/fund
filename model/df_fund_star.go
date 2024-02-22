package model

import "time"

type DfFundStar struct {
	Id                         int64      `gorm:"column:id"                            desc:""`
	Code                       string     `gorm:"column:code"                          desc:""`
	JiananJinxinSecuritiesDate string     `gorm:"column:Jianan_Jinxin_Securities_date" desc:"济安金信-更新时间"`
	JiananJinxinStar           string     `gorm:"column:Jianan_Jinxin_star"            desc:"济安金信-星"`
	JiananJinxinTrend          string     `gorm:"column:Jianan_Jinxin_trend"           desc:"济安金信-趋势 up down"`
	Name                       string     `gorm:"column:name"                          desc:""`
	ShanghaiSecuritiesDate     string     `gorm:"column:Shanghai_Securities_date"      desc:"上海证券-更新时间"`
	ShanghaiSecuritiesStar     string     `gorm:"column:Shanghai_Securities_star"      desc:"上海证券-星"`
	ShanghaiSecuritiesTrend    string     `gorm:"column:Shanghai_Securities_trend"     desc:"上海证券-趋势 up down"`
	ZhaoShangSecuritiesDate    string     `gorm:"column:ZhaoShang_Securities_date"     desc:"招商证券-更新时间"`
	ZhaoShangSecuritiesStar    string     `gorm:"column:ZhaoShang_Securities_star"     desc:"招商证券-星"`
	ZhaoShangSecuritiesTrend   string     `gorm:"column:ZhaoShang_Securities_trend"    desc:"招商证券-趋势 up down"`
	AddTime                    int64      `gorm:"column:add_time"         desc:"添加时间"`
	LastUpdateTime             int64      `gorm:"column:last_update_time" desc:"最后更新时间"`
	CreatedAt                  *time.Time `gorm:"column:created_at"       desc:"创建时间框架维护"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at"       desc:"更新时间框架维护"`
	DeletedAt                  *time.Time `gorm:"column:deleted_at"       desc:"软删字段"`
}

func (DfFundStar) TableName() string {
	return "df_fund_star"
}
