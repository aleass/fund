package common

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	vip      = viper.New()
	path     = "pkg/config.yaml"
	MyConfig = Config{}
)

// 运行
func InitConfig() {
	// 使用 os.Stat 函数获取文件的信息
	_, err := os.Stat(path)
	// 检查文件是否存在
	if os.IsNotExist(err) {
		path = "config.yaml"
	}
	vip.SetConfigFile(path)
	vip.SetConfigType("yaml")
	if err = vip.ReadInConfig(); err != nil {
		panic(fmt.Errorf("无法读取配置文件: %w", err))
	}

	if err = vip.Unmarshal(&MyConfig); err != nil {
		panic(fmt.Errorf("无法解析配置文件: %w", err))
	}
	//mysql
	InitMysql()
	//wechat
	initWeChat()
}

func initWeChat() {
	for _, v := range MyConfig.Wechat {
		wechatNoteMap[v.Notes] = v.Token
	}
}

type Config struct {
	// 结构映射
	Wechat []struct {
		Token string `mapstructure:"token"`
		Notes string `mapstructure:"note"`
	} `mapstructure:"wechat"`
	CaiYun struct {
		Token  string `json:"token"`
		Addres []struct {
			Addr        string `json:"addr"`
			WechatNotes string `json:"wechatNotes"`
			Coordinate  string `json:"coordinate"`
			Switch      bool   `json:"switch" desc:"开关"`
			AllowWeek   string `json:"allowWeek"`
		} `json:"addres"`
	} `json:"caiyun"`

	UrlConfigPass []struct {
		Name  string `json:"name"`
		Notes string `mapstructure:"note"`
	} `json:"urlConfigPass"`

	GeoMapToken string `json:"geoMapToken"`
	QqMapToken  string `json:"qqMapToken"`

	DB struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		DbName   string `json:"dbName"`
		Port     string `json:"port"`
	} `json:"db"`

	Fund struct {
		Host  string   `json:"host"`
		Notes []string `mapstructure:"notes"`
	} `json:"fund"`
}
