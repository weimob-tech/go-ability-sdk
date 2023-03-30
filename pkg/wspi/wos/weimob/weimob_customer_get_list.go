package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCustomerGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCustomerGetListRequest, PaasWeimobCustomerGetListResponse]
}

func (s PaasWeimobCustomerGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCustomerGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCustomerGetListRequest]{
		Params: &PaasWeimobCustomerGetListRequest{},
	}
}

func (s PaasWeimobCustomerGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCustomerGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCustomerGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCustomerGetListRequest]))
}

type PaasWeimobCustomerGetListRequest struct {
	WidList []int64 `json:"widList,omitempty"`
}

type PaasWeimobCustomerGetListResponse struct {
	UserBaseInfoList []PaasWeimobCustomerGetListResponseUserBaseInfoList `json:"userBaseInfoList,omitempty"`
}
type PaasWeimobCustomerGetListResponseUserBaseInfoList struct {
	MemberIdentifyInfoList []PaasWeimobCustomerGetListResponseMemberIdentifyInfoList `json:"memberIdentifyInfoList,omitempty"`
	ExtendInfo             []PaasWeimobCustomerGetListResponseExtendInfo             `json:"extendInfo,omitempty"`
	Phone                  string                                                    `json:"phone,omitempty"`
	Wid                    int64                                                     `json:"wid,omitempty"`
	Name                   string                                                    `json:"name,omitempty"`
	Nickname               string                                                    `json:"nickname,omitempty"`
	AvatarUrl              string                                                    `json:"avatarUrl,omitempty"`
	Birthday               int64                                                     `json:"birthday,omitempty"`
	IdentityCardNum        string                                                    `json:"identityCardNum,omitempty"`
	Gender                 int64                                                     `json:"gender,omitempty"`
	Email                  string                                                    `json:"email,omitempty"`
	Education              int64                                                     `json:"education,omitempty"`
	Hobby                  string                                                    `json:"hobby,omitempty"`
	Income                 int64                                                     `json:"income,omitempty"`
	Industry               string                                                    `json:"industry,omitempty"`
	BecomeCustomerTime     int64                                                     `json:"becomeCustomerTime,omitempty"`
	BecomeMemberTime       int64                                                     `json:"becomeMemberTime,omitempty"`
	BelongVid              int64                                                     `json:"belongVid,omitempty"`
	BelongVidBindTime      int64                                                     `json:"belongVidBindTime,omitempty"`
	IsCustomer             bool                                                      `json:"isCustomer,omitempty"`
	IsFans                 bool                                                      `json:"isFans,omitempty"`
	HasValidMemberCard     bool                                                      `json:"hasValidMemberCard,omitempty"`
}
type PaasWeimobCustomerGetListResponseMemberIdentifyInfoList struct {
	LevelId          int64 `json:"levelId,omitempty"`
	MembershipPlanId int64 `json:"membershipPlanId,omitempty"`
	MembershipType   int64 `json:"membershipType,omitempty"`
}
type PaasWeimobCustomerGetListResponseExtendInfo struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func CreatePaasWeimobCustomerGetListResponse() *PaasWeimobCustomerGetListResponse {
	return &PaasWeimobCustomerGetListResponse{}
}

// OnPaasWeimobCustomerGetListServiceInvocation
// @id 1090
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1090?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取客户基础信息)
func (s *Service) OnPaasWeimobCustomerGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCustomerGetListRequest, PaasWeimobCustomerGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCustomerGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCustomerGetListRequest, PaasWeimobCustomerGetListResponse](invocation),
		"PaasWeimobCustomerGetListService",
		"weimob.customer.getList",
	)
	return s
}
