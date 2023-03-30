package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreTargetCreate
// @id 1945
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1945?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增门店目标)
func (client *Client) StoreTargetCreate(request *StoreTargetCreateRequest) (response *StoreTargetCreateResponse, err error) {
	rpcResponse := CreateStoreTargetCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreTargetCreateRequest struct {
	*api.BaseRequest

	TargetList    []StoreTargetCreateRequestTargetList `json:"targetList,omitempty"`
	FiscalYear    int64                                `json:"fiscalYear,omitempty"`
	TargetVidList []int64                              `json:"targetVidList,omitempty"`
}

type StoreTargetCreateResponse struct {
	FailedItemList []StoreTargetCreateResponseFailedItemList `json:"failedItemList,omitempty"`
	Total          int64                                     `json:"total,omitempty"`
	Success        int64                                     `json:"success,omitempty"`
	Fail           int64                                     `json:"fail,omitempty"`
}

func CreateStoreTargetCreateRequest() (request *StoreTargetCreateRequest) {
	request = &StoreTargetCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "store/target/create", "apigw")
	request.Method = api.POST
	return
}

func CreateStoreTargetCreateResponse() (response *api.BaseResponse[StoreTargetCreateResponse]) {
	response = api.CreateResponse[StoreTargetCreateResponse](&StoreTargetCreateResponse{})
	return
}

type StoreTargetCreateRequestTargetList struct {
	Details    []StoreTargetCreateRequestDetails `json:"details,omitempty"`
	TargetType int64                             `json:"targetType,omitempty"`
}

type StoreTargetCreateRequestDetails struct {
	CycleId          int64 `json:"cycleId,omitempty"`
	CycleTargetValue int64 `json:"cycleTargetValue,omitempty"`
}

type StoreTargetCreateResponseFailedItemList struct {
	Reason  string  `json:"reason,omitempty"`
	VidList []int64 `json:"vidList,omitempty"`
}
