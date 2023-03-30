package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrganizationCreate
// @id 1683
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1683?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建组织)
func (client *Client) OrganizationCreate(request *OrganizationCreateRequest) (response *OrganizationCreateResponse, err error) {
	rpcResponse := CreateOrganizationCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrganizationCreateRequest struct {
	*api.BaseRequest

	ContactTels   OrganizationCreateRequestContactTels     `json:"contactTels,omitempty"`
	BusinessHours []OrganizationCreateRequestBusinessHours `json:"businessHours,omitempty"`
	ParentVid     int64                                    `json:"parentVid,omitempty"`
	VidName       string                                   `json:"vidName,omitempty"`
	VidCode       string                                   `json:"vidCode,omitempty"`
	VidType       int64                                    `json:"vidType,omitempty"`
	VidStatus     int64                                    `json:"vidStatus,omitempty"`
	OnlineStatus  int64                                    `json:"onlineStatus,omitempty"`
	Logo          string                                   `json:"logo,omitempty"`
	PictureList   []int64                                  `json:"pictureList,omitempty"`
	Area          int64                                    `json:"area,omitempty"`
	Address       string                                   `json:"address,omitempty"`
	Floor         string                                   `json:"floor,omitempty"`
	IsUnifyLogo   int64                                    `json:"isUnifyLogo,omitempty"`
}

type OrganizationCreateResponse struct {
	Code OrganizationCreateResponseCode `json:"code,omitempty"`
	Data OrganizationCreateResponseData `json:"data,omitempty"`
}

func CreateOrganizationCreateRequest() (request *OrganizationCreateRequest) {
	request = &OrganizationCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "organization/create", "apigw")
	request.Method = api.POST
	return
}

func CreateOrganizationCreateResponse() (response *api.BaseResponse[OrganizationCreateResponse]) {
	response = api.CreateResponse[OrganizationCreateResponse](&OrganizationCreateResponse{})
	return
}

type OrganizationCreateRequestContactTels struct {
	Zone string `json:"zone,omitempty"`
	Tel  string `json:"tel,omitempty"`
}

type OrganizationCreateRequestBusinessHours struct {
	Hours []OrganizationCreateRequestHours `json:"hours,omitempty"`
	Days  []int64                          `json:"days,omitempty"`
}

type OrganizationCreateRequestHours struct {
	Begin string `json:"begin,omitempty"`
	End   string `json:"end,omitempty"`
}

type OrganizationCreateResponseCode struct {
	Errcode string `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type OrganizationCreateResponseData struct {
	Vid int64 `json:"vid,omitempty"`
}
