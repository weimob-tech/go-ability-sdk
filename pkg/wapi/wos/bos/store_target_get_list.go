package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreTargetGetList
// @id 1944
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1944?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询门店目标列表)
func (client *Client) StoreTargetGetList(request *StoreTargetGetListRequest) (response *StoreTargetGetListResponse, err error) {
	rpcResponse := CreateStoreTargetGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreTargetGetListRequest struct {
	*api.BaseRequest

	TargetVid  int64 `json:"targetVid,omitempty"`
	FiscalYear int64 `json:"fiscalYear,omitempty"`
	PageSize   int64 `json:"pageSize,omitempty"`
	PageNum    int64 `json:"pageNum,omitempty"`
}

type StoreTargetGetListResponse struct {
	PageList   []StoreTargetGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                `json:"pageNum,omitempty"`
	PageSize   int64                                `json:"pageSize,omitempty"`
	TotalCount int64                                `json:"totalCount,omitempty"`
}

func CreateStoreTargetGetListRequest() (request *StoreTargetGetListRequest) {
	request = &StoreTargetGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "store/target/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateStoreTargetGetListResponse() (response *api.BaseResponse[StoreTargetGetListResponse]) {
	response = api.CreateResponse[StoreTargetGetListResponse](&StoreTargetGetListResponse{})
	return
}

type StoreTargetGetListResponsePageList struct {
	CycleSalesTargetList       []StoreTargetGetListResponseCycleSalesTargetList       `json:"cycleSalesTargetList,omitempty"`
	CycleChargeTargetList      []StoreTargetGetListResponseCycleChargeTargetList      `json:"cycleChargeTargetList,omitempty"`
	CycleOpenCardTargetList    []StoreTargetGetListResponseCycleOpenCardTargetList    `json:"cycleOpenCardTargetList,omitempty"`
	CycleAttractFansTargetList []StoreTargetGetListResponseCycleAttractFansTargetList `json:"cycleAttractFansTargetList,omitempty"`
	FiscalYear                 int64                                                  `json:"fiscalYear,omitempty"`
	StoreStatus                int64                                                  `json:"storeStatus,omitempty"`
	Vid                        int64                                                  `json:"vid,omitempty"`
	StoreName                  string                                                 `json:"storeName,omitempty"`
	StoreNumber                string                                                 `json:"storeNumber,omitempty"`
	YearSalesTarget            int64                                                  `json:"yearSalesTarget,omitempty"`
	ChargeTarget               int64                                                  `json:"chargeTarget,omitempty"`
	OpenCardTarget             int64                                                  `json:"openCardTarget,omitempty"`
	AttractFansTarget          int64                                                  `json:"attractFansTarget,omitempty"`
	IsUsed                     bool                                                   `json:"isUsed,omitempty"`
}

type StoreTargetGetListResponseCycleSalesTargetList struct {
	CycleId          int64  `json:"cycleId,omitempty"`
	CycleTitle       string `json:"cycleTitle,omitempty"`
	CycleTargetValue int64  `json:"cycleTargetValue,omitempty"`
}

type StoreTargetGetListResponseCycleChargeTargetList struct {
	CycleId          int64  `json:"cycleId,omitempty"`
	CycleTitle       string `json:"cycleTitle,omitempty"`
	CycleTargetValue int64  `json:"cycleTargetValue,omitempty"`
}

type StoreTargetGetListResponseCycleOpenCardTargetList struct {
	CycleId          int64  `json:"cycleId,omitempty"`
	CycleTitle       string `json:"cycleTitle,omitempty"`
	CycleTargetValue int64  `json:"cycleTargetValue,omitempty"`
}

type StoreTargetGetListResponseCycleAttractFansTargetList struct {
	CycleId          int64  `json:"cycleId,omitempty"`
	CycleTitle       string `json:"cycleTitle,omitempty"`
	CycleTargetValue int64  `json:"cycleTargetValue,omitempty"`
}
