package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShareprofitInitShareProfit
// @id 1518
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=预分账明细通知)
func (client *Client) ShareprofitInitShareProfit(request *ShareprofitInitShareProfitRequest) (response *ShareprofitInitShareProfitResponse, err error) {
	rpcResponse := CreateShareprofitInitShareProfitResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShareprofitInitShareProfitRequest struct {
	*api.BaseRequest

	DetailList []ShareprofitInitShareProfitRequestDetailList `json:"detailList,omitempty"`
	StoreId    int64                                         `json:"storeId,omitempty"`
	OrderNo    int64                                         `json:"orderNo,omitempty"`
}

type ShareprofitInitShareProfitResponse struct {
}

func CreateShareprofitInitShareProfitRequest() (request *ShareprofitInitShareProfitRequest) {
	request = &ShareprofitInitShareProfitRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "shareprofit/initShareProfit", "api")
	request.Method = api.POST
	return
}

func CreateShareprofitInitShareProfitResponse() (response *api.BaseResponse[ShareprofitInitShareProfitResponse]) {
	response = api.CreateResponse[ShareprofitInitShareProfitResponse](&ShareprofitInitShareProfitResponse{})
	return
}

type ShareprofitInitShareProfitRequestDetailList struct {
	ShareProfitBizType   int     `json:"shareProfitBizType,omitempty"`
	ShareAmount          float64 `json:"shareAmount,omitempty"`
	ReceiverId           string  `json:"receiverId,omitempty"`
	ShareProfitExtraInfo string  `json:"shareProfitExtraInfo,omitempty"`
}
