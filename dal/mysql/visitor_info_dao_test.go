package mysql

import (
	"context"
	"gtools/conf"
	"gtools/consts"
	"gtools/utils"
	"testing"
)

func TestInsertVisitorInfo(t *testing.T) {
	ctx := context.Background()
	conf.TestInit()
	Init(ctx, conf.GetMysql())

	visitorInfo := &VisitorInfo{
		Success:    false,
		Time:       utils.ParseTime("2024-08-17 17:46:18", consts.TimeFormat),
		Week:       "星期六",
		IP:         "191.123.211.168",
		Location:   "北京-北京市",
		Browser:    "Edge",
		BrowserVer: "127.0.0.0",
		System:     "Windows 10",
		Low:        "23°C",
		High:       "30°C",
		TQ:         "小雨",
		FL:         "1-3级",
		Tip:        "天太热了，吃个西瓜~",
		FengXiang:  "南风",
		Path:       "https://aa.com/index/页面",
		Domain:     "aaa.com",
	}

	err := NewVisitInfoDao().InsertVisitorInfo(ctx, "test", visitorInfo)
	if err != nil {
		t.Error(err)
	}
}
