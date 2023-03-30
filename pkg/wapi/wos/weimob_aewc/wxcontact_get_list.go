package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WxcontactGetList
// @id 2085
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2085?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取渠道活码列表)
func (client *Client) WxcontactGetList(request *WxcontactGetListRequest) (response *WxcontactGetListResponse, err error) {
	rpcResponse := CreateWxcontactGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WxcontactGetListRequest struct {
	*api.BaseRequest

	State       int64  `json:"state,omitempty"`
	ContactName string `json:"contactName,omitempty"`
	StartTime   int64  `json:"startTime,omitempty"`
	EndTime     int64  `json:"endTime,omitempty"`
	PageNum     int64  `json:"pageNum,omitempty"`
	PageSize    int64  `json:"pageSize,omitempty"`
}

type WxcontactGetListResponse struct {
	Results    []WxcontactGetListResponseResults `json:"results,omitempty"`
	TotalPage  int64                             `json:"totalPage,omitempty"`
	PageSize   int64                             `json:"pageSize,omitempty"`
	TotalCount int64                             `json:"totalCount,omitempty"`
	PageNum    int64                             `json:"pageNum,omitempty"`
}

func CreateWxcontactGetListRequest() (request *WxcontactGetListRequest) {
	request = &WxcontactGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "wxcontact/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateWxcontactGetListResponse() (response *api.BaseResponse[WxcontactGetListResponse]) {
	response = api.CreateResponse[WxcontactGetListResponse](&WxcontactGetListResponse{})
	return
}

type WxcontactGetListResponseResults struct {
	UserBaseInfoVOS []WxcontactGetListResponseUserBaseInfoVOS `json:"userBaseInfoVOS,omitempty"`
	ContactName     string                                    `json:"contactName,omitempty"`
	ContactId       int64                                     `json:"contactId,omitempty"`
	State           int64                                     `json:"state,omitempty"`
	LiveCodeAddSum  int64                                     `json:"liveCodeAddSum,omitempty"`
	QrCodeUrl       string                                    `json:"qrCodeUrl,omitempty"`
	CreateTime      string                                    `json:"createTime,omitempty"`
	ConfigId        string                                    `json:"configId,omitempty"`
}

type WxcontactGetListResponseUserBaseInfoVOS struct {
	OrgUserId     string `json:"orgUserId,omitempty"`
	UserId        string `json:"userId,omitempty"`
	UserName      string `json:"userName,omitempty"`
	HaveBindGuide bool   `json:"haveBindGuide,omitempty"`
	DepartmentId  int64  `json:"departmentId,omitempty"`
}
