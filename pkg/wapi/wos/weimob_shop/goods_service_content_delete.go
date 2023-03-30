package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsServiceContentDelete
// @id 1838
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1838?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品服务值)
func (client *Client) GoodsServiceContentDelete(request *GoodsServiceContentDeleteRequest) (response *GoodsServiceContentDeleteResponse, err error) {
	rpcResponse := CreateGoodsServiceContentDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsServiceContentDeleteRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsServiceContentDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	PropValueId int64                                     `json:"propValueId,omitempty"`
	PropId      int64                                     `json:"propId,omitempty"`
}

type GoodsServiceContentDeleteResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGoodsServiceContentDeleteRequest() (request *GoodsServiceContentDeleteRequest) {
	request = &GoodsServiceContentDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/service/content/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsServiceContentDeleteResponse() (response *api.BaseResponse[GoodsServiceContentDeleteResponse]) {
	response = api.CreateResponse[GoodsServiceContentDeleteResponse](&GoodsServiceContentDeleteResponse{})
	return
}

type GoodsServiceContentDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
