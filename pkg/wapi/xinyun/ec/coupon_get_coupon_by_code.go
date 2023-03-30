package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetCouponByCode
// @id 535
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=通过券码查询优惠券详情（C端）)
func (client *Client) CouponGetCouponByCode(request *CouponGetCouponByCodeRequest) (response *CouponGetCouponByCodeResponse, err error) {
	rpcResponse := CreateCouponGetCouponByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetCouponByCodeRequest struct {
	*api.BaseRequest

	Code       string `json:"code,omitempty"`
	UseScene   int64  `json:"useScene,omitempty"`
	StoreId    int64  `json:"storeId,omitempty"`
	CheckStore bool   `json:"checkStore,omitempty"`
}

type CouponGetCouponByCodeResponse struct {
	Name                string `json:"name,omitempty"`
	Type                int64  `json:"type,omitempty"`
	FavorRemark         string `json:"favorRemark,omitempty"`
	AcceptGoodsType     int64  `json:"acceptGoodsType,omitempty"`
	UseRemark           string `json:"useRemark,omitempty"`
	Code                string `json:"code,omitempty"`
	Wid                 string `json:"wid,omitempty"`
	UserNickName        string `json:"userNickName,omitempty"`
	CardTemplateId      int64  `json:"cardTemplateId,omitempty"`
	StartDate           int64  `json:"startDate,omitempty"`
	ExpDate             int64  `json:"expDate,omitempty"`
	ConsumeDate         int64  `json:"consumeDate,omitempty"`
	CashTicketCondition int64  `json:"cashTicketCondition,omitempty"`
	CashTicketAmt       int64  `json:"cashTicketAmt,omitempty"`
	Discount            int64  `json:"discount,omitempty"`
	Status              int64  `json:"status,omitempty"`
	UnAvailableReason   string `json:"unAvailableReason,omitempty"`
	OrderId             string `json:"orderId,omitempty"`
	UseScene            int64  `json:"useScene,omitempty"`
	UseSceneList        int64  `json:"useSceneList,omitempty"`
	SubCouponType       int64  `json:"subCouponType,omitempty"`
	AcquireStoreId      int64  `json:"acquireStoreId,omitempty"`
}

func CreateCouponGetCouponByCodeRequest() (request *CouponGetCouponByCodeRequest) {
	request = &CouponGetCouponByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/getCouponByCode", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetCouponByCodeResponse() (response *api.BaseResponse[CouponGetCouponByCodeResponse]) {
	response = api.CreateResponse[CouponGetCouponByCodeResponse](&CouponGetCouponByCodeResponse{})
	return
}
