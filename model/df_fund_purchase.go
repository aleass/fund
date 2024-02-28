package model

import "time"

type DfFuncPurchase struct {
	Id              int64      `gorm:"column:id"               desc:""`
	Code            string     `gorm:"column:code"             desc:"代码"`
	Name            string     `gorm:"column:name"             desc:"名字"`
	HoldingQuantity float64    `gorm:"holding_quantity"        desc:"持有份额"`
	AddTime         int64      `gorm:"column:add_time"         desc:"添加时间"`
	LastUpdateTime  int64      `gorm:"column:last_update_time" desc:"最后更新时间"`
	CreatedAt       *time.Time `gorm:"column:created_at"       desc:"创建时间框架维护"`
	UpdatedAt       *time.Time `gorm:"column:updated_at"       desc:"更新时间框架维护"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"       desc:"软删字段"`
}

func (DfFuncPurchase) TableName() string {
	return "df_func_purchase"
}
