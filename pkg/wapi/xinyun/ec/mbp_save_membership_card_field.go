package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpSaveMembershipCardField
// @id 1680
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=保存客户资料设置)
func (client *Client) MbpSaveMembershipCardField(request *MbpSaveMembershipCardFieldRequest) (response *MbpSaveMembershipCardFieldResponse, err error) {
	rpcResponse := CreateMbpSaveMembershipCardFieldResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpSaveMembershipCardFieldRequest struct {
	*api.BaseRequest

	ListGroup []MbpSaveMembershipCardFieldRequestListGroup `json:"listGroup,omitempty"`
}

type MbpSaveMembershipCardFieldResponse struct {
}

func CreateMbpSaveMembershipCardFieldRequest() (request *MbpSaveMembershipCardFieldRequest) {
	request = &MbpSaveMembershipCardFieldRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/saveMembershipCardField", "api")
	request.Method = api.POST
	return
}

func CreateMbpSaveMembershipCardFieldResponse() (response *api.BaseResponse[MbpSaveMembershipCardFieldResponse]) {
	response = api.CreateResponse[MbpSaveMembershipCardFieldResponse](&MbpSaveMembershipCardFieldResponse{})
	return
}

type MbpSaveMembershipCardFieldRequestListGroup struct {
	FieldList     []MbpSaveMembershipCardFieldRequestFieldList `json:"fieldList,omitempty"`
	GroupId       int64                                        `json:"groupId,omitempty"`
	GroupName     string                                       `json:"groupName,omitempty"`
	Type          int                                          `json:"type,omitempty"`
	IsFieldRepeat int                                          `json:"isFieldRepeat,omitempty"`
	MaxLimitGroup int                                          `json:"maxLimitGroup,omitempty"`
	Sort          int                                          `json:"sort,omitempty"`
}

type MbpSaveMembershipCardFieldRequestFieldList struct {
	OptionList  []MbpSaveMembershipCardFieldRequestOptionList `json:"optionList,omitempty"`
	FieldRule   MbpSaveMembershipCardFieldRequestFieldRule    `json:"fieldRule,omitempty"`
	Id          int                                           `json:"id,omitempty"`
	Type        int                                           `json:"type,omitempty"`
	Key         string                                        `json:"key,omitempty"`
	Name        string                                        `json:"name,omitempty"`
	Required    bool                                          `json:"required,omitempty"`
	CanModify   bool                                          `json:"canModify,omitempty"`
	Enabled     bool                                          `json:"enabled,omitempty"`
	Seq         int                                           `json:"seq,omitempty"`
	Placeholder string                                        `json:"placeholder,omitempty"`
	OptionType  string                                        `json:"optionType,omitempty"`
	MaxLen      int                                           `json:"maxLen,omitempty"`
}

type MbpSaveMembershipCardFieldRequestOptionList struct {
	Title string `json:"title,omitempty"`
}

type MbpSaveMembershipCardFieldRequestFieldRule struct {
	ModifyType  int     `json:"modifyType,omitempty"`
	ModifyCount int     `json:"modifyCount,omitempty"`
	StartTime   int     `json:"startTime,omitempty"`
	EndTime     int     `json:"endTime,omitempty"`
	MinNumber   float64 `json:"minNumber,omitempty"`
	MaxNumber   float64 `json:"maxNumber,omitempty"`
}
