package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategorySkuspecInsert
// @id 1589
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1589?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品规格)
func (client *Client) GoodsCategorySkuspecInsert(request *GoodsCategorySkuspecInsertRequest) (response *GoodsCategorySkuspecInsertResponse, err error) {
	rpcResponse := CreateGoodsCategorySkuspecInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategorySkuspecInsertRequest struct {
	*api.BaseRequest

	BasicInfo  GoodsCategorySkuspecInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	CategoryId int64                                      `json:"categoryId,omitempty"`
	SpecName   string                                     `json:"specName,omitempty"`
}

type GoodsCategorySkuspecInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateGoodsCategorySkuspecInsertRequest() (request *GoodsCategorySkuspecInsertRequest) {
	request = &GoodsCategorySkuspecInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/skuspec/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategorySkuspecInsertResponse() (response *api.BaseResponse[GoodsCategorySkuspecInsertResponse]) {
	response = api.CreateResponse[GoodsCategorySkuspecInsertResponse](&GoodsCategorySkuspecInsertResponse{})
	return
}

type GoodsCategorySkuspecInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
