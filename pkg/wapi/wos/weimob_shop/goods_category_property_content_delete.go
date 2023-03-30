package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategoryPropertyContentDelete
// @id 1856
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1856?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品参数值)
func (client *Client) GoodsCategoryPropertyContentDelete(request *GoodsCategoryPropertyContentDeleteRequest) (response *GoodsCategoryPropertyContentDeleteResponse, err error) {
	rpcResponse := CreateGoodsCategoryPropertyContentDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategoryPropertyContentDeleteRequest struct {
	*api.BaseRequest

	BasicInfo           GoodsCategoryPropertyContentDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	CategoryPropValueId int64                                              `json:"categoryPropValueId,omitempty"`
}

type GoodsCategoryPropertyContentDeleteResponse struct {
	ReturnResult bool   `json:"returnResult,omitempty"`
	Id           int64  `json:"id,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func CreateGoodsCategoryPropertyContentDeleteRequest() (request *GoodsCategoryPropertyContentDeleteRequest) {
	request = &GoodsCategoryPropertyContentDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/property/content/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategoryPropertyContentDeleteResponse() (response *api.BaseResponse[GoodsCategoryPropertyContentDeleteResponse]) {
	response = api.CreateResponse[GoodsCategoryPropertyContentDeleteResponse](&GoodsCategoryPropertyContentDeleteResponse{})
	return
}

type GoodsCategoryPropertyContentDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
