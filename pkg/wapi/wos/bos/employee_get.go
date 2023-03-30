package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeGet
// @id 2080
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2080?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询员工的基本信息)
func (client *Client) EmployeeGet(request *EmployeeGetRequest) (response *EmployeeGetResponse, err error) {
	rpcResponse := CreateEmployeeGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeGetRequest struct {
	*api.BaseRequest

	Wid        int64  `json:"wid,omitempty"`
	Phone      string `json:"phone,omitempty"`
	EmployeeNo string `json:"employeeNo,omitempty"`
}

type EmployeeGetResponse struct {
	Phone      string `json:"phone,omitempty"`
	Name       string `json:"name,omitempty"`
	EmployeeNo string `json:"employeeNo,omitempty"`
	Status     int64  `json:"status,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}

func CreateEmployeeGetRequest() (request *EmployeeGetRequest) {
	request = &EmployeeGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/get", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeGetResponse() (response *api.BaseResponse[EmployeeGetResponse]) {
	response = api.CreateResponse[EmployeeGetResponse](&EmployeeGetResponse{})
	return
}
