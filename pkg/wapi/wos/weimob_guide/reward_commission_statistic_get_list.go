package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RewardCommissionStatisticGetList
// @id 1739
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1739?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购提成统计)
func (client *Client) RewardCommissionStatisticGetList(request *RewardCommissionStatisticGetListRequest) (response *RewardCommissionStatisticGetListResponse, err error) {
	rpcResponse := CreateRewardCommissionStatisticGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RewardCommissionStatisticGetListRequest struct {
	*api.BaseRequest

	PageNum     int64  `json:"pageNum,omitempty"`
	PageSize    int64  `json:"pageSize,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	VidType     int64  `json:"vidType,omitempty"`
}

type RewardCommissionStatisticGetListResponse struct {
	PageList   []RewardCommissionStatisticGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                              `json:"pageNum,omitempty"`
	PageSize   int64                                              `json:"pageSize,omitempty"`
	TotalCount int64                                              `json:"totalCount,omitempty"`
}

func CreateRewardCommissionStatisticGetListRequest() (request *RewardCommissionStatisticGetListRequest) {
	request = &RewardCommissionStatisticGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "reward/commission/statistic/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateRewardCommissionStatisticGetListResponse() (response *api.BaseResponse[RewardCommissionStatisticGetListResponse]) {
	response = api.CreateResponse[RewardCommissionStatisticGetListResponse](&RewardCommissionStatisticGetListResponse{})
	return
}

type RewardCommissionStatisticGetListResponsePageList struct {
	FiscalYear       int64  `json:"fiscalYear,omitempty"`
	CycleName        string `json:"cycleName,omitempty"`
	GuideName        string `json:"guideName,omitempty"`
	GuidePhone       string `json:"guidePhone,omitempty"`
	PlanTitle        string `json:"planTitle,omitempty"`
	StoreName        string `json:"storeName,omitempty"`
	TotalCommission  int64  `json:"totalCommission,omitempty"`
	GoodsCommission  int64  `json:"goodsCommission,omitempty"`
	OrderCommission  int64  `json:"orderCommission,omitempty"`
	ChargeCommission int64  `json:"chargeCommission,omitempty"`
	FansCommission   int64  `json:"fansCommission,omitempty"`
	CardCommission   int64  `json:"cardCommission	,omitempty"`
}
