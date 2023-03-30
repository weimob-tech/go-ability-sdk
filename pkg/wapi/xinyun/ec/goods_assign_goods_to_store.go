package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAssignGoodsToStore
// @id 1676
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=分配商品)
func (client *Client) GoodsAssignGoodsToStore(request *GoodsAssignGoodsToStoreRequest) (response *GoodsAssignGoodsToStoreResponse, err error) {
	rpcResponse := CreateGoodsAssignGoodsToStoreResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAssignGoodsToStoreRequest struct {
	*api.BaseRequest

	StoreIdList []int `json:"storeIdList,omitempty"`
	GoodsId     int64 `json:"goodsId,omitempty"`
	AssignType  int   `json:"assignType,omitempty"`
}

type GoodsAssignGoodsToStoreResponse struct {
}

func CreateGoodsAssignGoodsToStoreRequest() (request *GoodsAssignGoodsToStoreRequest) {
	request = &GoodsAssignGoodsToStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/assignGoodsToStore", "api")
	request.Method = api.POST
	return
}

func CreateGoodsAssignGoodsToStoreResponse() (response *api.BaseResponse[GoodsAssignGoodsToStoreResponse]) {
	response = api.CreateResponse[GoodsAssignGoodsToStoreResponse](&GoodsAssignGoodsToStoreResponse{})
	return
}
