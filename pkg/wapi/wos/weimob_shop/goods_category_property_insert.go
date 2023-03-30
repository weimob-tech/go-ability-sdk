package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategoryPropertyInsert
// @id 1587
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1587?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品参数)
func (client *Client) GoodsCategoryPropertyInsert(request *GoodsCategoryPropertyInsertRequest) (response *GoodsCategoryPropertyInsertResponse, err error) {
	rpcResponse := CreateGoodsCategoryPropertyInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategoryPropertyInsertRequest struct {
	*api.BaseRequest

	BasicInfo        GoodsCategoryPropertyInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	CategoryId       int64                                       `json:"categoryId,omitempty"`
	CategoryPropName string                                      `json:"categoryPropName,omitempty"`
}

type GoodsCategoryPropertyInsertResponse struct {
	Id           int64 `json:"id,omitempty"`
	ReturnResult bool  `json:"returnResult,omitempty"`
}

func CreateGoodsCategoryPropertyInsertRequest() (request *GoodsCategoryPropertyInsertRequest) {
	request = &GoodsCategoryPropertyInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/property/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategoryPropertyInsertResponse() (response *api.BaseResponse[GoodsCategoryPropertyInsertResponse]) {
	response = api.CreateResponse[GoodsCategoryPropertyInsertResponse](&GoodsCategoryPropertyInsertResponse{})
	return
}

type GoodsCategoryPropertyInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
