package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponcodeGetValidCouponCode
// @id 1526
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠码详情)
func (client *Client) CouponcodeGetValidCouponCode(request *CouponcodeGetValidCouponCodeRequest) (response *CouponcodeGetValidCouponCodeResponse, err error) {
	rpcResponse := CreateCouponcodeGetValidCouponCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponcodeGetValidCouponCodeRequest struct {
	*api.BaseRequest

	StoreId int64  `json:"storeId,omitempty"`
	Code    string `json:"code,omitempty"`
}

type CouponcodeGetValidCouponCodeResponse struct {
	CodeTemplateId  int64   `json:"codeTemplateId,omitempty"`
	Code            string  `json:"code,omitempty"`
	Name            string  `json:"name,omitempty"`
	CodeType        int     `json:"codeType,omitempty"`
	CouponType      int     `json:"couponType,omitempty"`
	CouponAmount    float64 `json:"couponAmount,omitempty"`
	UseCondition    float64 `json:"useCondition,omitempty"`
	UseLimitNum     int     `json:"useLimitNum,omitempty"`
	UseScenes       []int   `json:"useScenes,omitempty"`
	ExpireTimeBegin int     `json:"expireTimeBegin,omitempty"`
	ExpireTimeEnd   int     `json:"expireTimeEnd,omitempty"`
	AcceptGoodsType int     `json:"acceptGoodsType,omitempty"`
	SelectStoreType int     `json:"selectStoreType,omitempty"`
	TotalStock      int     `json:"totalStock,omitempty"`
	Stock           int     `json:"stock,omitempty"`
}

func CreateCouponcodeGetValidCouponCodeRequest() (request *CouponcodeGetValidCouponCodeRequest) {
	request = &CouponcodeGetValidCouponCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "couponcode/getValidCouponCode", "api")
	request.Method = api.POST
	return
}

func CreateCouponcodeGetValidCouponCodeResponse() (response *api.BaseResponse[CouponcodeGetValidCouponCodeResponse]) {
	response = api.CreateResponse[CouponcodeGetValidCouponCodeResponse](&CouponcodeGetValidCouponCodeResponse{})
	return
}
