package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserChannelUnbind
// @id 2046
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2046?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=释放用户渠道)
func (client *Client) UserChannelUnbind(request *UserChannelUnbindRequest) (response *UserChannelUnbindResponse, err error) {
	rpcResponse := CreateUserChannelUnbindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserChannelUnbindRequest struct {
	*api.BaseRequest

	BosId      int64  `json:"bosId,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
	Source     int64  `json:"source,omitempty"`
	AppId      string `json:"appId,omitempty"`
	OriginalId string `json:"originalId,omitempty"`
}

type UserChannelUnbindResponse struct {
}

func CreateUserChannelUnbindRequest() (request *UserChannelUnbindRequest) {
	request = &UserChannelUnbindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/channel/unbind", "apigw")
	request.Method = api.POST
	return
}

func CreateUserChannelUnbindResponse() (response *api.BaseResponse[UserChannelUnbindResponse]) {
	response = api.CreateResponse[UserChannelUnbindResponse](&UserChannelUnbindResponse{})
	return
}
