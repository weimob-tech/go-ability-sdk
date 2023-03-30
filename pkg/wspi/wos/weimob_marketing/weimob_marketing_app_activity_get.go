package weimob_marketing

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobMarketingAppActivityGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobMarketingAppActivityGetRequest, PaasWeimobMarketingAppActivityGetResponse]
}

func (s PaasWeimobMarketingAppActivityGetService) NewRequest() spi.InvocationRequest[PaasWeimobMarketingAppActivityGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobMarketingAppActivityGetRequest]{
		Params: &PaasWeimobMarketingAppActivityGetRequest{},
	}
}

func (s PaasWeimobMarketingAppActivityGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobMarketingAppActivityGetRequest],
) (
	spi.InvocationResponse[PaasWeimobMarketingAppActivityGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobMarketingAppActivityGetRequest]))
}

type PaasWeimobMarketingAppActivityGetRequest struct {
	ActivityId        int64 `json:"activityId,omitempty"`
	BosId             int64 `json:"bosId,omitempty"`
	Wid               int64 `json:"wid,omitempty"`
	CurrentVid        int64 `json:"currentVid,omitempty"`
	CurrentBosId      int64 `json:"currentBosId,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	Id                int64 `json:"id,omitempty"`
}

type PaasWeimobMarketingAppActivityGetResponse struct {
	VNodes            []PaasWeimobMarketingAppActivityGetResponseVNodes  `json:"vNodes,omitempty"`
	QrCodes           []PaasWeimobMarketingAppActivityGetResponseQrCodes `json:"qrCodes,omitempty"`
	Id                int64                                              `json:"id,omitempty"`
	BosId             int64                                              `json:"bosId,omitempty"`
	ProductId         int64                                              `json:"productId,omitempty"`
	ProductInstanceId int64                                              `json:"productInstanceId,omitempty"`
	Vid               int64                                              `json:"vid,omitempty"`
	VidType           int64                                              `json:"vidType,omitempty"`
	ActivityId        int64                                              `json:"activityId,omitempty"`
	ActivityName      string                                             `json:"activityName,omitempty"`
	ActivityPic       string                                             `json:"activityPic,omitempty"`
	ActivityDesc      string                                             `json:"activityDesc,omitempty"`
	StartTime         string                                             `json:"startTime,omitempty"`
	EndTime           string                                             `json:"endTime,omitempty"`
	MatchAllVid       bool                                               `json:"matchAllVid,omitempty"`
	Status            int64                                              `json:"status,omitempty"`
}
type PaasWeimobMarketingAppActivityGetResponseVNodes struct {
	BosId             int64  `json:"bosId,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Id                int64  `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Avatar            string `json:"avatar,omitempty"`
}
type PaasWeimobMarketingAppActivityGetResponseQrCodes struct {
	Channel int64  `json:"channel,omitempty"`
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	Page    string `json:"page,omitempty"`
}

func CreatePaasWeimobMarketingAppActivityGetResponse() *PaasWeimobMarketingAppActivityGetResponse {
	return &PaasWeimobMarketingAppActivityGetResponse{}
}

// OnPaasWeimobMarketingAppActivityGetServiceInvocation
// @id 941
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/941?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取外部活动详情)
func (s *Service) OnPaasWeimobMarketingAppActivityGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobMarketingAppActivityGetRequest, PaasWeimobMarketingAppActivityGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobMarketingAppActivityGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobMarketingAppActivityGetRequest, PaasWeimobMarketingAppActivityGetResponse](invocation),
		"PaasWeimobMarketingAppActivityGetService",
		"weimobMarketing.app.activity.get",
	)
	return s
}
