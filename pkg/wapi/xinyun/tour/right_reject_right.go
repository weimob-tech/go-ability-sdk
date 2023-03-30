package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightRejectRight
// @id 1009
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=商家拒绝维权)
func (client *Client) RightRejectRight(request *RightRejectRightRequest) (response *RightRejectRightResponse, err error) {
	rpcResponse := CreateRightRejectRightResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightRejectRightRequest struct {
	*api.BaseRequest

	OrderNo      string `json:"orderNo,omitempty"`
	RightNo      string `json:"rightNo,omitempty"`
	RefuseReason string `json:"refuseReason,omitempty"`
}

type RightRejectRightResponse struct {
}

func CreateRightRejectRightRequest() (request *RightRejectRightRequest) {
	request = &RightRejectRightRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "right/rejectRight", "api")
	request.Method = api.POST
	return
}

func CreateRightRejectRightResponse() (response *api.BaseResponse[RightRejectRightResponse]) {
	response = api.CreateResponse[RightRejectRightResponse](&RightRejectRightResponse{})
	return
}
