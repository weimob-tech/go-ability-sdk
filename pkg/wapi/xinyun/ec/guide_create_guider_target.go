package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideCreateGuiderTarget
// @id 1253
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增导购目标)
func (client *Client) GuideCreateGuiderTarget(request *GuideCreateGuiderTargetRequest) (response *GuideCreateGuiderTargetResponse, err error) {
	rpcResponse := CreateGuideCreateGuiderTargetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideCreateGuiderTargetRequest struct {
	*api.BaseRequest

	GuiderWidList []int                                      `json:"guiderWidList,omitempty"`
	TargetList    []GuideCreateGuiderTargetRequestTargetList `json:"targetList,omitempty"`
	FiscalYear    int                                        `json:"fiscalYear,omitempty"`
}

type GuideCreateGuiderTargetResponse struct {
}

func CreateGuideCreateGuiderTargetRequest() (request *GuideCreateGuiderTargetRequest) {
	request = &GuideCreateGuiderTargetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/createGuiderTarget", "api")
	request.Method = api.POST
	return
}

func CreateGuideCreateGuiderTargetResponse() (response *api.BaseResponse[GuideCreateGuiderTargetResponse]) {
	response = api.CreateResponse[GuideCreateGuiderTargetResponse](&GuideCreateGuiderTargetResponse{})
	return
}

type GuideCreateGuiderTargetRequestTargetList struct {
	Details    []GuideCreateGuiderTargetRequestDetails `json:"details,omitempty"`
	TargetType int                                     `json:"targetType,omitempty"`
}

type GuideCreateGuiderTargetRequestDetails struct {
	CycleId     int     `json:"cycleId,omitempty"`
	CycleTarget float64 `json:"cycleTarget,omitempty"`
}
