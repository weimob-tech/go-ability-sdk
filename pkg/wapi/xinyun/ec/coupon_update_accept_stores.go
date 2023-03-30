package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUpdateAcceptStores
// @id 1696
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券模板适用门店)
func (client *Client) CouponUpdateAcceptStores(request *CouponUpdateAcceptStoresRequest) (response *CouponUpdateAcceptStoresResponse, err error) {
	rpcResponse := CreateCouponUpdateAcceptStoresResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUpdateAcceptStoresRequest struct {
	*api.BaseRequest

	AcceptStoreIds  []string `json:"acceptStoreIds,omitempty"`
	StoreId         int64    `json:"storeId,omitempty"`
	CardTemplateId  int64    `json:"cardTemplateId,omitempty"`
	AcceptStoreType int      `json:"acceptStoreType,omitempty"`
}

type CouponUpdateAcceptStoresResponse struct {
}

func CreateCouponUpdateAcceptStoresRequest() (request *CouponUpdateAcceptStoresRequest) {
	request = &CouponUpdateAcceptStoresRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/updateAcceptStores", "api")
	request.Method = api.POST
	return
}

func CreateCouponUpdateAcceptStoresResponse() (response *api.BaseResponse[CouponUpdateAcceptStoresResponse]) {
	response = api.CreateResponse[CouponUpdateAcceptStoresResponse](&CouponUpdateAcceptStoresResponse{})
	return
}
