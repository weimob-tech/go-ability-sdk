package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsSkuSpecGetList
// @id 1254
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1254?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询规格列表)
func (client *Client) GoodsSkuSpecGetList(request *GoodsSkuSpecGetListRequest) (response *GoodsSkuSpecGetListResponse, err error) {
	rpcResponse := CreateGoodsSkuSpecGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsSkuSpecGetListRequest struct {
	*api.BaseRequest

	Input GoodsSkuSpecGetListRequestInput `json:"input,omitempty"`
}

type GoodsSkuSpecGetListResponse struct {
	Output GoodsSkuSpecGetListResponseOutput `json:"output,omitempty"`
}

func CreateGoodsSkuSpecGetListRequest() (request *GoodsSkuSpecGetListRequest) {
	request = &GoodsSkuSpecGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/skuSpec/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsSkuSpecGetListResponse() (response *api.BaseResponse[GoodsSkuSpecGetListResponse]) {
	response = api.CreateResponse[GoodsSkuSpecGetListResponse](&GoodsSkuSpecGetListResponse{})
	return
}

type GoodsSkuSpecGetListRequestInput struct {
}

type GoodsSkuSpecGetListResponseOutput struct {
}
