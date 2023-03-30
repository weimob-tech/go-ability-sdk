package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardCreate
// @id 1770
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1770?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建会员卡)
func (client *Client) MembercardCreate(request *MembercardCreateRequest) (response *MembercardCreateResponse, err error) {
	rpcResponse := CreateMembercardCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardCreateRequest struct {
	*api.BaseRequest

	GroupFieldInfos       []MembercardCreateRequestGroupFieldInfos `json:"groupFieldInfos,omitempty"`
	MembershipCardSceneId int64                                    `json:"membershipCardSceneId,omitempty"`
	MembershipPlanId      int64                                    `json:"membershipPlanId,omitempty"`
	LevelId               int64                                    `json:"levelId,omitempty"`
	MembershipCardScene   int64                                    `json:"membershipCardScene,omitempty"`
	Wid                   int64                                    `json:"wid,omitempty"`
	ExpireTime            int64                                    `json:"expireTime,omitempty"`
	EffectiveTime         int64                                    `json:"effectiveTime,omitempty"`
	MembershipCardChannel int64                                    `json:"membershipCardChannel,omitempty"`
	Vid                   int64                                    `json:"vid,omitempty"`
	VidType               int64                                    `json:"vidType,omitempty"`
}

type MembercardCreateResponse struct {
	CustomCardNo string `json:"customCardNo,omitempty"`
}

func CreateMembercardCreateRequest() (request *MembercardCreateRequest) {
	request = &MembercardCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/create", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardCreateResponse() (response *api.BaseResponse[MembercardCreateResponse]) {
	response = api.CreateResponse[MembercardCreateResponse](&MembercardCreateResponse{})
	return
}

type MembercardCreateRequestGroupFieldInfos struct {
	FieldInfo MembercardCreateRequestFieldInfo `json:"fieldInfo,omitempty"`
	GroupId   int64                            `json:"groupId,omitempty"`
}

type MembercardCreateRequestFieldInfo struct {
	GroupWholeFieldInfos []MembercardCreateRequestGroupWholeFieldInfos `json:"groupWholeFieldInfos,omitempty"`
}

type MembercardCreateRequestGroupWholeFieldInfos struct {
	FieldValues []MembercardCreateRequestFieldValues `json:"fieldValues,omitempty"`
}

type MembercardCreateRequestFieldValues struct {
	FieldId    int64  `json:"fieldId,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
}
