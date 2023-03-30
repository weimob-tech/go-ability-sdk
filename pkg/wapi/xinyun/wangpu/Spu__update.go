package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SpuUpdate
// @id 19
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新某些商品字段)
func (client *Client) SpuUpdate(request *SpuUpdateRequest) (response *SpuUpdateResponse, err error) {
	rpcResponse := CreateSpuUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SpuUpdateRequest struct {
	*api.BaseRequest

	SpuInfo SpuUpdateRequestSpu_info `json:"spu_info,omitempty"`
}

type SpuUpdateResponse struct {
}

func CreateSpuUpdateRequest() (request *SpuUpdateRequest) {
	request = &SpuUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Spu/Update", "api")
	request.Method = api.POST
	return
}

func CreateSpuUpdateResponse() (response *api.BaseResponse[SpuUpdateResponse]) {
	response = api.CreateResponse[SpuUpdateResponse](&SpuUpdateResponse{})
	return
}

type SpuUpdateRequestSpu_info struct {
}
