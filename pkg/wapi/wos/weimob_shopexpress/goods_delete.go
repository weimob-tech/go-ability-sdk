package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsDelete
// @id 2110
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2110?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品)
func (client *Client) GoodsDelete(request *GoodsDeleteRequest) (response *GoodsDeleteResponse, err error) {
	rpcResponse := CreateGoodsDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsDeleteRequest struct {
	*api.BaseRequest

	GoodsCodeList []int64 `json:"goodsCodeList,omitempty"`
}

type GoodsDeleteResponse struct {
}

func CreateGoodsDeleteRequest() (request *GoodsDeleteRequest) {
	request = &GoodsDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsDeleteResponse() (response *api.BaseResponse[GoodsDeleteResponse]) {
	response = api.CreateResponse[GoodsDeleteResponse](&GoodsDeleteResponse{})
	return
}
