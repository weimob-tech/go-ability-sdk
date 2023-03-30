package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeCreate
// @id 2079
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2079?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加员工)
func (client *Client) EmployeeCreate(request *EmployeeCreateRequest) (response *EmployeeCreateResponse, err error) {
	rpcResponse := CreateEmployeeCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeCreateRequest struct {
	*api.BaseRequest

	RegisterType int64  `json:"registerType,omitempty"`
	Name         string `json:"name,omitempty"`
	EmployeeNo   string `json:"employeeNo,omitempty"`
	RegisterId   string `json:"registerId,omitempty"`
	Zone         string `json:"zone,omitempty"`
}

type EmployeeCreateResponse struct {
	Wid int64 `json:"wid,omitempty"`
}

func CreateEmployeeCreateRequest() (request *EmployeeCreateRequest) {
	request = &EmployeeCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/create", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeCreateResponse() (response *api.BaseResponse[EmployeeCreateResponse]) {
	response = api.CreateResponse[EmployeeCreateResponse](&EmployeeCreateResponse{})
	return
}
