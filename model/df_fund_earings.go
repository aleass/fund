package model

import "time"

type DfFundEarnings struct {
	Id              int64      `gorm:"column:id"                desc:""`
	Code            string     `gorm:"column:code"              desc:"基金代码"`
	Date            string     `gorm:"column:date"              desc:"数据的最新日期"`
	CumulativeNav   string     `gorm:"column:cumulative_nav"    desc:"累计净值"`
	DailyGrowthRate string     `gorm:"column:daily_growth_rate" desc:"日增长百分率"`
	Name            string     `gorm:"column:name"              desc:"基金简称"`
	NavPerUnit      float64    `gorm:"column:nav_per_unit"      desc:"单位净值"`
	Past1Month      string     `gorm:"column:past_1_month"      desc:"近1个月增长率"`
	Past1Week       string     `gorm:"column:past_1_week"       desc:"近1周增长率"`
	Past1Year       string     `gorm:"column:past_1_year"       desc:"近1年增长率"`
	Past2Years      string     `gorm:"column:past_2_years"      desc:"近2年增长率"`
	Past3Months     string     `gorm:"column:past_3_months"     desc:"近3个月增长率"`
	Past3Years      string     `gorm:"column:past_3_years"      desc:"近3年增长率"`
	Past5Years      string     `gorm:"column:past_5_years"      desc:"近5年增长率"`
	Past6Months     string     `gorm:"column:past_6_months"     desc:"近6个月增长率"`
	SinceInception  string     `gorm:"column:since_inception"   desc:"成立来增长率"`
	ThisYear        string     `gorm:"column:this_year"         desc:"今年来增长率"`
	AddTime         int64      `gorm:"column:add_time"         desc:"添加时间"`
	LastUpdateTime  int64      `gorm:"column:last_update_time" desc:"最后更新时间"`
	CreatedAt       *time.Time `gorm:"column:created_at"       desc:"创建时间框架维护"`
	UpdatedAt       *time.Time `gorm:"column:updated_at"       desc:"更新时间框架维护"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"       desc:"软删字段"`
}

func (DfFundEarnings) TableName() string {
	return "df_fund_earnings"
}
