package spi

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"testing"
)

func TestRegisterSpiServiceRemote(t *testing.T) {
	config.SetDefault("spi.register.url", "http://yym-open.weimobqa.com/open/extend/core/server/v2/registerExtendPointSecret")
	config.SetDefault("spi.register.version", SdkVersionV2)
	config.SetDefault("weimob_cloud_appId", "480209412352_go-demo")
	config.SetDefault("weimob_cloud_env", "1")

	Convey("Regsiter remote spi service", t, func(c C) {
		clientInfo := auth.ClientInfo{
			ClientId:     "4C2AC87EE1CE4C25E4CC16C5D8B33D55",
			ClientSecret: "C93EA7B6CA4AFD48D10C94ADFDEA8AD9",
		}
		registerInfo := RegisterInfo{
			SpiIfaceInfo: []RegisterSpiInfo{
				{
					InterfaceName: "PaasConnectorVerification1GoodsGetTwoService",
					ImplFullName:  "implPaasConnectorVerification1GoodsGetOneService",
					BeanName:      "beanPaasConnectorVerification1GoodsGetOneService",
					MethodName:    "excute",
					SpiBelongType: SpecTypeWos,
				},
			},
		}

		err := RegistryRemoteForClient(registerInfo, clientInfo)
		So(err, ShouldBeNil)
	})
}
