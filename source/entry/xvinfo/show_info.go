package xvinfo

import (
	"github.com/starryck/x-lib-go/source/core/utility/xblogger"

	"github.com/starryck/x-srv-account/source/core/base/xvcfg"
)

func Execute() error {
	xblogger.WithFields(xblogger.Fields{
		"config": xvcfg.GetConfig(),
	}).Info("Log info message.")
	return nil
}
