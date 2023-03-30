package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointTransferGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointTransferGetListRequest, PaasWeimobPointTransferGetListResponse]
}

func (s PaasWeimobPointTransferGetListService) NewRequest() spi.InvocationRequest[PaasWeimobPointTransferGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointTransferGetListRequest]{
		Params: &PaasWeimobPointTransferGetListRequest{},
	}
}

func (s PaasWeimobPointTransferGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointTransferGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobPointTransferGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointTransferGetListRequest]))
}

type PaasWeimobPointTransferGetListRequest struct {
	ModifyType        int64 `json:"modifyType,omitempty"`
	PageNum           int64 `json:"pageNum,omitempty"`
	PageSize          int64 `json:"pageSize,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	PointPlanId       int64 `json:"pointPlanId,omitempty"`
	Wid               int64 `json:"wid,omitempty"`
	BosId             int64 `json:"bosId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
}

type PaasWeimobPointTransferGetListResponse struct {
	PageList   []PaasWeimobPointTransferGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                            `json:"pageNum,omitempty"`
	PageSize   int64                                            `json:"pageSize,omitempty"`
	TotalCount int64                                            `json:"totalCount,omitempty"`
}
type PaasWeimobPointTransferGetListResponsePageList struct {
	Point        string `json:"point,omitempty"`
	LeftPoint    string `json:"leftPoint,omitempty"`
	ChangeTime   string `json:"changeTime,omitempty"`
	ChangeReason string `json:"changeReason,omitempty"`
	OperatorName string `json:"operatorName,omitempty"`
}

func CreatePaasWeimobPointTransferGetListResponse() *PaasWeimobPointTransferGetListResponse {
	return &PaasWeimobPointTransferGetListResponse{}
}

// OnPaasWeimobPointTransferGetListServiceInvocation
// @id 1104
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1104?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户积分日志)
func (s *Service) OnPaasWeimobPointTransferGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointTransferGetListRequest, PaasWeimobPointTransferGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobPointTransferGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointTransferGetListRequest, PaasWeimobPointTransferGetListResponse](invocation),
		"PaasWeimobPointTransferGetListService",
		"weimob.point.transfer.getList",
	)
	return s
}
