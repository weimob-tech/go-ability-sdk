package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdepartmentQueryDepartmentRootNode
// @id 1838
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询当前pid指定的根节点对应部门id)
func (client *Client) MerchantdepartmentQueryDepartmentRootNode(request *MerchantdepartmentQueryDepartmentRootNodeRequest) (response *MerchantdepartmentQueryDepartmentRootNodeResponse, err error) {
	rpcResponse := CreateMerchantdepartmentQueryDepartmentRootNodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdepartmentQueryDepartmentRootNodeRequest struct {
	*api.BaseRequest
}

type MerchantdepartmentQueryDepartmentRootNodeResponse struct {
	DepartmentId   int64  `json:"departmentId,omitempty"`
	DepartmentName string `json:"departmentName,omitempty"`
}

func CreateMerchantdepartmentQueryDepartmentRootNodeRequest() (request *MerchantdepartmentQueryDepartmentRootNodeRequest) {
	request = &MerchantdepartmentQueryDepartmentRootNodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdepartment/queryDepartmentRootNode", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdepartmentQueryDepartmentRootNodeResponse() (response *api.BaseResponse[MerchantdepartmentQueryDepartmentRootNodeResponse]) {
	response = api.CreateResponse[MerchantdepartmentQueryDepartmentRootNodeResponse](&MerchantdepartmentQueryDepartmentRootNodeResponse{})
	return
}
