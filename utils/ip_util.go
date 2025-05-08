package utils

import (
	"context"
	"fmt"
)

type IPInfoData struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

const (
	ipUrl = "https://ipinfo.io/%s/json?token=%s"
)

func GetIPInfo(ctx context.Context, ip, tokn string) (*IPInfoData, error) {
	reqUrl := fmt.Sprintf(ipUrl, ip, tokn)
	resp := &IPInfoData{}
	err := DoGet(ctx, reqUrl, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
