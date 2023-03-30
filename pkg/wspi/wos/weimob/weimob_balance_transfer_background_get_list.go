package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceTransferBackgroundGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceTransferBackgroundGetListRequest, PaasWeimobBalanceTransferBackgroundGetListResponse]
}

func (s PaasWeimobBalanceTransferBackgroundGetListService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceTransferBackgroundGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceTransferBackgroundGetListRequest]{
		Params: &PaasWeimobBalanceTransferBackgroundGetListRequest{},
	}
}

func (s PaasWeimobBalanceTransferBackgroundGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceTransferBackgroundGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceTransferBackgroundGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceTransferBackgroundGetListRequest]))
}

type PaasWeimobBalanceTransferBackgroundGetListRequest struct {
	IdentityType      int64   `json:"identityType,omitempty"`
	IdentityNo        string  `json:"identityNo,omitempty"`
	TransBeginDate    string  `json:"transBeginDate,omitempty"`
	TransEndDate      string  `json:"transEndDate,omitempty"`
	StoreVidList      []int64 `json:"storeVidList,omitempty"`
	ChannelTypeList   []int64 `json:"channelTypeList,omitempty"`
	ChangeTypeList    []int64 `json:"changeTypeList,omitempty"`
	OperateName       string  `json:"operateName,omitempty"`
	TransNo           string  `json:"transNo,omitempty"`
	OutTransNo        string  `json:"outTransNo,omitempty"`
	StartAmount       int64   `json:"startAmount,omitempty"`
	EndAmount         int64   `json:"endAmount,omitempty"`
	PageNum           int64   `json:"pageNum,omitempty"`
	PageSize          int64   `json:"pageSize,omitempty"`
	BosId             int64   `json:"bosId,omitempty"`
	Uid               int64   `json:"uid,omitempty"`
	Vid               int64   `json:"vid,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
}

type PaasWeimobBalanceTransferBackgroundGetListResponse struct {
	PageList   []PaasWeimobBalanceTransferBackgroundGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                                        `json:"pageNum,omitempty"`
	PageSize   int64                                                        `json:"pageSize,omitempty"`
	TotalCount int64                                                        `json:"totalCount,omitempty"`
}
type PaasWeimobBalanceTransferBackgroundGetListResponsePageList struct {
	TransNo         string `json:"transNo,omitempty"`
	TransDate       string `json:"transDate,omitempty"`
	Wid             string `json:"wid,omitempty"`
	WidName         string `json:"widName,omitempty"`
	ChangeType      string `json:"changeType,omitempty"`
	ChangeTypeDesc  string `json:"changeTypeDesc,omitempty"`
	RuleName        string `json:"ruleName,omitempty"`
	BalanceChange   string `json:"balanceChange,omitempty"`
	Balance         string `json:"balance,omitempty"`
	PrincipalChange string `json:"principalChange,omitempty"`
	Principal       string `json:"principal,omitempty"`
	BonusChange     string `json:"bonusChange,omitempty"`
	Bonus           string `json:"bonus,omitempty"`
	OutTransNo      string `json:"outTransNo,omitempty"`
	ChannelType     int64  `json:"channelType,omitempty"`
	ChannelTypeDesc string `json:"channelTypeDesc,omitempty"`
	Remark          string `json:"remark,omitempty"`
	OccurVName      string `json:"occurVName,omitempty"`
	OperateWName    string `json:"operateWName,omitempty"`
	PhoneNo         string `json:"phoneNo,omitempty"`
	IsPositive      int64  `json:"isPositive,omitempty"`
}

func CreatePaasWeimobBalanceTransferBackgroundGetListResponse() *PaasWeimobBalanceTransferBackgroundGetListResponse {
	return &PaasWeimobBalanceTransferBackgroundGetListResponse{}
}

// OnPaasWeimobBalanceTransferBackgroundGetListServiceInvocation
// @id 1106
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1106?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商户查询用户余额日志)
func (s *Service) OnPaasWeimobBalanceTransferBackgroundGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceTransferBackgroundGetListRequest, PaasWeimobBalanceTransferBackgroundGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceTransferBackgroundGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceTransferBackgroundGetListRequest, PaasWeimobBalanceTransferBackgroundGetListResponse](invocation),
		"PaasWeimobBalanceTransferBackgroundGetListService",
		"weimob.balance.transfer.background.getList",
	)
	return s
}
