package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShareprofitConfirmShareProfit
// @id 1520
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=确认分账通知)
func (client *Client) ShareprofitConfirmShareProfit(request *ShareprofitConfirmShareProfitRequest) (response *ShareprofitConfirmShareProfitResponse, err error) {
	rpcResponse := CreateShareprofitConfirmShareProfitResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShareprofitConfirmShareProfitRequest struct {
	*api.BaseRequest

	DetailList []ShareprofitConfirmShareProfitRequestDetailList `json:"detailList,omitempty"`
	StoreId    int64                                            `json:"storeId,omitempty"`
	OrderNo    int64                                            `json:"orderNo,omitempty"`
}

type ShareprofitConfirmShareProfitResponse struct {
}

func CreateShareprofitConfirmShareProfitRequest() (request *ShareprofitConfirmShareProfitRequest) {
	request = &ShareprofitConfirmShareProfitRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "shareprofit/confirmShareProfit", "api")
	request.Method = api.POST
	return
}

func CreateShareprofitConfirmShareProfitResponse() (response *api.BaseResponse[ShareprofitConfirmShareProfitResponse]) {
	response = api.CreateResponse[ShareprofitConfirmShareProfitResponse](&ShareprofitConfirmShareProfitResponse{})
	return
}

type ShareprofitConfirmShareProfitRequestDetailList struct {
	ShareProfitBizType   int     `json:"shareProfitBizType,omitempty"`
	ShareAmount          float64 `json:"shareAmount,omitempty"`
	ReceiverId           string  `json:"receiverId,omitempty"`
	ShareProfitExtraInfo string  `json:"shareProfitExtraInfo,omitempty"`
}
