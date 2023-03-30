package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardQueryCard
// @id 238
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询卡券详情)
func (client *Client) CardQueryCard(request *CardQueryCardRequest) (response *CardQueryCardResponse, err error) {
	rpcResponse := CreateCardQueryCardResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardQueryCardRequest struct {
	*api.BaseRequest

	CardCode string `json:"cardCode,omitempty"`
}

type CardQueryCardResponse struct {
}

func CreateCardQueryCardRequest() (request *CardQueryCardRequest) {
	request = &CardQueryCardRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "card/queryCard", "api")
	request.Method = api.POST
	return
}

func CreateCardQueryCardResponse() (response *api.BaseResponse[CardQueryCardResponse]) {
	response = api.CreateResponse[CardQueryCardResponse](&CardQueryCardResponse{})
	return
}
