package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightAllowRight
// @id 1008
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=商家同意维权)
func (client *Client) RightAllowRight(request *RightAllowRightRequest) (response *RightAllowRightResponse, err error) {
	rpcResponse := CreateRightAllowRightResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightAllowRightRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
	RightNo string `json:"rightNo,omitempty"`
}

type RightAllowRightResponse struct {
}

func CreateRightAllowRightRequest() (request *RightAllowRightRequest) {
	request = &RightAllowRightRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "right/allowRight", "api")
	request.Method = api.POST
	return
}

func CreateRightAllowRightResponse() (response *api.BaseResponse[RightAllowRightResponse]) {
	response = api.CreateResponse[RightAllowRightResponse](&RightAllowRightResponse{})
	return
}
