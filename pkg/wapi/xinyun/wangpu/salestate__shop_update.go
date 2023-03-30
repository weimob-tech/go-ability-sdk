package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SalestateShopUpdate
// @id 288
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新门店商品上下架状态在售状态接口)
func (client *Client) SalestateShopUpdate(request *SalestateShopUpdateRequest) (response *SalestateShopUpdateResponse, err error) {
	rpcResponse := CreateSalestateShopUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SalestateShopUpdateRequest struct {
	*api.BaseRequest

	ShopId        int64 `json:"shop_id,omitempty"`
	ShopGoodsIds  []int `json:"shop_goods_ids,omitempty"`
	CanPutOnsale  bool  `json:"can_put_onsale,omitempty"`
	IsPutOnsale   bool  `json:"is_put_onsale,omitempty"`
	CanShopOnsale bool  `json:"can_shop_onsale,omitempty"`
	IsShopOnsale  bool  `json:"is_shop_onsale,omitempty"`
}

type SalestateShopUpdateResponse struct {
}

func CreateSalestateShopUpdateRequest() (request *SalestateShopUpdateRequest) {
	request = &SalestateShopUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "salestate/ShopUpdate", "api")
	request.Method = api.POST
	return
}

func CreateSalestateShopUpdateResponse() (response *api.BaseResponse[SalestateShopUpdateResponse]) {
	response = api.CreateResponse[SalestateShopUpdateResponse](&SalestateShopUpdateResponse{})
	return
}
