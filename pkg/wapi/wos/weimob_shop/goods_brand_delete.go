package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBrandDelete
// @id 1833
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1833?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品品牌)
func (client *Client) GoodsBrandDelete(request *GoodsBrandDeleteRequest) (response *GoodsBrandDeleteResponse, err error) {
	rpcResponse := CreateGoodsBrandDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBrandDeleteRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsBrandDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	BrandIdList []int64                          `json:"brandIdList,omitempty"`
}

type GoodsBrandDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsBrandDeleteRequest() (request *GoodsBrandDeleteRequest) {
	request = &GoodsBrandDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/brand/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsBrandDeleteResponse() (response *api.BaseResponse[GoodsBrandDeleteResponse]) {
	response = api.CreateResponse[GoodsBrandDeleteResponse](&GoodsBrandDeleteResponse{})
	return
}

type GoodsBrandDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
