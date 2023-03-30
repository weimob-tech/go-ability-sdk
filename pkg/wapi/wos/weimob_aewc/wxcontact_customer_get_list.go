package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WxcontactCustomerGetList
// @id 2084
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2084?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取指定渠道活码的客户添加情况)
func (client *Client) WxcontactCustomerGetList(request *WxcontactCustomerGetListRequest) (response *WxcontactCustomerGetListResponse, err error) {
	rpcResponse := CreateWxcontactCustomerGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WxcontactCustomerGetListRequest struct {
	*api.BaseRequest

	ContactId int64 `json:"contactId,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
	PageNum   int64 `json:"pageNum,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
}

type WxcontactCustomerGetListResponse struct {
	Results    []WxcontactCustomerGetListResponseResults `json:"results,omitempty"`
	PageSize   int64                                     `json:"pageSize,omitempty"`
	TotalPage  int64                                     `json:"totalPage,omitempty"`
	PageNum    int64                                     `json:"pageNum,omitempty"`
	TotalCount int64                                     `json:"totalCount,omitempty"`
}

func CreateWxcontactCustomerGetListRequest() (request *WxcontactCustomerGetListRequest) {
	request = &WxcontactCustomerGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "wxcontact/customer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateWxcontactCustomerGetListResponse() (response *api.BaseResponse[WxcontactCustomerGetListResponse]) {
	response = api.CreateResponse[WxcontactCustomerGetListResponse](&WxcontactCustomerGetListResponse{})
	return
}

type WxcontactCustomerGetListResponseResults struct {
	UserBaseInfoVO   WxcontactCustomerGetListResponseUserBaseInfoVO `json:"userBaseInfoVO,omitempty"`
	AddTime          string                                         `json:"addTime,omitempty"`
	LoseStatus       int64                                          `json:"loseStatus,omitempty"`
	Sex              int64                                          `json:"sex,omitempty"`
	CustomerId       string                                         `json:"customerId,omitempty"`
	ExternalUserName string                                         `json:"externalUserName,omitempty"`
}

type WxcontactCustomerGetListResponseUserBaseInfoVO struct {
	UserName      string `json:"userName,omitempty"`
	DepartmentId  int64  `json:"departmentId,omitempty"`
	HaveBindGuide bool   `json:"haveBindGuide,omitempty"`
	UserId        string `json:"userId,omitempty"`
	OrgUserId     string `json:"orgUserId,omitempty"`
}
