package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerOrganizationUnbind
// @id 1740
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1740?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解绑客户归属门店)
func (client *Client) CustomerOrganizationUnbind(request *CustomerOrganizationUnbindRequest) (response *CustomerOrganizationUnbindResponse, err error) {
	rpcResponse := CreateCustomerOrganizationUnbindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerOrganizationUnbindRequest struct {
	*api.BaseRequest

	Vid     int64   `json:"vid,omitempty"`
	WidList []int64 `json:"widList,omitempty"`
	Scene   int64   `json:"scene,omitempty"`
	Origin  int64   `json:"origin,omitempty"`
	Reason  string  `json:"reason,omitempty"`
}

type CustomerOrganizationUnbindResponse struct {
	SuccessWidList []int64 `json:"successWidList,omitempty"`
}

func CreateCustomerOrganizationUnbindRequest() (request *CustomerOrganizationUnbindRequest) {
	request = &CustomerOrganizationUnbindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/organization/unbind", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerOrganizationUnbindResponse() (response *api.BaseResponse[CustomerOrganizationUnbindResponse]) {
	response = api.CreateResponse[CustomerOrganizationUnbindResponse](&CustomerOrganizationUnbindResponse{})
	return
}
