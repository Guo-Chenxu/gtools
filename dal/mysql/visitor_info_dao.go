package mysql

import (
	"context"
	"gtools/consts"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VisitorInfo struct {
	ID         uint      `gorm:"primary_key;auto_increment" json:"-"`
	Success    bool      `gorm:"column:success" json:"success"`
	Time       time.Time `gorm:"column:time" json:"time"`
	Week       string    `gorm:"column:week" json:"week"`
	IP         string    `gorm:"column:ip" json:"ip"`
	Location   string    `gorm:"column:location" json:"location"`
	Browser    string    `gorm:"column:browser" json:"browser"`
	BrowserVer string    `gorm:"column:browser_ver" json:"browser_ver"`
	System     string    `gorm:"column:system" json:"system"`
	Low        string    `gorm:"column:low" json:"low"`
	High       string    `gorm:"column:high" json:"high"`
	TQ         string    `gorm:"column:tq" json:"tq"`
	FL         string    `gorm:"column:fl" json:"fl"`
	Tip        string    `gorm:"column:tip" json:"tip"`
	FengXiang  string    `gorm:"column:fengxiang" json:"fengxiang"`
	Path       string    `gorm:"column:path" json:"path"`
	Domain     string    `gorm:"column:domain" json:"domain"`
}

type VisitorInfoDao struct {
}

var visitorInfoDao *VisitorInfoDao

var visitorInfoDaoOnce sync.Once

func NewVisitInfoDao() *VisitorInfoDao {
	visitorInfoDaoOnce.Do(func() {
		visitorInfoDao = &VisitorInfoDao{}
	})
	return visitorInfoDao
}

func (visitInfoDao *VisitorInfoDao) InsertVisitorInfo(ctx context.Context, tableName string, visitorInfo *VisitorInfo) *consts.BizCode {
	res := mysqlClient.Debug().Table(tableName).Create(visitorInfo)
	if res.Error != nil {
		hlog.CtxErrorf(ctx, "write db err, table: %s, %v", tableName, res.Error)
		return &consts.BizCode{Code: consts.WriteDbError.Code, Msg: consts.WriteDbError.Msg}
	}
	return nil
}
