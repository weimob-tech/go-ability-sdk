package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsDelete
// @id 1633
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1633?tab=1)
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

	BasicInfo   GoodsDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsIdList []int64                     `json:"goodsIdList,omitempty"`
}

type GoodsDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsDeleteRequest() (request *GoodsDeleteRequest) {
	request = &GoodsDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsDeleteResponse() (response *api.BaseResponse[GoodsDeleteResponse]) {
	response = api.CreateResponse[GoodsDeleteResponse](&GoodsDeleteResponse{})
	return
}

type GoodsDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
