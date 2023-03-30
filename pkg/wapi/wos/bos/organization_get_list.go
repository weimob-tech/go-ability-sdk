package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrganizationGetList
// @id 1367
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1367?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询bos下的组织列表)
func (client *Client) OrganizationGetList(request *OrganizationGetListRequest) (response *OrganizationGetListResponse, err error) {
	rpcResponse := CreateOrganizationGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrganizationGetListRequest struct {
	*api.BaseRequest

	ParentVid int64  `json:"parentVid,omitempty"`
	ChildVid  int64  `json:"childVid,omitempty"`
	IsDirect  int64  `json:"isDirect,omitempty"`
	VidType   int64  `json:"vidType,omitempty"`
	VidStatus int64  `json:"vidStatus,omitempty"`
	PageSize  int64  `json:"pageSize,omitempty"`
	PageNum   int64  `json:"pageNum,omitempty"`
	VidName   string `json:"vidName,omitempty"`
}

type OrganizationGetListResponse struct {
	Data      []OrganizationGetListResponseData `json:"data,omitempty"`
	PageSize  int64                             `json:"pageSize,omitempty"`
	PageNum   int64                             `json:"pageNum,omitempty"`
	TotalSize int64                             `json:"totalSize,omitempty"`
}

func CreateOrganizationGetListRequest() (request *OrganizationGetListRequest) {
	request = &OrganizationGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "organization/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrganizationGetListResponse() (response *api.BaseResponse[OrganizationGetListResponse]) {
	response = api.CreateResponse[OrganizationGetListResponse](&OrganizationGetListResponse{})
	return
}

type OrganizationGetListResponseData struct {
	VidName   string `json:"vidName,omitempty"`
	Vid       int64  `json:"vid,omitempty"`
	ParentVid int64  `json:"parentVid,omitempty"`
	VidCode   string `json:"vidCode,omitempty"`
	VidType   int64  `json:"vidType,omitempty"`
	VidStatus int64  `json:"vidStatus,omitempty"`
}
