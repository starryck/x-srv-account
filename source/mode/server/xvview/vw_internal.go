package xvview

import (
	"github.com/starryck/x-srv-account/source/core/base/xvcfg"
	"github.com/starryck/x-srv-account/source/core/base/xvmtmsg"
)

var ReadInternalInfoHandlers = []Handler{
	InternalRequestHandler,
	ReadInternalInfoHandler,
}

func ReadInternalInfoHandler(ctx *Context) {
	flow := &ReadInternalInfoHandlerFlow{}
	flow.Initiate(ctx)
	flow.SetResult()
}

type ReadInternalInfoHandlerFlow struct {
	APISIXFlow
}

func (flow *ReadInternalInfoHandlerFlow) SetResult() {
	if flow.HasError() {
		return
	}

	info := xvcfg.GetConfig()
	flow.RespondJSON(xvmtmsg.IAV200, info, nil)
	return
}
