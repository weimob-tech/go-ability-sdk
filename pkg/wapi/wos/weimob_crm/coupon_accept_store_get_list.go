package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponAcceptStoreGetList
// @id 2089
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2089?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询券适用门店信息)
func (client *Client) CouponAcceptStoreGetList(request *CouponAcceptStoreGetListRequest) (response *CouponAcceptStoreGetListResponse, err error) {
	rpcResponse := CreateCouponAcceptStoreGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponAcceptStoreGetListRequest struct {
	*api.BaseRequest

	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
}

type CouponAcceptStoreGetListResponse struct {
	VidInfoList []CouponAcceptStoreGetListResponseVidInfoList `json:"vidInfoList,omitempty"`
}

func CreateCouponAcceptStoreGetListRequest() (request *CouponAcceptStoreGetListRequest) {
	request = &CouponAcceptStoreGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/accept/store/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponAcceptStoreGetListResponse() (response *api.BaseResponse[CouponAcceptStoreGetListResponse]) {
	response = api.CreateResponse[CouponAcceptStoreGetListResponse](&CouponAcceptStoreGetListResponse{})
	return
}

type CouponAcceptStoreGetListResponseVidInfoList struct {
	Vid     int64  `json:"vid,omitempty"`
	VidCode string `json:"vidCode,omitempty"`
	VidType int64  `json:"vidType,omitempty"`
}
