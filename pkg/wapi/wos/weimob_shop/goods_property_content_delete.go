package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyContentDelete
// @id 1864
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1864?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品属性值)
func (client *Client) GoodsPropertyContentDelete(request *GoodsPropertyContentDeleteRequest) (response *GoodsPropertyContentDeleteResponse, err error) {
	rpcResponse := CreateGoodsPropertyContentDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyContentDeleteRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsPropertyContentDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	PropValueId int64                                      `json:"propValueId,omitempty"`
	PropId      int64                                      `json:"propId,omitempty"`
}

type GoodsPropertyContentDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsPropertyContentDeleteRequest() (request *GoodsPropertyContentDeleteRequest) {
	request = &GoodsPropertyContentDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/content/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyContentDeleteResponse() (response *api.BaseResponse[GoodsPropertyContentDeleteResponse]) {
	response = api.CreateResponse[GoodsPropertyContentDeleteResponse](&GoodsPropertyContentDeleteResponse{})
	return
}

type GoodsPropertyContentDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
