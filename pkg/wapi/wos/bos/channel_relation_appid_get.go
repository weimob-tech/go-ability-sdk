package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ChannelRelationAppidGet
// @id 2165
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2165?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询店铺下渠道应用appId)
func (client *Client) ChannelRelationAppidGet(request *ChannelRelationAppidGetRequest) (response *ChannelRelationAppidGetResponse, err error) {
	rpcResponse := CreateChannelRelationAppidGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ChannelRelationAppidGetRequest struct {
	*api.BaseRequest

	ChannelType string `json:"channelType,omitempty"`
}

type ChannelRelationAppidGetResponse struct {
	AuthorizerAppid string `json:"authorizerAppid,omitempty"`
}

func CreateChannelRelationAppidGetRequest() (request *ChannelRelationAppidGetRequest) {
	request = &ChannelRelationAppidGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "channel/relation/appid/get", "apigw")
	request.Method = api.POST
	return
}

func CreateChannelRelationAppidGetResponse() (response *api.BaseResponse[ChannelRelationAppidGetResponse]) {
	response = api.CreateResponse[ChannelRelationAppidGetResponse](&ChannelRelationAppidGetResponse{})
	return
}
