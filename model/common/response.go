/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-21 22:52:39
 * @LastEditTime: 2024-09-21 23:10:53
 */
package common

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	Success = 0
	Error   = -1
)

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse() *Response {
	return NewResponse(Success, "成功", nil)
}

func SuccessWithMessage(message string) *Response {
	return NewResponse(Success, message, nil)
}

func SuccesswithData(data interface{}) *Response {
	return NewResponse(Success, "成功", data)
}

func SuccessWithComplete(message string, data interface{}) *Response {
	return NewResponse(Success, message, data)
}

func ErrorResponse() *Response {
	return NewResponse(Error, "失败", nil)
}

func ErrorWithMessage(message string) *Response {
	return NewResponse(Error, message, nil)
}

func ErrorWithCodeAndMessage(code int, message string) *Response {
	return NewResponse(code, message, nil)
}

func ErrorWithData(data interface{}) *Response {
	return NewResponse(Error, "失败", data)
}
func ErrorWithCodeAndData(code int, data interface{}) *Response {
	return NewResponse(code, "失败", data)
}

func ErrorWithComplete(message string, data interface{}) *Response {
	return NewResponse(Error, message, data)
}

func ErrorWithCodeAndComplete(code int, message string, data interface{}) *Response {
	return NewResponse(code, message, data)
}
