package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideCreateStoreTarget
// @id 1254
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增门店目标)
func (client *Client) GuideCreateStoreTarget(request *GuideCreateStoreTargetRequest) (response *GuideCreateStoreTargetResponse, err error) {
	rpcResponse := CreateGuideCreateStoreTargetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideCreateStoreTargetRequest struct {
	*api.BaseRequest

	StoreIdList []int64                                   `json:"storeIdList,omitempty"`
	TargetList  []GuideCreateStoreTargetRequestTargetList `json:"targetList,omitempty"`
	FiscalYear  int                                       `json:"fiscalYear,omitempty"`
}

type GuideCreateStoreTargetResponse struct {
}

func CreateGuideCreateStoreTargetRequest() (request *GuideCreateStoreTargetRequest) {
	request = &GuideCreateStoreTargetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/createStoreTarget", "api")
	request.Method = api.POST
	return
}

func CreateGuideCreateStoreTargetResponse() (response *api.BaseResponse[GuideCreateStoreTargetResponse]) {
	response = api.CreateResponse[GuideCreateStoreTargetResponse](&GuideCreateStoreTargetResponse{})
	return
}

type GuideCreateStoreTargetRequestTargetList struct {
	Details    []GuideCreateStoreTargetRequestDetails `json:"details,omitempty"`
	TargetType int                                    `json:"targetType,omitempty"`
}

type GuideCreateStoreTargetRequestDetails struct {
	CycleId     int     `json:"cycleId,omitempty"`
	CycleTarget float64 `json:"cycleTarget,omitempty"`
}
