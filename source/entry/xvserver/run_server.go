package xvserver

import (
	"fmt"
	"net/http"

	"github.com/starryck/x-lib-go/source/utility/xbspvs"

	"github.com/starryck/x-srv-account/source/core/base/xvcfg"
	"github.com/starryck/x-srv-account/source/mode/server/xvsrv"
)

func Execute() error {
	supervisor := xbspvs.GetSupervisor(nil)
	supervisor.Handle(&ServerProcess{})
	supervisor.RunForever()
	return nil
}

type ServerProcess struct {
	xbspvs.ServerProcess
}

func (process *ServerProcess) Setup() error {
	process.SetServer(&http.Server{
		Addr:    fmt.Sprintf(":%d", xvcfg.GetServicePort()),
		Handler: xvsrv.GetRouter().GetEngine(),
	})
	return nil
}
