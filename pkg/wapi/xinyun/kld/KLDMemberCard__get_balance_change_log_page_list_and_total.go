package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardGetBalanceChangeLogPageListAndTotal
// @id 215
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取余额流水)
func (client *Client) KLDMemberCardGetBalanceChangeLogPageListAndTotal(request *KLDMemberCardGetBalanceChangeLogPageListAndTotalRequest) (response *KLDMemberCardGetBalanceChangeLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateKLDMemberCardGetBalanceChangeLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardGetBalanceChangeLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo    string `json:"memberCardNo,omitempty"`
	Sort            int    `json:"sort,omitempty"`
	BeginTime       int64  `json:"beginTime,omitempty"`
	EndTime         int64  `json:"endTime,omitempty"`
	PageIndex       int    `json:"pageIndex,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	IsOnlyEffective bool   `json:"isOnlyEffective,omitempty"`
}

type KLDMemberCardGetBalanceChangeLogPageListAndTotalResponse struct {
}

func CreateKLDMemberCardGetBalanceChangeLogPageListAndTotalRequest() (request *KLDMemberCardGetBalanceChangeLogPageListAndTotalRequest) {
	request = &KLDMemberCardGetBalanceChangeLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/GetBalanceChangeLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardGetBalanceChangeLogPageListAndTotalResponse() (response *api.BaseResponse[KLDMemberCardGetBalanceChangeLogPageListAndTotalResponse]) {
	response = api.CreateResponse[KLDMemberCardGetBalanceChangeLogPageListAndTotalResponse](&KLDMemberCardGetBalanceChangeLogPageListAndTotalResponse{})
	return
}
