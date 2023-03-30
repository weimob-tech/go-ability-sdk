package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUpdateCouponStock
// @id 1097
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券库存)
func (client *Client) CouponUpdateCouponStock(request *CouponUpdateCouponStockRequest) (response *CouponUpdateCouponStockResponse, err error) {
	rpcResponse := CreateCouponUpdateCouponStockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUpdateCouponStockRequest struct {
	*api.BaseRequest

	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
	StockNumChange int   `json:"stockNumChange,omitempty"`
}

type CouponUpdateCouponStockResponse struct {
}

func CreateCouponUpdateCouponStockRequest() (request *CouponUpdateCouponStockRequest) {
	request = &CouponUpdateCouponStockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "coupon/updateCouponStock", "api")
	request.Method = api.POST
	return
}

func CreateCouponUpdateCouponStockResponse() (response *api.BaseResponse[CouponUpdateCouponStockResponse]) {
	response = api.CreateResponse[CouponUpdateCouponStockResponse](&CouponUpdateCouponStockResponse{})
	return
}
