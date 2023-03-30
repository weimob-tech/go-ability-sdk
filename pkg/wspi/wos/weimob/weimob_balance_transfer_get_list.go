package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceTransferGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceTransferGetListRequest, PaasWeimobBalanceTransferGetListResponse]
}

func (s PaasWeimobBalanceTransferGetListService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceTransferGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceTransferGetListRequest]{
		Params: &PaasWeimobBalanceTransferGetListRequest{},
	}
}

func (s PaasWeimobBalanceTransferGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceTransferGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceTransferGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceTransferGetListRequest]))
}

type PaasWeimobBalanceTransferGetListRequest struct {
	AmountType        int64 `json:"amountType,omitempty"`
	PageNum           int64 `json:"pageNum,omitempty"`
	PageSize          int64 `json:"pageSize,omitempty"`
	BosId             int64 `json:"bosId,omitempty"`
	Uid               int64 `json:"uid,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
}

type PaasWeimobBalanceTransferGetListResponse struct {
	PageList   []PaasWeimobBalanceTransferGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                              `json:"pageNum,omitempty"`
	PageSize   int64                                              `json:"pageSize,omitempty"`
	TotalCount int64                                              `json:"totalCount,omitempty"`
}
type PaasWeimobBalanceTransferGetListResponsePageList struct {
	TransNo         string `json:"transNo,omitempty"`
	TransDate       string `json:"transDate,omitempty"`
	ChangeType      string `json:"changeType,omitempty"`
	IsPositive      int64  `json:"isPositive,omitempty"`
	BalanceChange   string `json:"balanceChange,omitempty"`
	PrincipalChange string `json:"principalChange,omitempty"`
	BonusChange     string `json:"bonusChange,omitempty"`
	Balance         string `json:"balance,omitempty"`
	Principal       string `json:"principal,omitempty"`
	Bonus           string `json:"bonus,omitempty"`
	Remark          string `json:"remark,omitempty"`
	RuleName        string `json:"ruleName,omitempty"`
}

func CreatePaasWeimobBalanceTransferGetListResponse() *PaasWeimobBalanceTransferGetListResponse {
	return &PaasWeimobBalanceTransferGetListResponse{}
}

// OnPaasWeimobBalanceTransferGetListServiceInvocation
// @id 1107
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1107?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户余额日志)
func (s *Service) OnPaasWeimobBalanceTransferGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceTransferGetListRequest, PaasWeimobBalanceTransferGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceTransferGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceTransferGetListRequest, PaasWeimobBalanceTransferGetListResponse](invocation),
		"PaasWeimobBalanceTransferGetListService",
		"weimob.balance.transfer.getList",
	)
	return s
}
