package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinanceBillAccountGet
// @id 1737
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1737?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询账户余额)
func (client *Client) FinanceBillAccountGet(request *FinanceBillAccountGetRequest) (response *FinanceBillAccountGetResponse, err error) {
	rpcResponse := CreateFinanceBillAccountGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinanceBillAccountGetRequest struct {
	*api.BaseRequest

	Wids []int64 `json:"wids,omitempty"`
}

type FinanceBillAccountGetResponse struct {
	BalanceDetailVOS []FinanceBillAccountGetResponseBalanceDetailVOS `json:"balanceDetailVOS,omitempty"`
}

func CreateFinanceBillAccountGetRequest() (request *FinanceBillAccountGetRequest) {
	request = &FinanceBillAccountGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/bill/account/get", "apigw")
	request.Method = api.POST
	return
}

func CreateFinanceBillAccountGetResponse() (response *api.BaseResponse[FinanceBillAccountGetResponse]) {
	response = api.CreateResponse[FinanceBillAccountGetResponse](&FinanceBillAccountGetResponse{})
	return
}

type FinanceBillAccountGetResponseBalanceDetailVOS struct {
	Wid    string `json:"wid,omitempty"`
	Amount string `json:"amount,omitempty"`
}
