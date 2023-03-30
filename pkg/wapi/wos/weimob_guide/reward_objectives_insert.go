package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RewardObjectivesInsert
// @id 1755
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1755?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增导购目标)
func (client *Client) RewardObjectivesInsert(request *RewardObjectivesInsertRequest) (response *RewardObjectivesInsertResponse, err error) {
	rpcResponse := CreateRewardObjectivesInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RewardObjectivesInsertRequest struct {
	*api.BaseRequest

	TargetList    []RewardObjectivesInsertRequestTargetList `json:"targetList,omitempty"`
	FiscalYear    int64                                     `json:"fiscalYear,omitempty"`
	GuiderWidList []int64                                   `json:"guiderWidList,omitempty"`
}

type RewardObjectivesInsertResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateRewardObjectivesInsertRequest() (request *RewardObjectivesInsertRequest) {
	request = &RewardObjectivesInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "reward/objectives/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateRewardObjectivesInsertResponse() (response *api.BaseResponse[RewardObjectivesInsertResponse]) {
	response = api.CreateResponse[RewardObjectivesInsertResponse](&RewardObjectivesInsertResponse{})
	return
}

type RewardObjectivesInsertRequestTargetList struct {
	Details    []RewardObjectivesInsertRequestDetails `json:"details,omitempty"`
	TargetType int64                                  `json:"targetType,omitempty"`
}

type RewardObjectivesInsertRequestDetails struct {
	CycleId     int64 `json:"cycleId,omitempty"`
	CycleTarget int64 `json:"cycleTarget,omitempty"`
}
