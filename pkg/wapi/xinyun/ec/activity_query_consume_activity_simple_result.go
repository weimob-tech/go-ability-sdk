package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityQueryConsumeActivitySimpleResult
// @id 1823
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询参与活动的订单号)
func (client *Client) ActivityQueryConsumeActivitySimpleResult(request *ActivityQueryConsumeActivitySimpleResultRequest) (response *ActivityQueryConsumeActivitySimpleResultResponse, err error) {
	rpcResponse := CreateActivityQueryConsumeActivitySimpleResultResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityQueryConsumeActivitySimpleResultRequest struct {
	*api.BaseRequest

	Pid        int64  `json:"pid,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"`
	ActivityId int64  `json:"activityId,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}

type ActivityQueryConsumeActivitySimpleResultResponse struct {
}

func CreateActivityQueryConsumeActivitySimpleResultRequest() (request *ActivityQueryConsumeActivitySimpleResultRequest) {
	request = &ActivityQueryConsumeActivitySimpleResultRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "activity/queryConsumeActivitySimpleResult", "api")
	request.Method = api.POST
	return
}

func CreateActivityQueryConsumeActivitySimpleResultResponse() (response *api.BaseResponse[ActivityQueryConsumeActivitySimpleResultResponse]) {
	response = api.CreateResponse[ActivityQueryConsumeActivitySimpleResultResponse](&ActivityQueryConsumeActivitySimpleResultResponse{})
	return
}
