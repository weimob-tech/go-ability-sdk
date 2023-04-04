package generic

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/hook"
)

func init() {
	config.SetDefault("client.log.lvl", "body")
	config.SetDefault("client.schema", "http")
	config.SetDefault("client.domain", "dopen.weimob.com")
	config.SetDefault("client.oauth.domain", "dopen.weimob.com")
	config.SetDefault("weimob.cloud.foo.clientId", "F580B0441D733E0A8B34367D168772DD")
	config.SetDefault("weimob.cloud.foo.clientSecret", "E51E9913EAC1C0461354B7B74184189A")

	hook.ApplyPostStartHook()
}

func TestGenericClientDoRequest(t *testing.T) {
	Convey("create new generic client", t, func() {
		clientInfo := auth.GetClientInfo("foo")
		cli, err := NewClientWithClientKey(clientInfo.ClientId, clientInfo.ClientSecret)
		So(err, ShouldBeNil)

		Convey("post request should success", func() {
			var res = api.CreateResponse(&map[string]string{})
			var payload = `{"source":"1234"}`
			var req = NewPostRequest("/apigw/bos/v2.0/security/encrypt")
			err = req.SetBody([]byte(payload))
			So(err, ShouldBeNil)
			err = cli.DoRequest(req, res)
			So(err, ShouldBeNil)
			So(res.IsSuccess(), ShouldBeTrue)

			result, err := res.GetData()
			So(err, ShouldBeNil)
			So((*result)["result"], ShouldNotBeEmpty)
			_, _ = Println(*result)
		})
	})
}
