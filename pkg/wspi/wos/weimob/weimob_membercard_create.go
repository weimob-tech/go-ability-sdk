package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobMembercardCreateService struct {
	handler spi.WosInvocationHandler[PaasWeimobMembercardCreateRequest, PaasWeimobMembercardCreateResponse]
}

func (s PaasWeimobMembercardCreateService) NewRequest() spi.InvocationRequest[PaasWeimobMembercardCreateRequest] {
	return &spi.WosInvocationRequest[PaasWeimobMembercardCreateRequest]{
		Params: &PaasWeimobMembercardCreateRequest{},
	}
}

func (s PaasWeimobMembercardCreateService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobMembercardCreateRequest],
) (
	spi.InvocationResponse[PaasWeimobMembercardCreateResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobMembercardCreateRequest]))
}

type PaasWeimobMembercardCreateRequest struct {
	GroupFieldInfos       []PaasWeimobMembercardCreateRequestGroupFieldInfos  `json:"groupFieldInfos,omitempty"`
	SourceObjectList      []PaasWeimobMembercardCreateRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Phone                 string                                              `json:"phone,omitempty"`
	MembershipPlanId      int64                                               `json:"membershipPlanId,omitempty"`
	Wid                   int64                                               `json:"wid,omitempty"`
	MembershipCardScene   int64                                               `json:"membershipCardScene,omitempty"`
	MembershipCardSceneId int64                                               `json:"membershipCardSceneId,omitempty"`
	LevelId               int64                                               `json:"levelId,omitempty"`
	EffectiveTime         int64                                               `json:"effectiveTime,omitempty"`
	ExpireTime            int64                                               `json:"expireTime,omitempty"`
	GuiderWid             int64                                               `json:"guiderWid,omitempty"`
}
type PaasWeimobMembercardCreateRequestGroupFieldInfos struct {
	FieldInfo PaasWeimobMembercardCreateRequestFieldInfo `json:"fieldInfo,omitempty"`
	GroupId   int64                                      `json:"groupId,omitempty"`
}
type PaasWeimobMembercardCreateRequestFieldInfo struct {
	GroupWholeFieldInfos []PaasWeimobMembercardCreateRequestGroupWholeFieldInfos `json:"groupWholeFieldInfos,omitempty"`
}
type PaasWeimobMembercardCreateRequestGroupWholeFieldInfos struct {
	FieldValues []PaasWeimobMembercardCreateRequestFieldValues `json:"fieldValues,omitempty"`
}
type PaasWeimobMembercardCreateRequestFieldValues struct {
	FieldId    int64  `json:"fieldId,omitempty"`
	FieldKey   string `json:"fieldKey,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
}
type PaasWeimobMembercardCreateRequestSourceObjectList struct {
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
	Source       int64  `json:"source,omitempty"`
}

type PaasWeimobMembercardCreateResponse struct {
	CustomCardNo string `json:"customCardNo,omitempty"`
}

func CreatePaasWeimobMembercardCreateResponse() *PaasWeimobMembercardCreateResponse {
	return &PaasWeimobMembercardCreateResponse{}
}

// OnPaasWeimobMembercardCreateServiceInvocation
// @id 1098
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1098?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=开通会员卡)
func (s *Service) OnPaasWeimobMembercardCreateServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobMembercardCreateRequest, PaasWeimobMembercardCreateResponse],
) (service *Service) {
	var invocation = &PaasWeimobMembercardCreateService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobMembercardCreateRequest, PaasWeimobMembercardCreateResponse](invocation),
		"PaasWeimobMembercardCreateService",
		"weimob.membercard.create",
	)
	return s
}
