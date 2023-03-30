package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceAdjust
// @id 1665
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1665?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=调整余额)
func (client *Client) BalanceAdjust(request *BalanceAdjustRequest) (response *BalanceAdjustResponse, err error) {
	rpcResponse := CreateBalanceAdjustResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceAdjustRequest struct {
	*api.BaseRequest

	AdjustAmountInfoList []BalanceAdjustRequestAdjustAmountInfoList `json:"adjustAmountInfoList,omitempty"`
	SettleRule           []BalanceAdjustRequestSettleRule           `json:"settleRule,omitempty"`
	RequestId            string                                     `json:"requestId,omitempty"`
	RequestType          string                                     `json:"requestType,omitempty"`
	AdjustType           string                                     `json:"adjustType,omitempty"`
	ChangeReason         string                                     `json:"changeReason,omitempty"`
	OperatorWid          string                                     `json:"operatorWid,omitempty"`
	ChangCapitalAmount   string                                     `json:"changCapitalAmount,omitempty"`
	ChangeGrantsAmount   string                                     `json:"changeGrantsAmount,omitempty"`
	OperateStoreVid      int64                                      `json:"operateStoreVid,omitempty"`
	SourceVid            int64                                      `json:"sourceVid,omitempty"`
	ChangeType           string                                     `json:"changeType,omitempty"`
	OutTransNo           string                                     `json:"outTransNo,omitempty"`
	OutTransType         string                                     `json:"outTransType,omitempty"`
}

type BalanceAdjustResponse struct {
	List   []BalanceAdjustResponselist `json:"list,omitempty"`
	Result bool                        `json:"result,omitempty"`
}

func CreateBalanceAdjustRequest() (request *BalanceAdjustRequest) {
	request = &BalanceAdjustRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/adjust", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceAdjustResponse() (response *api.BaseResponse[BalanceAdjustResponse]) {
	response = api.CreateResponse[BalanceAdjustResponse](&BalanceAdjustResponse{})
	return
}

type BalanceAdjustRequestAdjustAmountInfoList struct {
	Vid           int64 `json:"vid,omitempty"`
	Wid           int64 `json:"wid,omitempty"`
	BalancePlanId int64 `json:"balancePlanId,omitempty"`
}

type BalanceAdjustRequestSettleRule struct {
	Vid  int64 `json:"vid,omitempty"`
	Rate int64 `json:"rate,omitempty"`
}

type BalanceAdjustResponselist struct {
	Status  int64  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	WidName string `json:"widName,omitempty"`
}
