package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WaySaveResource
// @id 1000
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/编辑线路资源)
func (client *Client) WaySaveResource(request *WaySaveResourceRequest) (response *WaySaveResourceResponse, err error) {
	rpcResponse := CreateWaySaveResourceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WaySaveResourceRequest struct {
	*api.BaseRequest

	GoodsCode     string `json:"goodsCode,omitempty"`
	InsuranceIds  string `json:"insuranceIds,omitempty"`
	AdditionalIds string `json:"additionalIds,omitempty"`
}

type WaySaveResourceResponse struct {
}

func CreateWaySaveResourceRequest() (request *WaySaveResourceRequest) {
	request = &WaySaveResourceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "way/saveResource", "api")
	request.Method = api.POST
	return
}

func CreateWaySaveResourceResponse() (response *api.BaseResponse[WaySaveResourceResponse]) {
	response = api.CreateResponse[WaySaveResourceResponse](&WaySaveResourceResponse{})
	return
}
