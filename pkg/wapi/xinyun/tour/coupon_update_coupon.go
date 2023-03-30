package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUpdateCoupon
// @id 1096
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券)
func (client *Client) CouponUpdateCoupon(request *CouponUpdateCouponRequest) (response *CouponUpdateCouponResponse, err error) {
	rpcResponse := CreateCouponUpdateCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUpdateCouponRequest struct {
	*api.BaseRequest

	CardTemplateId int64   `json:"cardTemplateId,omitempty"`
	Version        int64   `json:"version,omitempty"`
	ColorType      int     `json:"colorType,omitempty"`
	UserTakeLimit  int     `json:"userTakeLimit,omitempty"`
	StartDate      int     `json:"startDate,omitempty"`
	ExpireDate     int     `json:"expireDate,omitempty"`
	MaxAmount      float64 `json:"maxAmount,omitempty"`
	UseNotice      string  `json:"useNotice,omitempty"`
	ServiceTel     string  `json:"serviceTel,omitempty"`
	LogoUrl        string  `json:"logoUrl,omitempty"`
	GoodsIds       []int64 `json:"goodsIds,omitempty"`
}

type CouponUpdateCouponResponse struct {
}

func CreateCouponUpdateCouponRequest() (request *CouponUpdateCouponRequest) {
	request = &CouponUpdateCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "coupon/updateCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponUpdateCouponResponse() (response *api.BaseResponse[CouponUpdateCouponResponse]) {
	response = api.CreateResponse[CouponUpdateCouponResponse](&CouponUpdateCouponResponse{})
	return
}
