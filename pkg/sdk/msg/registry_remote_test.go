package msg

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"testing"
)

func TestRegisterMsgServiceRemote(t *testing.T) {
	config.SetDefault("msg.sdk.version", SdkVersionV2)
	config.SetDefault("msg.register.url", "http://yym-open.weimobqa.com/open/extend/core/server/v2/registerExtendPointSecret")
	config.SetDefault("msg.register.version", SdkVersionV2)
	config.SetDefault("weimob_cloud_appId", "480209412352_go-demo")
	config.SetDefault("weimob_cloud_env", "1")

	Convey("Regsiter remote spi service", t, func(c C) {
		clientInfo := auth.ClientInfo{
			ClientId:     "477FD526146E55C173E128C3596DC790",
			ClientSecret: "D3C10901C731115D8DF6DA38D19CFDD1",
		}
		registerInfo := RegisterInfo{
			RegisterMsgInfo: RegisterMsgInfo{
				Path:      "weimob/cloud/wos/message/receive",
				SpecsType: SpecTypeWos,
			},
			InterfacePathVos: []any{},
		}

		err := RegistryRemoteForClient(registerInfo, clientInfo)
		So(err, ShouldBeNil)
	})
}
