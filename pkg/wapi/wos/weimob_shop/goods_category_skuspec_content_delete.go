package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategorySkuspecContentDelete
// @id 1858
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1858?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品规格值)
func (client *Client) GoodsCategorySkuspecContentDelete(request *GoodsCategorySkuspecContentDeleteRequest) (response *GoodsCategorySkuspecContentDeleteResponse, err error) {
	rpcResponse := CreateGoodsCategorySkuspecContentDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategorySkuspecContentDeleteRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsCategorySkuspecContentDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	SpecValueId int64                                             `json:"specValueId,omitempty"`
}

type GoodsCategorySkuspecContentDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsCategorySkuspecContentDeleteRequest() (request *GoodsCategorySkuspecContentDeleteRequest) {
	request = &GoodsCategorySkuspecContentDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/skuspec/content/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategorySkuspecContentDeleteResponse() (response *api.BaseResponse[GoodsCategorySkuspecContentDeleteResponse]) {
	response = api.CreateResponse[GoodsCategorySkuspecContentDeleteResponse](&GoodsCategorySkuspecContentDeleteResponse{})
	return
}

type GoodsCategorySkuspecContentDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
