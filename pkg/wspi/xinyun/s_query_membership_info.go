package xinyun

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasQueryMembershipInfoService struct {
	handler spi.XinyunInvocationHandler[PaasQueryMembershipInfoRequest, PaasQueryMembershipInfoResponse]
}

func (s PaasQueryMembershipInfoService) NewRequest() spi.InvocationRequest[PaasQueryMembershipInfoRequest] {
	return &spi.XinyunInvocationRequest[PaasQueryMembershipInfoRequest]{
		Params: &PaasQueryMembershipInfoRequest{},
	}
}

func (s PaasQueryMembershipInfoService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasQueryMembershipInfoRequest],
) (
	spi.InvocationResponse[PaasQueryMembershipInfoResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.XinyunInvocationRequest[PaasQueryMembershipInfoRequest]))
}

type PaasQueryMembershipInfoRequest struct {
	SourceObject PaasQueryMembershipInfoRequestSourceObject `json:"sourceObject,omitempty"`
	Pid          int64                                      `json:"pid,omitempty"`
	StoreId      int64                                      `json:"storeId,omitempty"`
	Wid          int64                                      `json:"wid,omitempty"`
}
type PaasQueryMembershipInfoRequestSourceObject struct {
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
	Source       string `json:"source,omitempty"`
}

type PaasQueryMembershipInfoResponse struct {
	PassMembershipInfos []int64 `json:"passMembershipInfos,omitempty"`
	Phone               string  `json:"phone,omitempty"`
	IsMember            bool    `json:"isMember,omitempty"`
	Status              int64   `json:"status,omitempty"`
}

func CreatePaasQueryMembershipInfoResponse() *PaasQueryMembershipInfoResponse {
	return &PaasQueryMembershipInfoResponse{}
}

// OnPaasQueryMembershipInfoServiceInvocation
// @id 1780
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=查询会员信息)
func (s *Service) OnPaasQueryMembershipInfoServiceInvocation(
	handler spi.XinyunInvocationHandler[PaasQueryMembershipInfoRequest, PaasQueryMembershipInfoResponse],
) (service *Service) {
	var invocation = &PaasQueryMembershipInfoService{handler}
	s.Register(
		spi.WrapInvoker[PaasQueryMembershipInfoRequest, PaasQueryMembershipInfoResponse](invocation),
		"PaasQueryMembershipInfoService",
		"sQueryMembershipInfo",
	)
	return s
}
