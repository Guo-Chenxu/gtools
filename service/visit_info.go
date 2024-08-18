package service

import (
	"context"
	"fmt"
	"gtools/biz/model/gtools"
	"gtools/consts"
	"gtools/dal/mysql"
	"gtools/utils"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func AddVisitorInfo(ctx context.Context, req *gtools.AddVisitorInfoReq) *consts.BizCode {
	hlog.CtxInfof(ctx, "信息入库")
	tableName := fmt.Sprintf(consts.VisitorInfoMySQLTablePrefix, strings.Replace(req.Domain, ".", "_", -1))
	return mysql.NewVisitInfoDao().InsertVisitorInfo(ctx, tableName, convertVisitorInfo(req))
}

func CountVisitorByPath(ctx context.Context, req *gtools.CountVisitorReq) (int64, *consts.BizCode) {
	tableName := fmt.Sprintf(consts.VisitorInfoMySQLTablePrefix, strings.Replace(req.Domain, ".", "_", -1))   
	return mysql.NewVisitInfoDao().CountVisitorByPath(ctx, tableName, req.Path)
}

func convertVisitorInfo(req *gtools.AddVisitorInfoReq) *mysql.VisitorInfo {
	return &mysql.VisitorInfo{
		Success:    req.Success,
		Time:       utils.ParseTime(req.Time, consts.TimeFormat),
		Week:       req.Week,
		IP:         req.IP,
		Location:   req.Location,
		Browser:    req.Browser,
		BrowserVer: req.BrowserVer,
		System:     req.System,
		Low:        req.Low,
		High:       req.High,
		TQ:         req.Tq,
		FL:         req.Fl,
		Tip:        req.Tip,
		FengXiang:  req.Fengxiang,
		Path:       req.Path,
		Domain:     req.Domain,
	}
}
