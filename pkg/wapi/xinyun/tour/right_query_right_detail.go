package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightQueryRightDetail
// @id 1007
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线路维权详情)
func (client *Client) RightQueryRightDetail(request *RightQueryRightDetailRequest) (response *RightQueryRightDetailResponse, err error) {
	rpcResponse := CreateRightQueryRightDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightQueryRightDetailRequest struct {
	*api.BaseRequest

	RightNo string `json:"rightNo,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
}

type RightQueryRightDetailResponse struct {
}

func CreateRightQueryRightDetailRequest() (request *RightQueryRightDetailRequest) {
	request = &RightQueryRightDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "right/queryRightDetail", "api")
	request.Method = api.POST
	return
}

func CreateRightQueryRightDetailResponse() (response *api.BaseResponse[RightQueryRightDetailResponse]) {
	response = api.CreateResponse[RightQueryRightDetailResponse](&RightQueryRightDetailResponse{})
	return
}
