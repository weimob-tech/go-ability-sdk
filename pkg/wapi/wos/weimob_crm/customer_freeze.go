package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerFreeze
// @id 2039
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2039?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=冻结客户)
func (client *Client) CustomerFreeze(request *CustomerFreezeRequest) (response *CustomerFreezeResponse, err error) {
	rpcResponse := CreateCustomerFreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerFreezeRequest struct {
	*api.BaseRequest

	WidList []int64 `json:"widList,omitempty"`
}

type CustomerFreezeResponse struct {
	FailList []int64 `json:"failList,omitempty"`
}

func CreateCustomerFreezeRequest() (request *CustomerFreezeRequest) {
	request = &CustomerFreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/freeze", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerFreezeResponse() (response *api.BaseResponse[CustomerFreezeResponse]) {
	response = api.CreateResponse[CustomerFreezeResponse](&CustomerFreezeResponse{})
	return
}
