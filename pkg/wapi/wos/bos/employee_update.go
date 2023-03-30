package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeUpdate
// @id 2124
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2124?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新员工基本信息)
func (client *Client) EmployeeUpdate(request *EmployeeUpdateRequest) (response *EmployeeUpdateResponse, err error) {
	rpcResponse := CreateEmployeeUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeUpdateRequest struct {
	*api.BaseRequest

	Wid        int64  `json:"wid,omitempty"`
	Name       string `json:"name,omitempty"`
	EmployeeNo string `json:"employeeNo,omitempty"`
}

type EmployeeUpdateResponse struct {
}

func CreateEmployeeUpdateRequest() (request *EmployeeUpdateRequest) {
	request = &EmployeeUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/update", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeUpdateResponse() (response *api.BaseResponse[EmployeeUpdateResponse]) {
	response = api.CreateResponse[EmployeeUpdateResponse](&EmployeeUpdateResponse{})
	return
}
