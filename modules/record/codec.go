package record

import (
	"github.com/irisnet/irismod/modules/record"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
