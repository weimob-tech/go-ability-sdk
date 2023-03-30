package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrganizationMappingVidGetList
// @id 2125
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2125?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取映射后的节点vid)
func (client *Client) OrganizationMappingVidGetList(request *OrganizationMappingVidGetListRequest) (response *OrganizationMappingVidGetListResponse, err error) {
	rpcResponse := CreateOrganizationMappingVidGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrganizationMappingVidGetListRequest struct {
	*api.BaseRequest

	NodeInfoList []OrganizationMappingVidGetListRequestNodeInfoList `json:"nodeInfoList,omitempty"`
}

type OrganizationMappingVidGetListResponse struct {
	VidInfoList []OrganizationMappingVidGetListResponseVidInfoList `json:"vidInfoList,omitempty"`
}

func CreateOrganizationMappingVidGetListRequest() (request *OrganizationMappingVidGetListRequest) {
	request = &OrganizationMappingVidGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "organization/mapping/vid/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrganizationMappingVidGetListResponse() (response *api.BaseResponse[OrganizationMappingVidGetListResponse]) {
	response = api.CreateResponse[OrganizationMappingVidGetListResponse](&OrganizationMappingVidGetListResponse{})
	return
}

type OrganizationMappingVidGetListRequestNodeInfoList struct {
	NodeId   int64 `json:"nodeId,omitempty"`
	NodeType int64 `json:"nodeType,omitempty"`
}

type OrganizationMappingVidGetListResponseVidInfoList struct {
	NodeId   int64 `json:"nodeId,omitempty"`
	NodeType int64 `json:"nodeType,omitempty"`
	Vid      int64 `json:"vid,omitempty"`
	VidType  int64 `json:"vidType,omitempty"`
}
