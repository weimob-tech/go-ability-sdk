package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerOrganizationBind
// @id 1502
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1502?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=绑定客户归属门店)
func (client *Client) CustomerOrganizationBind(request *CustomerOrganizationBindRequest) (response *CustomerOrganizationBindResponse, err error) {
	rpcResponse := CreateCustomerOrganizationBindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerOrganizationBindRequest struct {
	*api.BaseRequest

	OperateWid int64   `json:"operateWid,omitempty"`
	WidList    []int64 `json:"widList,omitempty"`
	Scene      int64   `json:"scene,omitempty"`
	Origin     int64   `json:"origin,omitempty"`
	Vid        int64   `json:"vid,omitempty"`
	Reason     string  `json:"reason,omitempty"`
}

type CustomerOrganizationBindResponse struct {
	SuccessWidList []int64 `json:"successWidList,omitempty"`
}

func CreateCustomerOrganizationBindRequest() (request *CustomerOrganizationBindRequest) {
	request = &CustomerOrganizationBindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/organization/bind", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerOrganizationBindResponse() (response *api.BaseResponse[CustomerOrganizationBindResponse]) {
	response = api.CreateResponse[CustomerOrganizationBindResponse](&CustomerOrganizationBindResponse{})
	return
}
