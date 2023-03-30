package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUpdateCouponStatus
// @id 1700
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新券码状态)
func (client *Client) CouponUpdateCouponStatus(request *CouponUpdateCouponStatusRequest) (response *CouponUpdateCouponStatusResponse, err error) {
	rpcResponse := CreateCouponUpdateCouponStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUpdateCouponStatusRequest struct {
	*api.BaseRequest

	UserCouponList []CouponUpdateCouponStatusRequestUserCouponList `json:"userCouponList,omitempty"`
	InvalidCoupon  CouponUpdateCouponStatusRequestInvalidCoupon    `json:"invalidCoupon,omitempty"`
	ActionId       int                                             `json:"actionId,omitempty"`
	ConsumeCoupon  int64                                           `json:"consumeCoupon,omitempty"`
}

type CouponUpdateCouponStatusResponse struct {
}

func CreateCouponUpdateCouponStatusRequest() (request *CouponUpdateCouponStatusRequest) {
	request = &CouponUpdateCouponStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/updateCouponStatus", "api")
	request.Method = api.POST
	return
}

func CreateCouponUpdateCouponStatusResponse() (response *api.BaseResponse[CouponUpdateCouponStatusResponse]) {
	response = api.CreateResponse[CouponUpdateCouponStatusResponse](&CouponUpdateCouponStatusResponse{})
	return
}

type CouponUpdateCouponStatusRequestUserCouponList struct {
	GoodsList      []CouponUpdateCouponStatusRequestGoodsList `json:"goodsList,omitempty"`
	StoreId        int64                                      `json:"storeId,omitempty"`
	Wid            int64                                      `json:"wid,omitempty"`
	CardTemplateId int64                                      `json:"cardTemplateId,omitempty"`
	Code           string                                     `json:"code,omitempty"`
	OrderId        string                                     `json:"orderId,omitempty"`
	Amount         float64                                    `json:"amount,omitempty"`
	UseScene       int64                                      `json:"useScene,omitempty"`
}

type CouponUpdateCouponStatusRequestGoodsList struct {
	GoodsId        int64   `json:"goodsId,omitempty"`
	CategoryIdList []int64 `json:"categoryIdList,omitempty"`
	ClassifyIdList []int64 `json:"classifyIdList,omitempty"`
}

type CouponUpdateCouponStatusRequestInvalidCoupon struct {
	Code string `json:"code,omitempty"`
}
