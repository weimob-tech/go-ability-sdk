package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceGetList
// @id 1669
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1669?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户余额列表)
func (client *Client) BalanceGetList(request *BalanceGetListRequest) (response *BalanceGetListResponse, err error) {
	rpcResponse := CreateBalanceGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceGetListRequest struct {
	*api.BaseRequest

	WidList       []int64 `json:"widList,omitempty"`
	Vid           int64   `json:"vid,omitempty"`
	OperatorWid   int64   `json:"operatorWid,omitempty"`
	UserType      int64   `json:"userType,omitempty"`
	ChannelType   int64   `json:"channelType,omitempty"`
	BalancePlanId int64   `json:"balancePlanId,omitempty"`
	PageNum       int64   `json:"pageNum,omitempty"`
	PageSize      int64   `json:"pageSize,omitempty"`
	Wid           int64   `json:"wid,omitempty"`
}

type BalanceGetListResponse struct {
	List []BalanceGetListResponselist `json:"list,omitempty"`
}

func CreateBalanceGetListRequest() (request *BalanceGetListRequest) {
	request = &BalanceGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceGetListResponse() (response *api.BaseResponse[BalanceGetListResponse]) {
	response = api.CreateResponse[BalanceGetListResponse](&BalanceGetListResponse{})
	return
}

type BalanceGetListResponselist struct {
	Wid                int64  `json:"wid,omitempty"`
	TotalCapitalAmount string `json:"totalCapitalAmount,omitempty"`
	RechargeTimes      int64  `json:"rechargeTimes,omitempty"`
}
