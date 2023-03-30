package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdepartmentQueryDepartmentListByDepId
// @id 1811
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询指定节点下级所有区域类型的部门)
func (client *Client) MerchantdepartmentQueryDepartmentListByDepId(request *MerchantdepartmentQueryDepartmentListByDepIdRequest) (response *MerchantdepartmentQueryDepartmentListByDepIdResponse, err error) {
	rpcResponse := CreateMerchantdepartmentQueryDepartmentListByDepIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdepartmentQueryDepartmentListByDepIdRequest struct {
	*api.BaseRequest

	DepartmentId int64 `json:"departmentId,omitempty"`
}

type MerchantdepartmentQueryDepartmentListByDepIdResponse struct {
	DepartmentList []MerchantdepartmentQueryDepartmentListByDepIdResponseDepartmentList `json:"departmentList,omitempty"`
}

func CreateMerchantdepartmentQueryDepartmentListByDepIdRequest() (request *MerchantdepartmentQueryDepartmentListByDepIdRequest) {
	request = &MerchantdepartmentQueryDepartmentListByDepIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdepartment/queryDepartmentListByDepId", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdepartmentQueryDepartmentListByDepIdResponse() (response *api.BaseResponse[MerchantdepartmentQueryDepartmentListByDepIdResponse]) {
	response = api.CreateResponse[MerchantdepartmentQueryDepartmentListByDepIdResponse](&MerchantdepartmentQueryDepartmentListByDepIdResponse{})
	return
}

type MerchantdepartmentQueryDepartmentListByDepIdResponseDepartmentList struct {
	DepartmentId     int64  `json:"departmentId,omitempty"`
	DepartmentName   string `json:"departmentName,omitempty"`
	ParentId         int64  `json:"parentId,omitempty"`
	DepartmentNumber string `json:"departmentNumber,omitempty"`
}
