package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserChannelBind
// @id 2038
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2038?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=绑定渠道和wid)
func (client *Client) UserChannelBind(request *UserChannelBindRequest) (response *UserChannelBindResponse, err error) {
	rpcResponse := CreateUserChannelBindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserChannelBindRequest struct {
	*api.BaseRequest

	BosId      int64  `json:"bosId,omitempty"`
	Source     int64  `json:"source,omitempty"`
	AppId      string `json:"appId,omitempty"`
	OriginalId string `json:"originalId,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}

type UserChannelBindResponse struct {
}

func CreateUserChannelBindRequest() (request *UserChannelBindRequest) {
	request = &UserChannelBindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/channel/bind", "apigw")
	request.Method = api.POST
	return
}

func CreateUserChannelBindResponse() (response *api.BaseResponse[UserChannelBindResponse]) {
	response = api.CreateResponse[UserChannelBindResponse](&UserChannelBindResponse{})
	return
}
