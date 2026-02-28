package xerr

import (
	"errors"
	"fmt"
	"runtime"
)

type XErr struct {
	code        Code
	inputDetail string
	stack       []uintptr
}

func New(code Code) *XErr {
	e := &XErr{
		code: code,
	}
	e.setCallerStack()
	return e
}

func NewWithDetail(code Code, detail string) *XErr {
	e := &XErr{
		code:        code,
		inputDetail: detail,
	}
	e.setCallerStack()
	return e
}

func NewWithError(code Code, err error, detail string) *XErr {
	e := &XErr{
		code:        code,
		inputDetail: fmt.Sprintf("%s: %v", detail, err),
	}
	e.setCallerStack()
	return e
}

func (e *XErr) Error() string {
	detail := e.inputDetail
	if detail == "" {
		detail = e.code.Detail()
	}
	return fmt.Sprintf("[%d] - %s", e.code, detail)
}

func (e *XErr) setCallerStack() {
	if e.code == CodeSuccess {
		return
	}
	const skip = 3  // 跳过3层栈帧
	const depth = 5 // 栈跟踪的最大深度
	stack := make([]uintptr, depth)
	n := runtime.Callers(skip, stack)
	e.stack = stack[:n]
}

func (e *XErr) FormatStack() string {
	if e.stack == nil {
		return ""
	}

	frames := runtime.CallersFrames(e.stack)
	var stackInfo string

	for {
		frame, more := frames.Next()
		stackInfo += fmt.Sprintf("\t%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}

	return stackInfo
}

func (e *XErr) Code() Code {
	return e.code
}

func (e *XErr) Detail() string {
	if e.inputDetail == "" {
		return e.code.Detail()
	}
	return e.inputDetail
}

type Code int

const (
	// 常规错误码
	CodeSuccess            Code = 200
	CodeInternalError      Code = 500
	CodeBadRequest         Code = 400
	CodeNotFound           Code = 404
	CodeUnauthorized       Code = 401
	CodeForbidden          Code = 403
	CodeTimeout            Code = 408
	CodeConflict           Code = 409
	CodeTooManyRequests    Code = 429
	CodeServiceUnavailable Code = 503

	// 外部组件错误码
	CodeMySQLError   Code = 1000
	CodeRedisError   Code = 1001
	CodeMongoDBError Code = 1002
	CodeKafkaError   Code = 1003
	CodeParamError   Code = 1004
)

var codeDetails = map[Code]string{
	CodeSuccess:            "CodeSuccess",
	CodeInternalError:      "CodeInternalError",
	CodeBadRequest:         "CodeBadRequest",
	CodeNotFound:           "CodeNotFound",
	CodeUnauthorized:       "CodeUnauthorized",
	CodeForbidden:          "CodeForbidden",
	CodeTimeout:            "CodeTimeout",
	CodeConflict:           "CodeConflict",
	CodeTooManyRequests:    "CodeTooManyRequests",
	CodeServiceUnavailable: "CodeServiceUnavailable",

	CodeMySQLError:   "CodeMySQLError",
	CodeRedisError:   "CodeRedisError",
	CodeMongoDBError: "CodeMongoDBError",
	CodeKafkaError:   "CodeKafkaError",
	CodeParamError:   "CodeParamError",
}

func RegisterCodeDetail(dmap map[Code]string) {
	for k, v := range dmap {
		codeDetails[k] = v
	}
}

func (c Code) Detail() string {
	if detail, exists := codeDetails[c]; exists {
		return detail
	}
	return "undefined error code"
}

// ToXErr 进程内传递的err，转化为xerr，返回值永远非空
func ToXErr(err error) *XErr {
	if err == nil {
		return New(CodeSuccess)
	}
	var e *XErr
	if errors.As(err, &e) {
		return e
	}
	return NewWithDetail(CodeInternalError, err.Error())
}
