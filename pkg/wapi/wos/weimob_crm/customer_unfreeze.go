package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerUnfreeze
// @id 2040
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2040?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解冻客户)
func (client *Client) CustomerUnfreeze(request *CustomerUnfreezeRequest) (response *CustomerUnfreezeResponse, err error) {
	rpcResponse := CreateCustomerUnfreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerUnfreezeRequest struct {
	*api.BaseRequest

	WidList []int64 `json:"widList,omitempty"`
}

type CustomerUnfreezeResponse struct {
	FailList []int64 `json:"failList,omitempty"`
}

func CreateCustomerUnfreezeRequest() (request *CustomerUnfreezeRequest) {
	request = &CustomerUnfreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/unfreeze", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerUnfreezeResponse() (response *api.BaseResponse[CustomerUnfreezeResponse]) {
	response = api.CreateResponse[CustomerUnfreezeResponse](&CustomerUnfreezeResponse{})
	return
}
