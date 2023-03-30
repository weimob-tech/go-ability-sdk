package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyDelete
// @id 1837
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1837?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品分组)
func (client *Client) GoodsClassifyDelete(request *GoodsClassifyDeleteRequest) (response *GoodsClassifyDeleteResponse, err error) {
	rpcResponse := CreateGoodsClassifyDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyDeleteRequest struct {
	*api.BaseRequest

	BasicInfo  GoodsClassifyDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	ClassifyId int64                               `json:"classifyId,omitempty"`
}

type GoodsClassifyDeleteResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGoodsClassifyDeleteRequest() (request *GoodsClassifyDeleteRequest) {
	request = &GoodsClassifyDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/classify/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyDeleteResponse() (response *api.BaseResponse[GoodsClassifyDeleteResponse]) {
	response = api.CreateResponse[GoodsClassifyDeleteResponse](&GoodsClassifyDeleteResponse{})
	return
}

type GoodsClassifyDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
