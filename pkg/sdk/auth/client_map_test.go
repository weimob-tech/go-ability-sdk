package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"testing"
)

func TestCreateClientInfo(t *testing.T) {
	config.SetDefault("weimob.cloud.foo.clientId", "fooId")
	config.SetDefault("weimob.cloud.foo.clientSecret", "fooSecret")
	config.SetDefault("weimob.cloud.bar.clientId", "barId")
	config.SetDefault("weimob.cloud.bar.clientSecret", "barSecret")

	Convey("Create client info", t, func() {
		expected := map[string]*auth.ClientInfo{
			"foo": {ClientId: "fooId", ClientSecret: "fooSecret"},
			"bar": {ClientId: "barId", ClientSecret: "barSecret"},
		}
		actual := CreateClientInfo()
		So(actual, ShouldResemble, expected)
	})
}
