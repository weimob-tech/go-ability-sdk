package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeDelete
// @id 2230
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2230?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除子账号)
func (client *Client) EmployeeDelete(request *EmployeeDeleteRequest) (response *EmployeeDeleteResponse, err error) {
	rpcResponse := CreateEmployeeDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeDeleteRequest struct {
	*api.BaseRequest

	Wid   int64  `json:"wid,omitempty"`
	Phone string `json:"phone,omitempty"`
	Zone  string `json:"zone,omitempty"`
}

type EmployeeDeleteResponse struct {
}

func CreateEmployeeDeleteRequest() (request *EmployeeDeleteRequest) {
	request = &EmployeeDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeDeleteResponse() (response *api.BaseResponse[EmployeeDeleteResponse]) {
	response = api.CreateResponse[EmployeeDeleteResponse](&EmployeeDeleteResponse{})
	return
}
