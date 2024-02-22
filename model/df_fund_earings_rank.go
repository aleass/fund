package model

import "time"

type DfFundEarningsRank struct {
	Id             int64      `gorm:"column:id"                desc:""`
	Date           string     `gorm:"column:date"                desc:"日期"`
	Name           string     `gorm:"column:name"              desc:"基金简称"`
	Code           string     `gorm:"column:code"              desc:"基金代码"`
	Rank           int        `gorm:"column:rank"      desc:"排名"`
	RankPrecent    float64    `gorm:"column:rank_precent"      desc:"排名百分比"`
	TotalRate      string     `gorm:"column:total_rate"      desc:"收益率"`
	KindAvgRate    string     `gorm:"column:kind_avg_rate"      desc:"同类平均收益率"`
	DayIncreRate   float64    `gorm:"column:day_incre_rate" desc:"日增长率"`
	DayIncreVal    float64    `gorm:"column:day_incre_val"  desc:"日增长值"`
	TotalNV        float64    `gorm:"column:total_NV"       desc:"累计净值"`
	UnitNV         float64    `gorm:"column:unit_NV"        desc:"单位净值"`
	Gain           *string    `gorm:"column:gain"      desc:"创建时间"`
	AddTime        int64      `gorm:"column:add_time"         desc:"添加时间"`
	LastUpdateTime int64      `gorm:"column:last_update_time" desc:"最后更新时间"`
	CreatedAt      *time.Time `gorm:"column:created_at"       desc:"创建时间框架维护"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"       desc:"更新时间框架维护"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"       desc:"软删字段"`
}

func (DfFundEarningsRank) TableName() string {
	return "df_fund_earnings_rank"
}
