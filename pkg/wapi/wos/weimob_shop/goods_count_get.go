package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCountGet
// @id 1862
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1862?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品总数和上架数)
func (client *Client) GoodsCountGet(request *GoodsCountGetRequest) (response *GoodsCountGetResponse, err error) {
	rpcResponse := CreateGoodsCountGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCountGetRequest struct {
	*api.BaseRequest

	BasicInfo GoodsCountGetRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsId   int64                         `json:"goodsId,omitempty"`
}

type GoodsCountGetResponse struct {
	TotalGoodsCount  int64 `json:"totalGoodsCount,omitempty"`
	OnlineGoodsCount int64 `json:"onlineGoodsCount,omitempty"`
}

func CreateGoodsCountGetRequest() (request *GoodsCountGetRequest) {
	request = &GoodsCountGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/count/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCountGetResponse() (response *api.BaseResponse[GoodsCountGetResponse]) {
	response = api.CreateResponse[GoodsCountGetResponse](&GoodsCountGetResponse{})
	return
}

type GoodsCountGetRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
