package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityQueryRechargeRewardDetail
// @id 1877
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据订单号查询充值活动赠送的奖品)
func (client *Client) ActivityQueryRechargeRewardDetail(request *ActivityQueryRechargeRewardDetailRequest) (response *ActivityQueryRechargeRewardDetailResponse, err error) {
	rpcResponse := CreateActivityQueryRechargeRewardDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityQueryRechargeRewardDetailRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
	Wid     int    `json:"wid,omitempty"`
}

type ActivityQueryRechargeRewardDetailResponse struct {
}

func CreateActivityQueryRechargeRewardDetailRequest() (request *ActivityQueryRechargeRewardDetailRequest) {
	request = &ActivityQueryRechargeRewardDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "activity/queryRechargeRewardDetail", "api")
	request.Method = api.POST
	return
}

func CreateActivityQueryRechargeRewardDetailResponse() (response *api.BaseResponse[ActivityQueryRechargeRewardDetailResponse]) {
	response = api.CreateResponse[ActivityQueryRechargeRewardDetailResponse](&ActivityQueryRechargeRewardDetailResponse{})
	return
}
