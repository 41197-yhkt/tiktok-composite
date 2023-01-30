package pack

import (
	"errors"
	"tiktok-composite/kitex_gen/composite"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *composite.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *composite.BaseResp {
	return &composite.BaseResp{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
