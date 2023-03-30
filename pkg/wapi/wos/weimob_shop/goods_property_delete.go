package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyDelete
// @id 1865
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1865?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品属性)
func (client *Client) GoodsPropertyDelete(request *GoodsPropertyDeleteRequest) (response *GoodsPropertyDeleteResponse, err error) {
	rpcResponse := CreateGoodsPropertyDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyDeleteRequest struct {
	*api.BaseRequest

	BasicInfo GoodsPropertyDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	PropId    int64                               `json:"propId,omitempty"`
}

type GoodsPropertyDeleteResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGoodsPropertyDeleteRequest() (request *GoodsPropertyDeleteRequest) {
	request = &GoodsPropertyDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyDeleteResponse() (response *api.BaseResponse[GoodsPropertyDeleteResponse]) {
	response = api.CreateResponse[GoodsPropertyDeleteResponse](&GoodsPropertyDeleteResponse{})
	return
}

type GoodsPropertyDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
