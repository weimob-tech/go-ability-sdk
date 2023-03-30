package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillVirtualLogisticsFulfill
// @id 2234
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2234?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=虚拟商品发货)
func (client *Client) FulfillVirtualLogisticsFulfill(request *FulfillVirtualLogisticsFulfillRequest) (response *FulfillVirtualLogisticsFulfillResponse, err error) {
	rpcResponse := CreateFulfillVirtualLogisticsFulfillResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillVirtualLogisticsFulfillRequest struct {
	*api.BaseRequest

	OperatorVo FulfillVirtualLogisticsFulfillRequestOperatorVo `json:"operatorVo,omitempty"`
	BasicInfo  FulfillVirtualLogisticsFulfillRequestBasicInfo  `json:"basicInfo,omitempty"`
	OrderNo    int64                                           `json:"orderNo,omitempty"`
}

type FulfillVirtualLogisticsFulfillResponse struct {
	Success bool `json:"success	,omitempty"`
}

func CreateFulfillVirtualLogisticsFulfillRequest() (request *FulfillVirtualLogisticsFulfillRequest) {
	request = &FulfillVirtualLogisticsFulfillRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/virtual/logistics/fulfill", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillVirtualLogisticsFulfillResponse() (response *api.BaseResponse[FulfillVirtualLogisticsFulfillResponse]) {
	response = api.CreateResponse[FulfillVirtualLogisticsFulfillResponse](&FulfillVirtualLogisticsFulfillResponse{})
	return
}

type FulfillVirtualLogisticsFulfillRequestOperatorVo struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type FulfillVirtualLogisticsFulfillRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
