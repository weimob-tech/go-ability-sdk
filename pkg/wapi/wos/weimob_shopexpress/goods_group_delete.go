package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupDelete
// @id 1975
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1975?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品分组删除)
func (client *Client) GoodsGroupDelete(request *GoodsGroupDeleteRequest) (response *GoodsGroupDeleteResponse, err error) {
	rpcResponse := CreateGoodsGroupDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupDeleteRequest struct {
	*api.BaseRequest

	GroupCode     string `json:"groupCode,omitempty"`
	OperationName string `json:"operationName,omitempty"`
}

type GoodsGroupDeleteResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateGoodsGroupDeleteRequest() (request *GoodsGroupDeleteRequest) {
	request = &GoodsGroupDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/group/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGroupDeleteResponse() (response *api.BaseResponse[GoodsGroupDeleteResponse]) {
	response = api.CreateResponse[GoodsGroupDeleteResponse](&GoodsGroupDeleteResponse{})
	return
}
