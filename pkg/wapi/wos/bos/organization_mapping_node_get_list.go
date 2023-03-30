package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrganizationMappingNodeGetList
// @id 2121
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2121?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据节点vid反查询映射源)
func (client *Client) OrganizationMappingNodeGetList(request *OrganizationMappingNodeGetListRequest) (response *OrganizationMappingNodeGetListResponse, err error) {
	rpcResponse := CreateOrganizationMappingNodeGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrganizationMappingNodeGetListRequest struct {
	*api.BaseRequest

	VidList []int64 `json:"vidList,omitempty"`
}

type OrganizationMappingNodeGetListResponse struct {
	NodeInfoList []OrganizationMappingNodeGetListResponseNodeInfoList `json:"nodeInfoList,omitempty"`
}

func CreateOrganizationMappingNodeGetListRequest() (request *OrganizationMappingNodeGetListRequest) {
	request = &OrganizationMappingNodeGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "organization/mapping/node/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrganizationMappingNodeGetListResponse() (response *api.BaseResponse[OrganizationMappingNodeGetListResponse]) {
	response = api.CreateResponse[OrganizationMappingNodeGetListResponse](&OrganizationMappingNodeGetListResponse{})
	return
}

type OrganizationMappingNodeGetListResponseNodeInfoList struct {
	Vid         int64 `json:"vid,omitempty"`
	NodeId      int64 `json:"nodeId,omitempty"`
	NodeType    int64 `json:"nodeType,omitempty"`
	IsMainStore int64 `json:"isMainStore,omitempty"`
}
