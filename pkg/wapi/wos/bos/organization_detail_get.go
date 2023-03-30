package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrganizationDetailGet
// @id 1368
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1368?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询组织基本信息)
func (client *Client) OrganizationDetailGet(request *OrganizationDetailGetRequest) (response *OrganizationDetailGetResponse, err error) {
	rpcResponse := CreateOrganizationDetailGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrganizationDetailGetRequest struct {
	*api.BaseRequest

	Vid     int64   `json:"vid,omitempty"`
	VidCode string  `json:"vidCode,omitempty"`
	Fields  []int64 `json:"fields,omitempty"`
}

type OrganizationDetailGetResponse struct {
	ContactTels   []OrganizationDetailGetResponseContactTels   `json:"contactTels,omitempty"`
	BusinessHours []OrganizationDetailGetResponseBusinessHours `json:"businessHours,omitempty"`
	VidName       string                                       `json:"vidName,omitempty"`
	Vid           int64                                        `json:"vid,omitempty"`
	ParentVid     int64                                        `json:"parentVid,omitempty"`
	VidCode       string                                       `json:"vidCode,omitempty"`
	VidType       int64                                        `json:"vidType,omitempty"`
	VidStatus     int64                                        `json:"vidStatus,omitempty"`
	Logo          string                                       `json:"logo,omitempty"`
	ProvinceName  string                                       `json:"provinceName,omitempty"`
	CityName      string                                       `json:"cityName,omitempty"`
	CountyName    string                                       `json:"countyName,omitempty"`
	DetailAddress string                                       `json:"detailAddress,omitempty"`
	OnlineStatus  int64                                        `json:"onlineStatus,omitempty"`
	IsUnifyLogo   int64                                        `json:"isUnifyLogo,omitempty"`
	PictureList   []int64                                      `json:"pictureList,omitempty"`
}

func CreateOrganizationDetailGetRequest() (request *OrganizationDetailGetRequest) {
	request = &OrganizationDetailGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "organization/detail/get", "apigw")
	request.Method = api.POST
	return
}

func CreateOrganizationDetailGetResponse() (response *api.BaseResponse[OrganizationDetailGetResponse]) {
	response = api.CreateResponse[OrganizationDetailGetResponse](&OrganizationDetailGetResponse{})
	return
}

type OrganizationDetailGetResponseContactTels struct {
	Zone string `json:"zone,omitempty"`
	Tel  string `json:"tel,omitempty"`
}

type OrganizationDetailGetResponseBusinessHours struct {
	Hours []OrganizationDetailGetResponseHours `json:"hours,omitempty"`
	Days  []int64                              `json:"days,omitempty"`
}

type OrganizationDetailGetResponseHours struct {
	Begin string `json:"begin,omitempty"`
	End   string `json:"end,omitempty"`
}
