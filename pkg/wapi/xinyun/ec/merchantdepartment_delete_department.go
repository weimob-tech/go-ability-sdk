package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdepartmentDeleteDepartment
// @id 1850
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除部门区域)
func (client *Client) MerchantdepartmentDeleteDepartment(request *MerchantdepartmentDeleteDepartmentRequest) (response *MerchantdepartmentDeleteDepartmentResponse, err error) {
	rpcResponse := CreateMerchantdepartmentDeleteDepartmentResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdepartmentDeleteDepartmentRequest struct {
	*api.BaseRequest

	DepartmentId int64 `json:"departmentId,omitempty"`
}

type MerchantdepartmentDeleteDepartmentResponse struct {
}

func CreateMerchantdepartmentDeleteDepartmentRequest() (request *MerchantdepartmentDeleteDepartmentRequest) {
	request = &MerchantdepartmentDeleteDepartmentRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdepartment/deleteDepartment", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdepartmentDeleteDepartmentResponse() (response *api.BaseResponse[MerchantdepartmentDeleteDepartmentResponse]) {
	response = api.CreateResponse[MerchantdepartmentDeleteDepartmentResponse](&MerchantdepartmentDeleteDepartmentResponse{})
	return
}
