package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategorySkuspecContentInsert
// @id 1590
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1590?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品规格值)
func (client *Client) GoodsCategorySkuspecContentInsert(request *GoodsCategorySkuspecContentInsertRequest) (response *GoodsCategorySkuspecContentInsertResponse, err error) {
	rpcResponse := CreateGoodsCategorySkuspecContentInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategorySkuspecContentInsertRequest struct {
	*api.BaseRequest

	BasicInfo     GoodsCategorySkuspecContentInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	CategoryId    int64                                             `json:"categoryId,omitempty"`
	SpecId        int64                                             `json:"specId,omitempty"`
	SpecValueName string                                            `json:"specValueName,omitempty"`
}

type GoodsCategorySkuspecContentInsertResponse struct {
	Id           int64 `json:"id,omitempty"`
	ReturnResult bool  `json:"returnResult,omitempty"`
}

func CreateGoodsCategorySkuspecContentInsertRequest() (request *GoodsCategorySkuspecContentInsertRequest) {
	request = &GoodsCategorySkuspecContentInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/skuspec/content/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategorySkuspecContentInsertResponse() (response *api.BaseResponse[GoodsCategorySkuspecContentInsertResponse]) {
	response = api.CreateResponse[GoodsCategorySkuspecContentInsertResponse](&GoodsCategorySkuspecContentInsertResponse{})
	return
}

type GoodsCategorySkuspecContentInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
