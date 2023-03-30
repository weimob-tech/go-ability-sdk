package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrganizationUpdate
// @id 1687
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1687?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新组织基本信息)
func (client *Client) OrganizationUpdate(request *OrganizationUpdateRequest) (response *OrganizationUpdateResponse, err error) {
	rpcResponse := CreateOrganizationUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrganizationUpdateRequest struct {
	*api.BaseRequest

	ContactTels   OrganizationUpdateRequestContactTels     `json:"contactTels,omitempty"`
	BusinessHours []OrganizationUpdateRequestBusinessHours `json:"businessHours,omitempty"`
	Vid           int64                                    `json:"vid,omitempty"`
	VidName       string                                   `json:"vidName,omitempty"`
	VidCode       string                                   `json:"vidCode,omitempty"`
	VidStatus     int64                                    `json:"vidStatus,omitempty"`
	OnlineStatus  int64                                    `json:"onlineStatus,omitempty"`
	Logo          string                                   `json:"logo,omitempty"`
	PictureList   []int64                                  `json:"pictureList,omitempty"`
	Area          int64                                    `json:"area,omitempty"`
	Address       string                                   `json:"address,omitempty"`
	Floor         string                                   `json:"floor,omitempty"`
	IsUnifyLogo   int64                                    `json:"isUnifyLogo,omitempty"`
}

type OrganizationUpdateResponse struct {
}

func CreateOrganizationUpdateRequest() (request *OrganizationUpdateRequest) {
	request = &OrganizationUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "organization/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrganizationUpdateResponse() (response *api.BaseResponse[OrganizationUpdateResponse]) {
	response = api.CreateResponse[OrganizationUpdateResponse](&OrganizationUpdateResponse{})
	return
}

type OrganizationUpdateRequestContactTels struct {
	Zone string `json:"zone,omitempty"`
	Tel  string `json:"tel,omitempty"`
}

type OrganizationUpdateRequestBusinessHours struct {
	Hours []OrganizationUpdateRequestHours `json:"hours,omitempty"`
	Days  []int64                          `json:"days,omitempty"`
}

type OrganizationUpdateRequestHours struct {
	Begin string `json:"begin,omitempty"`
	End   string `json:"end,omitempty"`
}
