package utils

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func DoPost(ctx context.Context, url string, data interface{}, res interface{}) error {
	start := time.Now()
	defer func() {
		hlog.CtxInfof(ctx, "out request, method:%s, cost:%v", url, TimeSub(start))
	}()
	// 建立链接
	client := &http.Client{}
	reqData := JSONMarshal(data)
	hlog.CtxInfof(ctx, "out request, method:%s, req:%s, ", url, reqData)
	req, _ := http.NewRequest("POST", url, strings.NewReader(reqData))
	req.Header.Set("Content-Type", "application/json")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Post] post fail, err:%v, data:%v", err, reqData)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		hlog.CtxErrorf(ctx, "[Http Post] status code not 200, url:%s  code:%d, data:%v", url, resp.StatusCode, JSONMarshal(reqData))
		return errors.New("status code not 200")
	}
	// 解析校验结果
	body, _ := io.ReadAll(resp.Body)
	err = sonic.Unmarshal(body, &res)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Post] json unmarshal fail, err:%v", err)
		return err
	}
	return nil
}

func PostByUrlencoded(ctx context.Context, url string, reqData url.Values, res interface{}) error {
	start := time.Now()
	defer func() {
		hlog.CtxInfof(ctx, "out request, method:%s, cost:%v", url, TimeSub(start))
	}()
	hlog.CtxInfof(ctx, "out request, method:%s, req:%s, ", url, JSONMarshal(reqData))

	// 建立链接
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(reqData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Post] post fail, err:%v, data:%v", err, reqData)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		hlog.CtxErrorf(ctx, "[Http Post] status code not 200,  code:%d, data:%v", resp.StatusCode, JSONMarshal(reqData))
		return errors.New("status code not 200")
	}
	// 解析校验结果
	body, _ := io.ReadAll(resp.Body)
	err = sonic.Unmarshal(body, &res)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Post] json unmarshal fail, err:%v", err)
		return err
	}
	return nil
}

func DoGet(ctx context.Context, baseUrl string, params url.Values, res any) error {
	start := time.Now()
	uri := baseUrl
	if params != nil {
		uri = fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	}
	defer func() {
		hlog.CtxInfof(ctx, "out request, method:%s, cost:%v", baseUrl, TimeSub(start))
	}()
	// 构建请求对象
	client := &http.Client{}
	client.Timeout = time.Second * 10
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Get] new request fail, err:%v, params:%v", err, params)
		return err
	}
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Get] get fail, err:%v, params:%v", err, params)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		hlog.CtxErrorf(ctx, "[Http Get] status code not 200,  code:%d, params:%v", resp.StatusCode, params)
		return errors.New("status code not 200")
	}
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Get] read body fail, err:%v, params:%v", err, params)
		return err
	}
	// 解析响应
	err = sonic.Unmarshal(body, &res)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Http Get] json unmarshal fail, err:%v, params:%v", err, params)
		return err
	}
	return nil
}
