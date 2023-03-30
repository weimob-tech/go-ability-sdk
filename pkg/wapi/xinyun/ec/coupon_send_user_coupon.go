package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponSendUserCoupon
// @id 1080
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=发送优惠券给指定客户)
func (client *Client) CouponSendUserCoupon(request *CouponSendUserCouponRequest) (response *CouponSendUserCouponResponse, err error) {
	rpcResponse := CreateCouponSendUserCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponSendUserCouponRequest struct {
	*api.BaseRequest

	StoreId        int64  `json:"storeId,omitempty"`
	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	RecipientIdent int    `json:"recipientIdent,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Code           string `json:"code,omitempty"`
	CreateNewUser  bool   `json:"createNewUser,omitempty"`
	SendWid        int64  `json:"sendWid,omitempty"`
	Pid            int64  `json:"pid,omitempty"`
	UnionId        string `json:"unionId,omitempty"`
}

type CouponSendUserCouponResponse struct {
	State      int    `json:"state,omitempty"`
	CanReceive int    `json:"canReceive,omitempty"`
	Code       string `json:"code,omitempty"`
}

func CreateCouponSendUserCouponRequest() (request *CouponSendUserCouponRequest) {
	request = &CouponSendUserCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/sendUserCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponSendUserCouponResponse() (response *api.BaseResponse[CouponSendUserCouponResponse]) {
	response = api.CreateResponse[CouponSendUserCouponResponse](&CouponSendUserCouponResponse{})
	return
}
