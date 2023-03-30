package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CampaignGetList
// @id 2070
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2070?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询有礼活动列表)
func (client *Client) CampaignGetList(request *CampaignGetListRequest) (response *CampaignGetListResponse, err error) {
	rpcResponse := CreateCampaignGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CampaignGetListRequest struct {
	*api.BaseRequest

	BizList                 []CampaignGetListRequestBizList `json:"bizList,omitempty"`
	BosId                   int64                           `json:"bosId,omitempty"`
	TargetProductInstanceId int64                           `json:"targetProductInstanceId,omitempty"`
	TargetProductId         int64                           `json:"targetProductId,omitempty"`
	ActivityStatus          int64                           `json:"activityStatus,omitempty"`
	ActivityType            string                          `json:"activityType,omitempty"`
	PageNum                 int64                           `json:"pageNum,omitempty"`
	PageSize                int64                           `json:"pageSize,omitempty"`
}

type CampaignGetListResponse struct {
	Items      []CampaignGetListResponseItems `json:"items,omitempty"`
	PageNum    int64                          `json:"pageNum,omitempty"`
	PageSize   int64                          `json:"pageSize,omitempty"`
	TotalCount int64                          `json:"totalCount,omitempty"`
}

func CreateCampaignGetListRequest() (request *CampaignGetListRequest) {
	request = &CampaignGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "campaign/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCampaignGetListResponse() (response *api.BaseResponse[CampaignGetListResponse]) {
	response = api.CreateResponse[CampaignGetListResponse](&CampaignGetListResponse{})
	return
}

type CampaignGetListRequestBizList struct {
	BizKey   string `json:"bizKey,omitempty"`
	BizValue string `json:"bizValue,omitempty"`
}

type CampaignGetListResponseItems struct {
	RewardRankList      []CampaignGetListResponseRewardRankList `json:"rewardRankList,omitempty"`
	RewardTimeRule      CampaignGetListResponseRewardTimeRule   `json:"rewardTimeRule,omitempty"`
	ActivityName        string                                  `json:"activityName,omitempty"`
	TimeType            int64                                   `json:"timeType,omitempty"`
	StartTime           int64                                   `json:"startTime,omitempty"`
	EndTime             int64                                   `json:"endTime,omitempty"`
	RewardConditionType int64                                   `json:"rewardConditionType,omitempty"`
	CrowdRuleId         int64                                   `json:"crowdRuleId,omitempty"`
}

type CampaignGetListResponseRewardRankList struct {
	RewardInfoList       []CampaignGetListResponseRewardInfoList `json:"rewardInfoList,omitempty"`
	RewardConditionValue string                                  `json:"rewardConditionValue,omitempty"`
}

type CampaignGetListResponseRewardInfoList struct {
	RewardKey      int64  `json:"rewardKey,omitempty"`
	RewardUnit     string `json:"rewardUnit,omitempty"`
	RewardMaxNum   int64  `json:"rewardMaxNum,omitempty"`
	RewardNum      int64  `json:"rewardNum,omitempty"`
	RewardMultiple string `json:"rewardMultiple,omitempty"`
}

type CampaignGetListResponseRewardTimeRule struct {
	RuleType  int64 `json:"ruleType,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
}
