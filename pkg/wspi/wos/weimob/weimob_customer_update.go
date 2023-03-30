package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCustomerUpdateService struct {
	handler spi.WosInvocationHandler[PaasWeimobCustomerUpdateRequest, PaasWeimobCustomerUpdateResponse]
}

func (s PaasWeimobCustomerUpdateService) NewRequest() spi.InvocationRequest[PaasWeimobCustomerUpdateRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCustomerUpdateRequest]{
		Params: &PaasWeimobCustomerUpdateRequest{},
	}
}

func (s PaasWeimobCustomerUpdateService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCustomerUpdateRequest],
) (
	spi.InvocationResponse[PaasWeimobCustomerUpdateResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCustomerUpdateRequest]))
}

type PaasWeimobCustomerUpdateRequest struct {
	ExtInfo         PaasWeimobCustomerUpdateRequestExtInfo `json:"extInfo,omitempty"`
	Wid             int64                                  `json:"wid,omitempty"`
	Phone           string                                 `json:"phone,omitempty"`
	Name            string                                 `json:"name,omitempty"`
	IdentityCardNum string                                 `json:"identityCardNum,omitempty"`
	Email           string                                 `json:"email,omitempty"`
	Birthday        int64                                  `json:"birthday,omitempty"`
	Gender          int64                                  `json:"gender,omitempty"`
	Education       string                                 `json:"education,omitempty"`
	Industry        string                                 `json:"industry,omitempty"`
	Hobby           string                                 `json:"hobby,omitempty"`
	Income          string                                 `json:"income,omitempty"`
	Province        string                                 `json:"province,omitempty"`
	ProvinceCode    int64                                  `json:"provinceCode,omitempty"`
	City            string                                 `json:"city,omitempty"`
	CityCode        int64                                  `json:"cityCode,omitempty"`
	Area            string                                 `json:"area,omitempty"`
	AreaCode        int64                                  `json:"areaCode,omitempty"`
	Address         string                                 `json:"address,omitempty"`
}
type PaasWeimobCustomerUpdateRequestExtInfo struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type PaasWeimobCustomerUpdateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreatePaasWeimobCustomerUpdateResponse() *PaasWeimobCustomerUpdateResponse {
	return &PaasWeimobCustomerUpdateResponse{}
}

// OnPaasWeimobCustomerUpdateServiceInvocation
// @id 1092
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1092?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新客户资料信息)
func (s *Service) OnPaasWeimobCustomerUpdateServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCustomerUpdateRequest, PaasWeimobCustomerUpdateResponse],
) (service *Service) {
	var invocation = &PaasWeimobCustomerUpdateService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCustomerUpdateRequest, PaasWeimobCustomerUpdateResponse](invocation),
		"PaasWeimobCustomerUpdateService",
		"weimob.customer.update",
	)
	return s
}
