package consts

type BizCode struct {
	// 错误码
	Code int32
	// 错误描述
	Msg string
}

var (
	ResSuccess         = BizCode{0, "success"}
	RetParamError      = BizCode{100, "参数错误"}
	ParamBindJsonError = BizCode{101, "参数解析错误"}
	SystemErr          = BizCode{102, "服务繁忙，请稍后重试"}

	QueryRecordError = BizCode{201, "数据查询异常"}
	WriteDbError     = BizCode{202, "数据写入异常"}
	VistorInfoExistError = BizCode{203, "访客信息已存在"}

	SendEmailError = BizCode{501, "邮件发送失败"}
)
