package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardGetStatus
// @id 131
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠券状态)
func (client *Client) CardGetStatus(request *CardGetStatusRequest) (response *CardGetStatusResponse, err error) {
	rpcResponse := CreateCardGetStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardGetStatusRequest struct {
	*api.BaseRequest

	CardCode string `json:"cardCode,omitempty"`
}

type CardGetStatusResponse struct {
}

func CreateCardGetStatusRequest() (request *CardGetStatusRequest) {
	request = &CardGetStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "card/getStatus", "api")
	request.Method = api.POST
	return
}

func CreateCardGetStatusResponse() (response *api.BaseResponse[CardGetStatusResponse]) {
	response = api.CreateResponse[CardGetStatusResponse](&CardGetStatusResponse{})
	return
}
