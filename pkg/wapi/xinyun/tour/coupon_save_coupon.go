package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponSaveCoupon
// @id 1095
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加优惠券)
func (client *Client) CouponSaveCoupon(request *CouponSaveCouponRequest) (response *CouponSaveCouponResponse, err error) {
	rpcResponse := CreateCouponSaveCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponSaveCouponRequest struct {
	*api.BaseRequest

	Type               int     `json:"type,omitempty"`
	Title              string  `json:"title,omitempty"`
	Store              int     `json:"store,omitempty"`
	UserTakeLimit      int     `json:"userTakeLimit,omitempty"`
	UseAmountCondition float64 `json:"useAmountCondition,omitempty"`
	MaxAmount          float64 `json:"maxAmount,omitempty"`
	UseAmount          float64 `json:"useAmount,omitempty"`
	ExpireDateType     int     `json:"expireDateType,omitempty"`
	StartDate          int     `json:"startDate,omitempty"`
	ExpireDate         int     `json:"expireDate,omitempty"`
	StartDayCount      int     `json:"startDayCount,omitempty"`
	ExpDayCount        int     `json:"expDayCount,omitempty"`
	UseNotice          string  `json:"useNotice,omitempty"`
	ServiceTel         string  `json:"serviceTel,omitempty"`
	CanGiveFriend      bool    `json:"canGiveFriend,omitempty"`
	MerchantName       string  `json:"merchantName,omitempty"`
	LogoUrl            string  `json:"logoUrl,omitempty"`
	GoodsIds           []int64 `json:"goodsIds,omitempty"`
}

type CouponSaveCouponResponse struct {
}

func CreateCouponSaveCouponRequest() (request *CouponSaveCouponRequest) {
	request = &CouponSaveCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "coupon/saveCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponSaveCouponResponse() (response *api.BaseResponse[CouponSaveCouponResponse]) {
	response = api.CreateResponse[CouponSaveCouponResponse](&CouponSaveCouponResponse{})
	return
}
