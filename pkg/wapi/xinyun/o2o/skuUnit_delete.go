package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SkuUnitDelete
// @id 265
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除菜品规格)
func (client *Client) SkuUnitDelete(request *SkuUnitDeleteRequest) (response *SkuUnitDeleteResponse, err error) {
	rpcResponse := CreateSkuUnitDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SkuUnitDeleteRequest struct {
	*api.BaseRequest

	SkuUnitId  int64 `json:"skuUnitId,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
}

type SkuUnitDeleteResponse struct {
}

func CreateSkuUnitDeleteRequest() (request *SkuUnitDeleteRequest) {
	request = &SkuUnitDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "skuUnit/delete", "api")
	request.Method = api.POST
	return
}

func CreateSkuUnitDeleteResponse() (response *api.BaseResponse[SkuUnitDeleteResponse]) {
	response = api.CreateResponse[SkuUnitDeleteResponse](&SkuUnitDeleteResponse{})
	return
}
