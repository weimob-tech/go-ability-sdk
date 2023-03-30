package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponcodeValidCouponCodeDetail
// @id 1710
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠码详情)
func (client *Client) CouponcodeValidCouponCodeDetail(request *CouponcodeValidCouponCodeDetailRequest) (response *CouponcodeValidCouponCodeDetailResponse, err error) {
	rpcResponse := CreateCouponcodeValidCouponCodeDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponcodeValidCouponCodeDetailRequest struct {
	*api.BaseRequest

	StoreId int64  `json:"storeId,omitempty"`
	Code    string `json:"code,omitempty"`
}

type CouponcodeValidCouponCodeDetailResponse struct {
	CodeTemplateId     int64   `json:"codeTemplateId,omitempty"`
	Code               string  `json:"code,omitempty"`
	Name               string  `json:"name,omitempty"`
	CodeType           int     `json:"codeType,omitempty"`
	CouponType         int     `json:"couponType,omitempty"`
	CouponAmount       float64 `json:"couponAmount,omitempty"`
	UseCondition       float64 `json:"useCondition,omitempty"`
	UseLimitNum        int     `json:"useLimitNum,omitempty"`
	UseScenes          []int   `json:"useScenes,omitempty"`
	ExpireTimeBegin    int     `json:"expireTimeBegin,omitempty"`
	ExpireTimeEnd      int     `json:"expireTimeEnd,omitempty"`
	AcceptGoodsType    int     `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIdList  []int   `json:"acceptGoodsIdList,omitempty"`
	ExistExcludeGoods  bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIdList []int   `json:"excludeGoodsIdList,omitempty"`
	SelectStoreType    int     `json:"selectStoreType,omitempty"`
	SelectStoreIdList  []int   `json:"selectStoreIdList,omitempty"`
	TotalStock         int     `json:"totalStock,omitempty"`
	Status             int     `json:"status,omitempty"`
}

func CreateCouponcodeValidCouponCodeDetailRequest() (request *CouponcodeValidCouponCodeDetailRequest) {
	request = &CouponcodeValidCouponCodeDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "couponcode/validCouponCodeDetail", "api")
	request.Method = api.POST
	return
}

func CreateCouponcodeValidCouponCodeDetailResponse() (response *api.BaseResponse[CouponcodeValidCouponCodeDetailResponse]) {
	response = api.CreateResponse[CouponcodeValidCouponCodeDetailResponse](&CouponcodeValidCouponCodeDetailResponse{})
	return
}
