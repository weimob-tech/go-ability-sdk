package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGetList
// @id 1354
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1354?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取客户基础信息)
func (client *Client) CustomerGetList(request *CustomerGetListRequest) (response *CustomerGetListResponse, err error) {
	rpcResponse := CreateCustomerGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGetListRequest struct {
	*api.BaseRequest

	Vid        int64   `json:"vid,omitempty"`
	WidList    []int64 `json:"widList,omitempty"`
	ResultType []int64 `json:"resultType,omitempty"`
}

type CustomerGetListResponse struct {
	UserBaseInfoList []CustomerGetListResponseUserBaseInfoList `json:"userBaseInfoList,omitempty"`
}

func CreateCustomerGetListRequest() (request *CustomerGetListRequest) {
	request = &CustomerGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerGetListResponse() (response *api.BaseResponse[CustomerGetListResponse]) {
	response = api.CreateResponse[CustomerGetListResponse](&CustomerGetListResponse{})
	return
}

type CustomerGetListResponseUserBaseInfoList struct {
	MemberIdentifyInfoList []CustomerGetListResponseMemberIdentifyInfoList `json:"memberIdentifyInfoList,omitempty"`
	ExtendInfoList         []CustomerGetListResponseExtendInfoList         `json:"extendInfoList,omitempty"`
	Wid                    int64                                           `json:"wid,omitempty"`
	Name                   string                                          `json:"name,omitempty"`
	Nickname               string                                          `json:"nickname,omitempty"`
	AvatarUrl              string                                          `json:"avatarUrl,omitempty"`
	Phone                  string                                          `json:"phone,omitempty"`
	IdentityCardNum        string                                          `json:"identityCardNum,omitempty"`
	Email                  string                                          `json:"email,omitempty"`
	Birthday               int64                                           `json:"birthday,omitempty"`
	Gender                 int64                                           `json:"gender,omitempty"`
	Education              int64                                           `json:"education,omitempty"`
	Income                 int64                                           `json:"income,omitempty"`
	Industry               string                                          `json:"industry,omitempty"`
	Hobby                  string                                          `json:"hobby,omitempty"`
	BelongVid              int64                                           `json:"belongVid,omitempty"`
	BelongVidBindTime      int64                                           `json:"belongVidBindTime,omitempty"`
	BecomeCustomerTime     int64                                           `json:"becomeCustomerTime,omitempty"`
	BecomeMemberTime       int64                                           `json:"becomeMemberTime,omitempty"`
	IsFans                 bool                                            `json:"isFans,omitempty"`
	IsCustomer             bool                                            `json:"isCustomer,omitempty"`
	HasValidMemberCard     bool                                            `json:"hasValidMemberCard,omitempty"`
	Status                 int64                                           `json:"status,omitempty"`
	SourceChannel          []int64                                         `json:"sourceChannel,omitempty"`
}

type CustomerGetListResponseMemberIdentifyInfoList struct {
	MembershipPlanId int64 `json:"membershipPlanId,omitempty"`
	MembershipType   int64 `json:"membershipType,omitempty"`
	LevelId          int64 `json:"levelId,omitempty"`
}

type CustomerGetListResponseExtendInfoList struct {
	Key      string `json:"key,omitempty"`
	Value    string `json:"value,omitempty"`
	GroupNum int64  `json:"groupNum,omitempty"`
	FieldId  int64  `json:"fieldId,omitempty"`
}
