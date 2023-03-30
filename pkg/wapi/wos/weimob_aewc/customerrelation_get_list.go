package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerrelationGetList
// @id 2280
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2280?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户与企微员工好友关系)
func (client *Client) CustomerrelationGetList(request *CustomerrelationGetListRequest) (response *CustomerrelationGetListResponse, err error) {
	rpcResponse := CreateCustomerrelationGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerrelationGetListRequest struct {
	*api.BaseRequest

	UserId string `json:"userId,omitempty"`
	Wid    int64  `json:"wid,omitempty"`
	Status int64  `json:"status,omitempty"`
}

type CustomerrelationGetListResponse struct {
	CorpId         string `json:"corpId,omitempty"`
	EmployeeId     string `json:"employeeId,omitempty"`
	CustomerId     string `json:"customerId,omitempty"`
	ExternalUserId string `json:"externalUserId,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
	AddWay         int64  `json:"addWay,omitempty"`
	Friended       int64  `json:"friended,omitempty"`
	AddTime        string `json:"addTime,omitempty"`
	Status         int64  `json:"status,omitempty"`
}

func CreateCustomerrelationGetListRequest() (request *CustomerrelationGetListRequest) {
	request = &CustomerrelationGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "customerrelation/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerrelationGetListResponse() (response *api.BaseResponse[CustomerrelationGetListResponse]) {
	response = api.CreateResponse[CustomerrelationGetListResponse](&CustomerrelationGetListResponse{})
	return
}
