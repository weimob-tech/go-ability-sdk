package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SkuGet
// @id 10
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商品规格信息)
func (client *Client) SkuGet(request *SkuGetRequest) (response *SkuGetResponse, err error) {
	rpcResponse := CreateSkuGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SkuGetRequest struct {
	*api.BaseRequest

	SpuCode string `json:"spu_code,omitempty"`
	SkuCode string `json:"sku_code,omitempty"`
}

type SkuGetResponse struct {
}

func CreateSkuGetRequest() (request *SkuGetRequest) {
	request = &SkuGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Sku/Get", "api")
	request.Method = api.POST
	return
}

func CreateSkuGetResponse() (response *api.BaseResponse[SkuGetResponse]) {
	response = api.CreateResponse[SkuGetResponse](&SkuGetResponse{})
	return
}
