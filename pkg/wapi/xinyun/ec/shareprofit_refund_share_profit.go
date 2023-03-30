package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShareprofitRefundShareProfit
// @id 1519
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=维权分账通知)
func (client *Client) ShareprofitRefundShareProfit(request *ShareprofitRefundShareProfitRequest) (response *ShareprofitRefundShareProfitResponse, err error) {
	rpcResponse := CreateShareprofitRefundShareProfitResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShareprofitRefundShareProfitRequest struct {
	*api.BaseRequest

	DetailList []ShareprofitRefundShareProfitRequestDetailList `json:"detailList,omitempty"`
	StoreId    int64                                           `json:"storeId,omitempty"`
	OrderNo    int64                                           `json:"orderNo,omitempty"`
	RightsNo   int64                                           `json:"rightsNo,omitempty"`
}

type ShareprofitRefundShareProfitResponse struct {
}

func CreateShareprofitRefundShareProfitRequest() (request *ShareprofitRefundShareProfitRequest) {
	request = &ShareprofitRefundShareProfitRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "shareprofit/refundShareProfit", "api")
	request.Method = api.POST
	return
}

func CreateShareprofitRefundShareProfitResponse() (response *api.BaseResponse[ShareprofitRefundShareProfitResponse]) {
	response = api.CreateResponse[ShareprofitRefundShareProfitResponse](&ShareprofitRefundShareProfitResponse{})
	return
}

type ShareprofitRefundShareProfitRequestDetailList struct {
	ShareProfitBizType   int     `json:"shareProfitBizType,omitempty"`
	RefundShareAmount    float64 `json:"refundShareAmount,omitempty"`
	ReceiverId           string  `json:"receiverId,omitempty"`
	ShareProfitExtraInfo string  `json:"shareProfitExtraInfo,omitempty"`
}
