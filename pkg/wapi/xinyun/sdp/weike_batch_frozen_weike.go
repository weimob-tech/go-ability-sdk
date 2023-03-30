package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeBatchFrozenWeike
// @id 1882
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量冻结/解冻微客)
func (client *Client) WeikeBatchFrozenWeike(request *WeikeBatchFrozenWeikeRequest) (response *WeikeBatchFrozenWeikeResponse, err error) {
	rpcResponse := CreateWeikeBatchFrozenWeikeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeBatchFrozenWeikeRequest struct {
	*api.BaseRequest

	WidList []int64 `json:"widList,omitempty"`
	State   int     `json:"state,omitempty"`
}

type WeikeBatchFrozenWeikeResponse struct {
}

func CreateWeikeBatchFrozenWeikeRequest() (request *WeikeBatchFrozenWeikeRequest) {
	request = &WeikeBatchFrozenWeikeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/batchFrozenWeike", "api")
	request.Method = api.POST
	return
}

func CreateWeikeBatchFrozenWeikeResponse() (response *api.BaseResponse[WeikeBatchFrozenWeikeResponse]) {
	response = api.CreateResponse[WeikeBatchFrozenWeikeResponse](&WeikeBatchFrozenWeikeResponse{})
	return
}
