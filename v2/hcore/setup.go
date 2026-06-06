package hcore

import (
	"context"
	"os"

	"github.com/hiddify/hiddify-core/v2/hcommon"
	"github.com/hiddify/hiddify-core/v2/service_manager"
)

func init() {
	os.WriteFile(os.TempDir()+"/hcore_init.txt", []byte("init_ok\n"), 0644)
}

var (
	sWorkingPath          string
	sTempPath             string
	sUserID               int
	sGroupID              int
	statusPropagationPort int64
)

func InitHiddifyService() error {
	return service_manager.StartServices()
}

func (s *CoreService) Setup(ctx context.Context, req *SetupRequest) (*hcommon.Response, error) {
	mu.Lock()
	existing := grpcServer[req.Mode]
	mu.Unlock()
	if existing != nil {
		return &hcommon.Response{Code: hcommon.ResponseCode_OK, Message: ""}, nil
	}
	err := Setup(req, nil)
	code := hcommon.ResponseCode_OK
	if err != nil {
		code = hcommon.ResponseCode_FAILED
	}
	return &hcommon.Response{Code: code, Message: err.Error()}, err
}
