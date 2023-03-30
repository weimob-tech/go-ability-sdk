package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClear
// @id 68
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=菜品估清)
func (client *Client) GoodsClear(request *GoodsClearRequest) (response *GoodsClearResponse, err error) {
	rpcResponse := CreateGoodsClearResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClearRequest struct {
	*api.BaseRequest

	GoodsId    int64 `json:"goodsId,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
}

type GoodsClearResponse struct {
}

func CreateGoodsClearRequest() (request *GoodsClearRequest) {
	request = &GoodsClearRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/clear", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClearResponse() (response *api.BaseResponse[GoodsClearResponse]) {
	response = api.CreateResponse[GoodsClearResponse](&GoodsClearResponse{})
	return
}
