package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategorySkuspecDelete
// @id 1859
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1859?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品规格)
func (client *Client) GoodsCategorySkuspecDelete(request *GoodsCategorySkuspecDeleteRequest) (response *GoodsCategorySkuspecDeleteResponse, err error) {
	rpcResponse := CreateGoodsCategorySkuspecDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategorySkuspecDeleteRequest struct {
	*api.BaseRequest

	BasicInfo GoodsCategorySkuspecDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	SpecId    int64                                      `json:"specId,omitempty"`
}

type GoodsCategorySkuspecDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsCategorySkuspecDeleteRequest() (request *GoodsCategorySkuspecDeleteRequest) {
	request = &GoodsCategorySkuspecDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/skuspec/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategorySkuspecDeleteResponse() (response *api.BaseResponse[GoodsCategorySkuspecDeleteResponse]) {
	response = api.CreateResponse[GoodsCategorySkuspecDeleteResponse](&GoodsCategorySkuspecDeleteResponse{})
	return
}

type GoodsCategorySkuspecDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
