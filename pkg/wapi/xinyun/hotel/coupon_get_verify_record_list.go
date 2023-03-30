package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetVerifyRecordList
// @id 1535
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取优惠券核销纪录)
func (client *Client) CouponGetVerifyRecordList(request *CouponGetVerifyRecordListRequest) (response *CouponGetVerifyRecordListResponse, err error) {
	rpcResponse := CreateCouponGetVerifyRecordListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetVerifyRecordListRequest struct {
	*api.BaseRequest

	StoreId   int64  `json:"storeId,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
	StartTime int    `json:"startTime,omitempty"`
	EndTime   int    `json:"endTime,omitempty"`
	Keyword   string `json:"keyword,omitempty"`
}

type CouponGetVerifyRecordListResponse struct {
}

func CreateCouponGetVerifyRecordListRequest() (request *CouponGetVerifyRecordListRequest) {
	request = &CouponGetVerifyRecordListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "coupon/getVerifyRecordList", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetVerifyRecordListResponse() (response *api.BaseResponse[CouponGetVerifyRecordListResponse]) {
	response = api.CreateResponse[CouponGetVerifyRecordListResponse](&CouponGetVerifyRecordListResponse{})
	return
}
