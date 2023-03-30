package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrganizationDetailGetList
// @id 2086
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2086?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询组织基本信息)
func (client *Client) OrganizationDetailGetList(request *OrganizationDetailGetListRequest) (response *OrganizationDetailGetListResponse, err error) {
	rpcResponse := CreateOrganizationDetailGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrganizationDetailGetListRequest struct {
	*api.BaseRequest

	VidList     []int64 `json:"vidList,omitempty"`
	VidCodeList []int64 `json:"vidCodeList,omitempty"`
	Fields      []int64 `json:"fields,omitempty"`
}

type OrganizationDetailGetListResponse struct {
	ContactTels   []OrganizationDetailGetListResponseContactTels   `json:"contactTels,omitempty"`
	BusinessHours []OrganizationDetailGetListResponseBusinessHours `json:"businessHours,omitempty"`
	VidName       string                                           `json:"vidName,omitempty"`
	Vid           int64                                            `json:"vid,omitempty"`
	ParentVid     int64                                            `json:"parentVid,omitempty"`
	VidCode       string                                           `json:"vidCode,omitempty"`
	VidType       int64                                            `json:"vidType,omitempty"`
	VidStatus     int64                                            `json:"vidStatus,omitempty"`
	Logo          string                                           `json:"logo,omitempty"`
	ProvinceName  string                                           `json:"provinceName,omitempty"`
	CityName      string                                           `json:"cityName,omitempty"`
	CountyName    string                                           `json:"countyName,omitempty"`
	DetailAddress string                                           `json:"detailAddress,omitempty"`
	OnlineStatus  int64                                            `json:"onlineStatus,omitempty"`
}

func CreateOrganizationDetailGetListRequest() (request *OrganizationDetailGetListRequest) {
	request = &OrganizationDetailGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "organization/detail/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrganizationDetailGetListResponse() (response *api.BaseResponse[OrganizationDetailGetListResponse]) {
	response = api.CreateResponse[OrganizationDetailGetListResponse](&OrganizationDetailGetListResponse{})
	return
}

type OrganizationDetailGetListResponseContactTels struct {
	Zone string `json:"zone,omitempty"`
	Tel  string `json:"tel,omitempty"`
}

type OrganizationDetailGetListResponseBusinessHours struct {
	Hours []OrganizationDetailGetListResponseHours `json:"hours,omitempty"`
	Days  []int64                                  `json:"days,omitempty"`
}

type OrganizationDetailGetListResponseHours struct {
	Begin string `json:"begin,omitempty"`
	End   string `json:"end,omitempty"`
}
