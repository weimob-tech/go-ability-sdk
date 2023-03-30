package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQueryGoodsIdAndCodeWithPage
// @id 1266
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1266?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=滚动查询商品id和商品编码)
func (client *Client) GoodsQueryGoodsIdAndCodeWithPage(request *GoodsQueryGoodsIdAndCodeWithPageRequest) (response *GoodsQueryGoodsIdAndCodeWithPageResponse, err error) {
	rpcResponse := CreateGoodsQueryGoodsIdAndCodeWithPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQueryGoodsIdAndCodeWithPageRequest struct {
	*api.BaseRequest

	Input GoodsQueryGoodsIdAndCodeWithPageRequestInput `json:"input,omitempty"`
}

type GoodsQueryGoodsIdAndCodeWithPageResponse struct {
	Output GoodsQueryGoodsIdAndCodeWithPageResponseOutput `json:"output,omitempty"`
}

func CreateGoodsQueryGoodsIdAndCodeWithPageRequest() (request *GoodsQueryGoodsIdAndCodeWithPageRequest) {
	request = &GoodsQueryGoodsIdAndCodeWithPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/queryGoodsIdAndCodeWithPage", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsQueryGoodsIdAndCodeWithPageResponse() (response *api.BaseResponse[GoodsQueryGoodsIdAndCodeWithPageResponse]) {
	response = api.CreateResponse[GoodsQueryGoodsIdAndCodeWithPageResponse](&GoodsQueryGoodsIdAndCodeWithPageResponse{})
	return
}

type GoodsQueryGoodsIdAndCodeWithPageRequestInput struct {
}

type GoodsQueryGoodsIdAndCodeWithPageResponseOutput struct {
}
