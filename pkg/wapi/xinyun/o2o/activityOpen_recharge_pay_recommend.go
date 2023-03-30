package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityOpenRechargePayRecommend
// @id 1912
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=储值并买单推荐接口)
func (client *Client) ActivityOpenRechargePayRecommend(request *ActivityOpenRechargePayRecommendRequest) (response *ActivityOpenRechargePayRecommendResponse, err error) {
	rpcResponse := CreateActivityOpenRechargePayRecommendResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityOpenRechargePayRecommendRequest struct {
	*api.BaseRequest

	Wid                 int64 `json:"wid,omitempty"`
	RankId              int64 `json:"rankId,omitempty"`
	NeedToPay           int64 `json:"needToPay,omitempty"`
	DepositPayOrderType int   `json:"depositPayOrderType,omitempty"`
	DepositType         int   `json:"depositType,omitempty"`
	StoreId             int64 `json:"storeId,omitempty"`
}

type ActivityOpenRechargePayRecommendResponse struct {
}

func CreateActivityOpenRechargePayRecommendRequest() (request *ActivityOpenRechargePayRecommendRequest) {
	request = &ActivityOpenRechargePayRecommendRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "activityOpen/rechargePayRecommend", "api")
	request.Method = api.POST
	return
}

func CreateActivityOpenRechargePayRecommendResponse() (response *api.BaseResponse[ActivityOpenRechargePayRecommendResponse]) {
	response = api.CreateResponse[ActivityOpenRechargePayRecommendResponse](&ActivityOpenRechargePayRecommendResponse{})
	return
}
