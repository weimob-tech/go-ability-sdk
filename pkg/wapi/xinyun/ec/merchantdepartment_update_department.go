package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdepartmentUpdateDepartment
// @id 1851
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改部门区域)
func (client *Client) MerchantdepartmentUpdateDepartment(request *MerchantdepartmentUpdateDepartmentRequest) (response *MerchantdepartmentUpdateDepartmentResponse, err error) {
	rpcResponse := CreateMerchantdepartmentUpdateDepartmentResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdepartmentUpdateDepartmentRequest struct {
	*api.BaseRequest

	DepartmentId     int64  `json:"departmentId,omitempty"`
	DepartmentName   string `json:"departmentName,omitempty"`
	DepartmentNumber string `json:"departmentNumber,omitempty"`
}

type MerchantdepartmentUpdateDepartmentResponse struct {
}

func CreateMerchantdepartmentUpdateDepartmentRequest() (request *MerchantdepartmentUpdateDepartmentRequest) {
	request = &MerchantdepartmentUpdateDepartmentRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdepartment/updateDepartment", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdepartmentUpdateDepartmentResponse() (response *api.BaseResponse[MerchantdepartmentUpdateDepartmentResponse]) {
	response = api.CreateResponse[MerchantdepartmentUpdateDepartmentResponse](&MerchantdepartmentUpdateDepartmentResponse{})
	return
}
