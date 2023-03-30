package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetPublishRecordList
// @id 1537
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取码库列表)
func (client *Client) CouponGetPublishRecordList(request *CouponGetPublishRecordListRequest) (response *CouponGetPublishRecordListResponse, err error) {
	rpcResponse := CreateCouponGetPublishRecordListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetPublishRecordListRequest struct {
	*api.BaseRequest

	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	PageIndex      int    `json:"pageIndex,omitempty"`
	PageSize       int    `json:"pageSize,omitempty"`
	Keyword        string `json:"keyword,omitempty"`
}

type CouponGetPublishRecordListResponse struct {
}

func CreateCouponGetPublishRecordListRequest() (request *CouponGetPublishRecordListRequest) {
	request = &CouponGetPublishRecordListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "coupon/getPublishRecordList", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetPublishRecordListResponse() (response *api.BaseResponse[CouponGetPublishRecordListResponse]) {
	response = api.CreateResponse[CouponGetPublishRecordListResponse](&CouponGetPublishRecordListResponse{})
	return
}
