package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RewardObjectivesGetList
// @id 1754
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1754?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购目标列表)
func (client *Client) RewardObjectivesGetList(request *RewardObjectivesGetListRequest) (response *RewardObjectivesGetListResponse, err error) {
	rpcResponse := CreateRewardObjectivesGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RewardObjectivesGetListRequest struct {
	*api.BaseRequest

	PageNum                 int64  `json:"pageNum,omitempty"`
	PageSize                int64  `json:"pageSize,omitempty"`
	FiscalYear              int64  `json:"fiscalYear,omitempty"`
	StoreId                 int64  `json:"storeId,omitempty"`
	GuiderWid               int64  `json:"guiderWid,omitempty"`
	GuiderPhone             string `json:"guiderPhone,omitempty"`
	JobNumber               string `json:"jobNumber,omitempty"`
	BosId                   int64  `json:"bosId,omitempty"`
	Vid                     int64  `json:"vid,omitempty"`
	VidType                 int64  `json:"vidType,omitempty"`
	TargetProductId         int64  `json:"targetProductId,omitempty"`
	TargetProductInstanceId int64  `json:"targetProductInstanceId,omitempty"`
}

type RewardObjectivesGetListResponse struct {
	PageList   []RewardObjectivesGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                     `json:"pageNum,omitempty"`
	PageSize   int64                                     `json:"pageSize,omitempty"`
	TotalCount int64                                     `json:"totalCount,omitempty"`
}

func CreateRewardObjectivesGetListRequest() (request *RewardObjectivesGetListRequest) {
	request = &RewardObjectivesGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "reward/objectives/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateRewardObjectivesGetListResponse() (response *api.BaseResponse[RewardObjectivesGetListResponse]) {
	response = api.CreateResponse[RewardObjectivesGetListResponse](&RewardObjectivesGetListResponse{})
	return
}

type RewardObjectivesGetListResponsePageList struct {
	TargetItemList []RewardObjectivesGetListResponseTargetItemList `json:"targetItemList,omitempty"`
	FiscalYear     int64                                           `json:"fiscalYear,omitempty"`
	GuiderPhone    string                                          `json:"guiderPhone,omitempty"`
	GuiderName     string                                          `json:"guiderName,omitempty"`
	JobNumber      string                                          `json:"jobNumber,omitempty"`
	GuiderWid      int64                                           `json:"guiderWid,omitempty"`
	StoreId        int64                                           `json:"storeId,omitempty"`
	StoreName      string                                          `json:"storeName,omitempty"`
	StoreNumber    string                                          `json:"storeNumber,omitempty"`
}

type RewardObjectivesGetListResponseTargetItemList struct {
	TargetDetails []RewardObjectivesGetListResponseTargetDetails `json:"targetDetails,omitempty"`
	TargetType    int64                                          `json:"targetType,omitempty"`
	YearTarget    int64                                          `json:"yearTarget,omitempty"`
}

type RewardObjectivesGetListResponseTargetDetails struct {
	CycleId     RewardObjectivesGetListResponseCycleId `json:"cycleId,omitempty"`
	CycleTarget int64                                  `json:"cycleTarget,omitempty"`
	CycleTitle  string                                 `json:"cycleTitle,omitempty"`
}

type RewardObjectivesGetListResponseCycleId struct {
}
