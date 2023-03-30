package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardGetRechargeLogPageListAndTotal
// @id 217
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取充值明细)
func (client *Client) KLDMemberCardGetRechargeLogPageListAndTotal(request *KLDMemberCardGetRechargeLogPageListAndTotalRequest) (response *KLDMemberCardGetRechargeLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateKLDMemberCardGetRechargeLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardGetRechargeLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo    string  `json:"memberCardNo,omitempty"`
	Name            string  `json:"name,omitempty"`
	StoreId         int64   `json:"storeId,omitempty"`
	Phone           string  `json:"phone,omitempty"`
	AmountMin       float64 `json:"amountMin,omitempty"`
	AmountMax       float64 `json:"amountMax,omitempty"`
	Sort            int     `json:"sort,omitempty"`
	Begintime       int64   `json:"begintime,omitempty"`
	Endtime         int64   `json:"endtime,omitempty"`
	PageIndex       int     `json:"pageIndex,omitempty"`
	PageSize        int     `json:"pageSize,omitempty"`
	IsOnlyEffective bool    `json:"isOnlyEffective,omitempty"`
	Operator        string  `json:"operator,omitempty"`
}

type KLDMemberCardGetRechargeLogPageListAndTotalResponse struct {
}

func CreateKLDMemberCardGetRechargeLogPageListAndTotalRequest() (request *KLDMemberCardGetRechargeLogPageListAndTotalRequest) {
	request = &KLDMemberCardGetRechargeLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/GetRechargeLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardGetRechargeLogPageListAndTotalResponse() (response *api.BaseResponse[KLDMemberCardGetRechargeLogPageListAndTotalResponse]) {
	response = api.CreateResponse[KLDMemberCardGetRechargeLogPageListAndTotalResponse](&KLDMemberCardGetRechargeLogPageListAndTotalResponse{})
	return
}
