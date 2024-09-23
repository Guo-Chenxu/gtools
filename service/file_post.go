package service

import (
	"context"
	"fmt"
	"gtools/biz/model/gtools"
	"gtools/conf"
	"gtools/consts"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
)

func FilePost(ctx context.Context, req *gtools.FilePostReq) (string, *consts.BizCode) {
	uid := strings.ReplaceAll(uuid.New().String(), "-", "")
	ext := filepath.Ext(req.Filename)
	filePath := filepath.Join(conf.GetConfig().StorePath, fmt.Sprintf("%s%s", uid, ext))
	err := os.WriteFile(filePath, req.File, 0644)
	if err != nil {
		hlog.Errorf("FilePost err: %+v", err)
		return "", &consts.SystemErr
	}
	return fmt.Sprintf("%s%s%s", consts.PostFileBaseUrl, uid, ext), nil
}
