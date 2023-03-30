package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShopStartOrCloseShop
// @id 285
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=启用/停用门店)
func (client *Client) ShopStartOrCloseShop(request *ShopStartOrCloseShopRequest) (response *ShopStartOrCloseShopResponse, err error) {
	rpcResponse := CreateShopStartOrCloseShopResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShopStartOrCloseShopRequest struct {
	*api.BaseRequest

	ShopId int64 `json:"shop_id,omitempty"`
	Status int   `json:"status,omitempty"`
}

type ShopStartOrCloseShopResponse struct {
}

func CreateShopStartOrCloseShopRequest() (request *ShopStartOrCloseShopRequest) {
	request = &ShopStartOrCloseShopRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "shop/StartOrCloseShop", "api")
	request.Method = api.POST
	return
}

func CreateShopStartOrCloseShopResponse() (response *api.BaseResponse[ShopStartOrCloseShopResponse]) {
	response = api.CreateResponse[ShopStartOrCloseShopResponse](&ShopStartOrCloseShopResponse{})
	return
}
