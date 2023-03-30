package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupList
// @id 57
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询菜品分组)
func (client *Client) GoodsGroupList(request *GoodsGroupListRequest) (response *GoodsGroupListResponse, err error) {
	rpcResponse := CreateGoodsGroupListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupListRequest struct {
	*api.BaseRequest

	StoreId       int64 `json:"storeId,omitempty"`
	GoodsGroupId  int64 `json:"goodsGroupId,omitempty"`
	MerchantId    int64 `json:"merchantId,omitempty"`
	SuppOrderType int   `json:"suppOrderType,omitempty"`
}

type GoodsGroupListResponse struct {
}

func CreateGoodsGroupListRequest() (request *GoodsGroupListRequest) {
	request = &GoodsGroupListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsGroup/list", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGroupListResponse() (response *api.BaseResponse[GoodsGroupListResponse]) {
	response = api.CreateResponse[GoodsGroupListResponse](&GoodsGroupListResponse{})
	return
}
