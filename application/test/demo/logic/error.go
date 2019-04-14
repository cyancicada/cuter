package logic

import "github.com/yakaa/cuter/common/baseerror"

var (
	ErrInvalidParam = baseerror.NewCodeError(10001, "传入参数有误")

	ErrPermissionDenied     = baseerror.NewCodeError(10001, "权限不足")
)
