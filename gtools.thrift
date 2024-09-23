namespace go gtools

struct BaseResp {
    1: i32 code
    2: string msg
}

struct SendEmailReq {
    1: string from
    2: list<string> to
    3: string secret
    4: string host
    5: i32 port
    6: string nickname
    7: string subject
    8: string body
    9: bool ssl
}

struct SendEmailResp {
    1: i32 code
    2: string msg
    3: bool data
}

struct AddVisitorInfoReq {
	1: bool success
	2: string time
	3: string week
	4: string ip
	5: string location
	6: string browser
	7: string browser_ver
	8: string system
	9: string low
	10: string high
	11: string tq
	12: string fl
	13: string tip
	14: string fengxiang
    15: string path
    16: string domain
}

struct AddVisitorInfoResp {
    1: i32 code
    2: string msg
    3: bool data
}

struct CountVisitorReq{
    1: string domain
    2: string path
}

struct CountVisitorResp {
    1: i32 code
    2: string msg
    3: i64 data
}

struct FilePostReq{
    1: string filename
    2: binary file (api.form="file")
}

struct FilePostResp{
    1: i32 code
    2: string msg
    3: string data
}

service ToolsHandler {
    SendEmailResp SendEmail(1: SendEmailReq req) (api.post="/api/tools/send_email")
    AddVisitorInfoResp AddVisitorInfo(1: AddVisitorInfoReq req) (api.post="/api/tools/add_visitor_info")
    CountVisitorResp CountVisitorByPath(1: CountVisitorReq req) (api.get="/api/tools/count_visitor_by_path")
    FilePostResp FilePost(1: FilePostReq req) (api.post="/api/tools/file_post")
}