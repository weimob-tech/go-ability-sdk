package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategoryPropertyDelete
// @id 1857
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1857?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品参数)
func (client *Client) GoodsCategoryPropertyDelete(request *GoodsCategoryPropertyDeleteRequest) (response *GoodsCategoryPropertyDeleteResponse, err error) {
	rpcResponse := CreateGoodsCategoryPropertyDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategoryPropertyDeleteRequest struct {
	*api.BaseRequest

	BasicInfo      GoodsCategoryPropertyDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	CategoryPropId int64                                       `json:"categoryPropId,omitempty"`
}

type GoodsCategoryPropertyDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsCategoryPropertyDeleteRequest() (request *GoodsCategoryPropertyDeleteRequest) {
	request = &GoodsCategoryPropertyDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/property/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategoryPropertyDeleteResponse() (response *api.BaseResponse[GoodsCategoryPropertyDeleteResponse]) {
	response = api.CreateResponse[GoodsCategoryPropertyDeleteResponse](&GoodsCategoryPropertyDeleteResponse{})
	return
}

type GoodsCategoryPropertyDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
