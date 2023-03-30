package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdepartmentChangeStoreDepartment
// @id 1809
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改门店所属区域)
func (client *Client) MerchantdepartmentChangeStoreDepartment(request *MerchantdepartmentChangeStoreDepartmentRequest) (response *MerchantdepartmentChangeStoreDepartmentResponse, err error) {
	rpcResponse := CreateMerchantdepartmentChangeStoreDepartmentResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdepartmentChangeStoreDepartmentRequest struct {
	*api.BaseRequest

	StoreId      int64 `json:"storeId,omitempty"`
	DepartmentId int64 `json:"departmentId,omitempty"`
}

type MerchantdepartmentChangeStoreDepartmentResponse struct {
}

func CreateMerchantdepartmentChangeStoreDepartmentRequest() (request *MerchantdepartmentChangeStoreDepartmentRequest) {
	request = &MerchantdepartmentChangeStoreDepartmentRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdepartment/changeStoreDepartment", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdepartmentChangeStoreDepartmentResponse() (response *api.BaseResponse[MerchantdepartmentChangeStoreDepartmentResponse]) {
	response = api.CreateResponse[MerchantdepartmentChangeStoreDepartmentResponse](&MerchantdepartmentChangeStoreDepartmentResponse{})
	return
}
