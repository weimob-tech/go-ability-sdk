package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsApplyRights
// @id 1423
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=申请售后)
func (client *Client) RightsApplyRights(request *RightsApplyRightsRequest) (response *RightsApplyRightsResponse, err error) {
	rpcResponse := CreateRightsApplyRightsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsApplyRightsRequest struct {
	*api.BaseRequest

	ApplyItemList []RightsApplyRightsRequestApplyItemList `json:"applyItemList,omitempty"`
	OrderNo       int64                                   `json:"orderNo,omitempty"`
}

type RightsApplyRightsResponse struct {
}

func CreateRightsApplyRightsRequest() (request *RightsApplyRightsRequest) {
	request = &RightsApplyRightsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "rights/applyRights", "api")
	request.Method = api.POST
	return
}

func CreateRightsApplyRightsResponse() (response *api.BaseResponse[RightsApplyRightsResponse]) {
	response = api.CreateResponse[RightsApplyRightsResponse](&RightsApplyRightsResponse{})
	return
}

type RightsApplyRightsRequestApplyItemList struct {
	ItemId             int64   `json:"itemId,omitempty"`
	ApplyAmount        float64 `json:"applyAmount,omitempty"`
	RightsType         int     `json:"rightsType,omitempty"`
	RefundMethod       int     `json:"refundMethod,omitempty"`
	RightsReason       string  `json:"rightsReason,omitempty"`
	CustomRightsReason string  `json:"customRightsReason,omitempty"`
	StoreId            int64   `json:"storeId,omitempty"`
}
