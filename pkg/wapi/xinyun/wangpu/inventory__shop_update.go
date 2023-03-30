package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// InventoryShopUpdate
// @id 328
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=门店库存接口)
func (client *Client) InventoryShopUpdate(request *InventoryShopUpdateRequest) (response *InventoryShopUpdateResponse, err error) {
	rpcResponse := CreateInventoryShopUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type InventoryShopUpdateRequest struct {
	*api.BaseRequest

	SkuList []map[string]any `json:"sku_list,omitempty"`
	ShopId  int64            `json:"shop_id,omitempty"`
	SpuCode string           `json:"spu_code,omitempty"`
}

type InventoryShopUpdateResponse struct {
}

func CreateInventoryShopUpdateRequest() (request *InventoryShopUpdateRequest) {
	request = &InventoryShopUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "inventory/ShopUpdate", "api")
	request.Method = api.POST
	return
}

func CreateInventoryShopUpdateResponse() (response *api.BaseResponse[InventoryShopUpdateResponse]) {
	response = api.CreateResponse[InventoryShopUpdateResponse](&InventoryShopUpdateResponse{})
	return
}
