package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCsServicerExclusiveSyncService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCsrExclusiveSyncRequest, PaasWeimobGuideCsrExclusiveSyncResponse]
}

func (s PaasWeimobGuideCsServicerExclusiveSyncService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCsrExclusiveSyncRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCsrExclusiveSyncRequest]{
		Params: &PaasWeimobGuideCsrExclusiveSyncRequest{},
	}
}

func (s PaasWeimobGuideCsServicerExclusiveSyncService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCsrExclusiveSyncRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCsrExclusiveSyncResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCsrExclusiveSyncRequest]))
}

type PaasWeimobGuideCsrExclusiveSyncRequest struct {
	BosId                   int64   `json:"bosId,omitempty"`
	ExclusiveType           int64   `json:"exclusiveType,omitempty"`
	TargetProductInstanceId int64   `json:"targetProductInstanceId,omitempty"`
	ServicerIdList          []int64 `json:"servicerIdList,omitempty"`
}

type PaasWeimobGuideCsrExclusiveSyncResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreatePaasWeimobGuideCsrExclusiveSyncResponse() *PaasWeimobGuideCsrExclusiveSyncResponse {
	return &PaasWeimobGuideCsrExclusiveSyncResponse{}
}

// OnPaasWeimobGuideCsServicerExclusiveSyncServiceInvocation
// @id 725
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/725?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=设置/取消专属客服)
func (s *Service) OnPaasWeimobGuideCsServicerExclusiveSyncServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCsrExclusiveSyncRequest, PaasWeimobGuideCsrExclusiveSyncResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCsServicerExclusiveSyncService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCsrExclusiveSyncRequest, PaasWeimobGuideCsrExclusiveSyncResponse](invocation),
		"PaasWeimobGuideCsServicerExclusiveSyncService",
		"weimobGuide.cs.servicer.exclusive.sync",
	)
	return s
}
