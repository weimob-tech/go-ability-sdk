package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdepartmentAddDepartment
// @id 1852
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加部门区域)
func (client *Client) MerchantdepartmentAddDepartment(request *MerchantdepartmentAddDepartmentRequest) (response *MerchantdepartmentAddDepartmentResponse, err error) {
	rpcResponse := CreateMerchantdepartmentAddDepartmentResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdepartmentAddDepartmentRequest struct {
	*api.BaseRequest

	DepartmentName   string `json:"departmentName,omitempty"`
	DepartmentNumber string `json:"departmentNumber,omitempty"`
	ParentId         int64  `json:"parentId,omitempty"`
}

type MerchantdepartmentAddDepartmentResponse struct {
}

func CreateMerchantdepartmentAddDepartmentRequest() (request *MerchantdepartmentAddDepartmentRequest) {
	request = &MerchantdepartmentAddDepartmentRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdepartment/addDepartment", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdepartmentAddDepartmentResponse() (response *api.BaseResponse[MerchantdepartmentAddDepartmentResponse]) {
	response = api.CreateResponse[MerchantdepartmentAddDepartmentResponse](&MerchantdepartmentAddDepartmentResponse{})
	return
}
