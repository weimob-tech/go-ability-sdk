package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCdpGuiderGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCdpGuiderGetListRequest, PaasWeimobCdpGuiderGetListResponse]
}

func (s PaasWeimobCdpGuiderGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCdpGuiderGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCdpGuiderGetListRequest]{
		Params: &PaasWeimobCdpGuiderGetListRequest{},
	}
}

func (s PaasWeimobCdpGuiderGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCdpGuiderGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCdpGuiderGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCdpGuiderGetListRequest]))
}

type PaasWeimobCdpGuiderGetListRequest struct {
	QueryParameter    PaasWeimobCdpGuiderGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	MerchantId        int64                                           `json:"merchantId,omitempty"`
	BosId             int64                                           `json:"bosId,omitempty"`
	Vid               int64                                           `json:"vid,omitempty"`
	VidType           int64                                           `json:"vidType,omitempty"`
	ProductId         int64                                           `json:"productId,omitempty"`
	ProductVersionId  int64                                           `json:"productVersionId,omitempty"`
	ProductInstanceId int64                                           `json:"productInstanceId,omitempty"`
	Cid               int64                                           `json:"cid,omitempty"`
	Tcode             string                                          `json:"tcode,omitempty"`
	PageNum           int64                                           `json:"pageNum,omitempty"`
	PageSize          int64                                           `json:"pageSize,omitempty"`
}
type PaasWeimobCdpGuiderGetListRequestQueryParameter struct {
	GuiderVid       int64   `json:"guiderVid,omitempty"`
	GuiderName      string  `json:"guiderName,omitempty"`
	GuiderPhoneList []int64 `json:"guiderPhoneList,omitempty"`
	GuiderWidList   []int64 `json:"guiderWidList,omitempty"`
}

type PaasWeimobCdpGuiderGetListResponse struct {
	PageList   []PaasWeimobCdpGuiderGetListResponsePageList `json:"pageList,omitempty"`
	TotalCount int64                                        `json:"totalCount,omitempty"`
	PageNum    int64                                        `json:"pageNum,omitempty"`
	PageSize   int64                                        `json:"pageSize,omitempty"`
}
type PaasWeimobCdpGuiderGetListResponsePageList struct {
	GuiderPhone    string `json:"guiderPhone,omitempty"`
	GuiderName     string `json:"guiderName,omitempty"`
	GuiderWid      int64  `json:"guiderWid,omitempty"`
	GuiderVid      int64  `json:"guiderVid,omitempty"`
	JobNumber      int64  `json:"jobNumber,omitempty"`
	GuiderId       string `json:"guiderId,omitempty"`
	IsUsed         int64  `json:"isUsed,omitempty"`
	IsExclusiveCus int64  `json:"isExclusiveCus,omitempty"`
}

func CreatePaasWeimobCdpGuiderGetListResponse() *PaasWeimobCdpGuiderGetListResponse {
	return &PaasWeimobCdpGuiderGetListResponse{}
}

// OnPaasWeimobCdpGuiderGetListServiceInvocation
// @id 1170
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1170?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导购-分页查询导购列表)
func (s *Service) OnPaasWeimobCdpGuiderGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCdpGuiderGetListRequest, PaasWeimobCdpGuiderGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCdpGuiderGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCdpGuiderGetListRequest, PaasWeimobCdpGuiderGetListResponse](invocation),
		"PaasWeimobCdpGuiderGetListService",
		"weimobCdp.guider.getList",
	)
	return s
}
