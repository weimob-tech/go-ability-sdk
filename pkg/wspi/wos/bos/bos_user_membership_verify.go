package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasBosUserMembershipVerifyService struct {
	handler spi.WosInvocationHandler[PaasBosUserMembershipVerifyRequest, PaasBosUserMembershipVerifyResponse]
}

func (s PaasBosUserMembershipVerifyService) NewRequest() spi.InvocationRequest[PaasBosUserMembershipVerifyRequest] {
	return &spi.WosInvocationRequest[PaasBosUserMembershipVerifyRequest]{
		Params: &PaasBosUserMembershipVerifyRequest{},
	}
}

func (s PaasBosUserMembershipVerifyService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasBosUserMembershipVerifyRequest],
) (
	spi.InvocationResponse[PaasBosUserMembershipVerifyResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasBosUserMembershipVerifyRequest]))
}

type PaasBosUserMembershipVerifyRequest struct {
}

type PaasBosUserMembershipVerifyResponse struct {
}

func CreatePaasBosUserMembershipVerifyResponse() *PaasBosUserMembershipVerifyResponse {
	return &PaasBosUserMembershipVerifyResponse{}
}

// OnPaasBosUserMembershipVerifyServiceInvocation
// @id 587
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/587?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=检查用户会员身份)
func (s *Service) OnPaasBosUserMembershipVerifyServiceInvocation(
	handler spi.WosInvocationHandler[PaasBosUserMembershipVerifyRequest, PaasBosUserMembershipVerifyResponse],
) (service *Service) {
	var invocation = &PaasBosUserMembershipVerifyService{handler}
	s.Register(
		spi.WrapInvoker[PaasBosUserMembershipVerifyRequest, PaasBosUserMembershipVerifyResponse](invocation),
		"PaasBosUserMembershipVerifyService",
		"bosUserMembershipVerify",
	)
	return s
}
