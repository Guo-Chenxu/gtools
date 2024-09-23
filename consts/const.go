package consts

import "time"

const (
	HttpHeaderUId         = "U-Id"
	HttpTraceId           = "Trace-Id"
	HttpHeaderAccessToken = "Access-Token"
	HttpHeaderAuthToken   = "Auth-Token"
	HttpHeaderBid         = "B-Id"
	HttpHeaderWebsite     = "Web-Site"
)

const (
	RedisVisitorKeyPrefix      = "GTOOLS:VISITOR:%s"
	RedisVisitorInfoExpireTime = 10 * time.Minute
)

const IsVisitorInfoSuccess = "success"

const VisitorInfoMySQLTablePrefix = "visitor_info_%s"

const TimeFormat = "2006-01-02 15:04:05"

var VisitorInfoKey = []string{"week", "ip", "location", "browser", "browser_ver", "system", "path"}

const PostFileBaseUrl = "https://www.chenxutalk.top/file/"