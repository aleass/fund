package service

import (
	"fmt"
	"go.uber.org/zap/buffer"
	"log"
	"time"
	"weather/common"
)

type daysPastTimeRank struct {
}

type sqlInfo struct {
	sql  string
	name string
}

func (s *daysPastTimeRank) Send() {
	common.Logger.Info("执行 基金购买情况")
	if len(common.MyConfig.Fund.Notes) == 0 {
		return
	}
	sqlArr := []sqlInfo{
		//{common.DaysPastTimeAverSql, "今年基金大于平均值"},
		{common.DaysPastTimeRankSql, "表现优秀的基金"},
	}
	var list []common.DaysPastTimeRank
	for _, sql := range sqlArr {
		if db := common.FuncDb.Raw(sql.sql).Find(&list); db.Error != nil {
			common.Logger.Error(db.Error.Error())
		}
		var msg = buffer.Buffer{}
		var _, err = msg.WriteString(time.Now().Format(common.UsualDate) + "   " + sql.name + "\n")
		if err != nil {
			log.Println(err.Error())
			common.Logger.Error(err.Error())
		}
		for _, info := range list {
			msg.WriteString(fmt.Sprintf("%s %s\n", info.Code, info.Name))
		}
		msg.WriteString(fmt.Sprintf("\n\n查看数据点击 http://%s:8080/fund/day", common.MyConfig.Fund.Host))

		//for _, note := range config.MyConfig.Fund.Notes {
		//	//common.Send(msg.String(), GetWechatUrl(note))
		//}
	}
}
