package core

import "encoding/json"

var (
	// OK
	OK  = response(200, "ok") // 通用成功
	Err = response(500, "")   // 通用错误

	// 服务级错误码
	ErrParam     = response(10001, "参数有误")
	ErrSignParam = response(10002, "签名参数有误")

	// 模块级错误码 - 用户模块
	ErrUserService = response(20100, "用户服务异常")
	ErrUserPhone   = response(20101, "用户手机号不合法")
	ErrUserCaptcha = response(20102, "用户验证码有误")
)

type Response struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
