package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeStatusUpdate
// @id 2127
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2127?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新员工状态)
func (client *Client) EmployeeStatusUpdate(request *EmployeeStatusUpdateRequest) (response *EmployeeStatusUpdateResponse, err error) {
	rpcResponse := CreateEmployeeStatusUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeStatusUpdateRequest struct {
	*api.BaseRequest

	Wid    int64 `json:"wid,omitempty"`
	Status int64 `json:"status,omitempty"`
}

type EmployeeStatusUpdateResponse struct {
}

func CreateEmployeeStatusUpdateRequest() (request *EmployeeStatusUpdateRequest) {
	request = &EmployeeStatusUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/status/update", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeStatusUpdateResponse() (response *api.BaseResponse[EmployeeStatusUpdateResponse]) {
	response = api.CreateResponse[EmployeeStatusUpdateResponse](&EmployeeStatusUpdateResponse{})
	return
}
