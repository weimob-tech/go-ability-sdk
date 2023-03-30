package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupContains
// @id 73
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新分组内菜品（添加、删除）)
func (client *Client) GoodsGroupContains(request *GoodsGroupContainsRequest) (response *GoodsGroupContainsResponse, err error) {
	rpcResponse := CreateGoodsGroupContainsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupContainsRequest struct {
	*api.BaseRequest

	GoodsGroupId int64 `json:"goodsGroupId,omitempty"`
	StoreGoodsId int64 `json:"storeGoodsId,omitempty"`
	StoreId      int64 `json:"storeId,omitempty"`
	MerchantId   int64 `json:"merchantId,omitempty"`
}

type GoodsGroupContainsResponse struct {
}

func CreateGoodsGroupContainsRequest() (request *GoodsGroupContainsRequest) {
	request = &GoodsGroupContainsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsGroup/contains", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGroupContainsResponse() (response *api.BaseResponse[GoodsGroupContainsResponse]) {
	response = api.CreateResponse[GoodsGroupContainsResponse](&GoodsGroupContainsResponse{})
	return
}
