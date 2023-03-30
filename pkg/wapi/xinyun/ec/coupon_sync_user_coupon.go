package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponSyncUserCoupon
// @id 1491
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=用户券信息同步)
func (client *Client) CouponSyncUserCoupon(request *CouponSyncUserCouponRequest) (response *CouponSyncUserCouponResponse, err error) {
	rpcResponse := CreateCouponSyncUserCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponSyncUserCouponRequest struct {
	*api.BaseRequest

	StoreId         int64  `json:"storeId,omitempty"`
	RecipientIdent  int    `json:"recipientIdent,omitempty"`
	CardTemplateId  int64  `json:"cardTemplateId,omitempty"`
	SendWid         int64  `json:"sendWid,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Code            string `json:"code,omitempty"`
	PublishTime     int64  `json:"publishTime,omitempty"`
	ExpireStartTime int64  `json:"expireStartTime,omitempty"`
	ExpireEndTime   int64  `json:"expireEndTime,omitempty"`
}

type CouponSyncUserCouponResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponSyncUserCouponRequest() (request *CouponSyncUserCouponRequest) {
	request = &CouponSyncUserCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/syncUserCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponSyncUserCouponResponse() (response *api.BaseResponse[CouponSyncUserCouponResponse]) {
	response = api.CreateResponse[CouponSyncUserCouponResponse](&CouponSyncUserCouponResponse{})
	return
}
