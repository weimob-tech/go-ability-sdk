package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointTransferBackgroundGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointTransferBackgroundGetListRequest, PaasWeimobPointTransferBackgroundGetListResponse]
}

func (s PaasWeimobPointTransferBackgroundGetListService) NewRequest() spi.InvocationRequest[PaasWeimobPointTransferBackgroundGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointTransferBackgroundGetListRequest]{
		Params: &PaasWeimobPointTransferBackgroundGetListRequest{},
	}
}

func (s PaasWeimobPointTransferBackgroundGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointTransferBackgroundGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobPointTransferBackgroundGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointTransferBackgroundGetListRequest]))
}

type PaasWeimobPointTransferBackgroundGetListRequest struct {
	IdentityType      int64   `json:"identityType,omitempty"`
	IdentityNo        string  `json:"identityNo,omitempty"`
	TransBeginDate    string  `json:"transBeginDate,omitempty"`
	TransEndDate      string  `json:"transEndDate,omitempty"`
	StoreVidList      []int64 `json:"storeVidList,omitempty"`
	ChannelTypeList   []int64 `json:"channelTypeList,omitempty"`
	ChangeTypeList    []int64 `json:"changeTypeList,omitempty"`
	OperateName       string  `json:"operateName,omitempty"`
	TransNo           string  `json:"transNo,omitempty"`
	FlawType          string  `json:"flawType,omitempty"`
	PageNum           int64   `json:"pageNum,omitempty"`
	PageSize          int64   `json:"pageSize,omitempty"`
	Vid               int64   `json:"vid,omitempty"`
	Wid               int64   `json:"wid,omitempty"`
	BosId             int64   `json:"bosId,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
}

type PaasWeimobPointTransferBackgroundGetListResponse struct {
	PageList   []PaasWeimobPointTransferBackgroundGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                                      `json:"pageNum,omitempty"`
	PageSize   int64                                                      `json:"pageSize,omitempty"`
	TotalCount int64                                                      `json:"totalCount,omitempty"`
}
type PaasWeimobPointTransferBackgroundGetListResponsePageList struct {
	TransNo         string `json:"transNo,omitempty"`
	TransDate       string `json:"transDate,omitempty"`
	Uid             string `json:"uid,omitempty"`
	ChangeType      string `json:"changeType,omitempty"`
	ChangeTypeDesc  string `json:"changeTypeDesc,omitempty"`
	RuleName        string `json:"ruleName,omitempty"`
	Customer        string `json:"customer,omitempty"`
	VariationRange  string `json:"variationRange,omitempty"`
	Amount          string `json:"amount,omitempty"`
	ChannelType     int64  `json:"channelType,omitempty"`
	ChannelTypeDesc string `json:"channelTypeDesc,omitempty"`
	OccurVName      string `json:"occurVName,omitempty"`
	Remark          string `json:"remark,omitempty"`
	OperateWid      string `json:"operateWid,omitempty"`
	OperateWName    string `json:"operateWName,omitempty"`
	PhoneNo         string `json:"phoneNo,omitempty"`
}

func CreatePaasWeimobPointTransferBackgroundGetListResponse() *PaasWeimobPointTransferBackgroundGetListResponse {
	return &PaasWeimobPointTransferBackgroundGetListResponse{}
}

// OnPaasWeimobPointTransferBackgroundGetListServiceInvocation
// @id 1103
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1103?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商户查询用户积分日志)
func (s *Service) OnPaasWeimobPointTransferBackgroundGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointTransferBackgroundGetListRequest, PaasWeimobPointTransferBackgroundGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobPointTransferBackgroundGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointTransferBackgroundGetListRequest, PaasWeimobPointTransferBackgroundGetListResponse](invocation),
		"PaasWeimobPointTransferBackgroundGetListService",
		"weimob.point.transfer.background.getList",
	)
	return s
}
