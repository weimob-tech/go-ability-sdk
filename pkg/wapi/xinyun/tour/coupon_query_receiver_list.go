package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponQueryReceiverList
// @id 1094
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询领取人列表)
func (client *Client) CouponQueryReceiverList(request *CouponQueryReceiverListRequest) (response *CouponQueryReceiverListResponse, err error) {
	rpcResponse := CreateCouponQueryReceiverListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponQueryReceiverListRequest struct {
	*api.BaseRequest

	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	Status         int    `json:"status,omitempty"`
	TransactionId  string `json:"transactionId,omitempty"`
	PageIndex      int    `json:"pageIndex,omitempty"`
	PageSize       int    `json:"pageSize,omitempty"`
}

type CouponQueryReceiverListResponse struct {
}

func CreateCouponQueryReceiverListRequest() (request *CouponQueryReceiverListRequest) {
	request = &CouponQueryReceiverListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "coupon/queryReceiverList", "api")
	request.Method = api.POST
	return
}

func CreateCouponQueryReceiverListResponse() (response *api.BaseResponse[CouponQueryReceiverListResponse]) {
	response = api.CreateResponse[CouponQueryReceiverListResponse](&CouponQueryReceiverListResponse{})
	return
}
