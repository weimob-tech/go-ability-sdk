package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategoryPropertyContentInsert
// @id 1588
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1588?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品参数值)
func (client *Client) GoodsCategoryPropertyContentInsert(request *GoodsCategoryPropertyContentInsertRequest) (response *GoodsCategoryPropertyContentInsertResponse, err error) {
	rpcResponse := CreateGoodsCategoryPropertyContentInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategoryPropertyContentInsertRequest struct {
	*api.BaseRequest

	BasicInfo             GoodsCategoryPropertyContentInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	CategoryPropId        int64                                              `json:"categoryPropId,omitempty"`
	CategoryPropValueName string                                             `json:"categoryPropValueName,omitempty"`
}

type GoodsCategoryPropertyContentInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateGoodsCategoryPropertyContentInsertRequest() (request *GoodsCategoryPropertyContentInsertRequest) {
	request = &GoodsCategoryPropertyContentInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/property/content/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategoryPropertyContentInsertResponse() (response *api.BaseResponse[GoodsCategoryPropertyContentInsertResponse]) {
	response = api.CreateResponse[GoodsCategoryPropertyContentInsertResponse](&GoodsCategoryPropertyContentInsertResponse{})
	return
}

type GoodsCategoryPropertyContentInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
