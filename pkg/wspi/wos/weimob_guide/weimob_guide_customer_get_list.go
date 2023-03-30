package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerGetListRequest, PaasWeimobGuideCustomerGetListResponse]
}

func (s PaasWeimobGuideCustomerGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerGetListRequest]{
		Params: &PaasWeimobGuideCustomerGetListRequest{},
	}
}

func (s PaasWeimobGuideCustomerGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerGetListRequest]))
}

type PaasWeimobGuideCustomerGetListRequest struct {
	Vid               int64   `json:"vid,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
	BosId             int64   `json:"bosId,omitempty"`
	VidType           int64   `json:"vidType,omitempty"`
	WidList           []int64 `json:"widList,omitempty"`
}

type PaasWeimobGuideCustomerGetListResponse struct {
	UserBaseInfoList []PaasWeimobGuideCustomerGetListResponseUserBaseInfoList `json:"userBaseInfoList,omitempty"`
}
type PaasWeimobGuideCustomerGetListResponseUserBaseInfoList struct {
	Wid               int64  `json:"wid,omitempty"`
	NickName          string `json:"nickName,omitempty"`
	AvatarUrl         int64  `json:"avatarUrl,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Name              int64  `json:"name,omitempty"`
	Sex               int64  `json:"sex,omitempty"`
	Gender            string `json:"gender,omitempty"`
	IsSubscribe       string `json:"isSubscribe,omitempty"`
	Email             int64  `json:"email,omitempty"`
	BelongVid         int64  `json:"belongVid,omitempty"`
	BelongVidBindTime string `json:"belongVidBindTime,omitempty"`
	IdentityCardNum   string `json:"identityCardNum,omitempty"`
}

func CreatePaasWeimobGuideCustomerGetListResponse() *PaasWeimobGuideCustomerGetListResponse {
	return &PaasWeimobGuideCustomerGetListResponse{}
}

// OnPaasWeimobGuideCustomerGetListServiceInvocation
// @id 750
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/750?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户列表)
func (s *Service) OnPaasWeimobGuideCustomerGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerGetListRequest, PaasWeimobGuideCustomerGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerGetListRequest, PaasWeimobGuideCustomerGetListResponse](invocation),
		"PaasWeimobGuideCustomerGetListService",
		"weimobGuide.customer.getList",
	)
	return s
}
