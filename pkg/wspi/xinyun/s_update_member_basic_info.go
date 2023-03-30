package xinyun

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasUpdateMemberBasicInfoService struct {
	handler spi.XinyunInvocationHandler[PaasUpdateMemberBasicInfoRequest, PaasUpdateMemberBasicInfoResponse]
}

func (s PaasUpdateMemberBasicInfoService) NewRequest() spi.InvocationRequest[PaasUpdateMemberBasicInfoRequest] {
	return &spi.XinyunInvocationRequest[PaasUpdateMemberBasicInfoRequest]{
		Params: &PaasUpdateMemberBasicInfoRequest{},
	}
}

func (s PaasUpdateMemberBasicInfoService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasUpdateMemberBasicInfoRequest],
) (
	spi.InvocationResponse[PaasUpdateMemberBasicInfoResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.XinyunInvocationRequest[PaasUpdateMemberBasicInfoRequest]))
}

type PaasUpdateMemberBasicInfoRequest struct {
	ExtInfo          PaasUpdateMemberBasicInfoRequestExtInfo            `json:"extInfo,omitempty"`
	SourceObjectList []PaasUpdateMemberBasicInfoRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Pid              int64                                              `json:"pid,omitempty"`
	StoreId          int64                                              `json:"storeId,omitempty"`
	Wid              int64                                              `json:"wid,omitempty"`
	Name             string                                             `json:"name,omitempty"`
	Birthday         string                                             `json:"birthday,omitempty"`
	Education        string                                             `json:"education,omitempty"`
	Income           string                                             `json:"income,omitempty"`
	Address          string                                             `json:"address,omitempty"`
	Mail             string                                             `json:"mail,omitempty"`
	Sex              int64                                              `json:"sex,omitempty"`
	Hobby            string                                             `json:"hobby,omitempty"`
	Industry         string                                             `json:"industry,omitempty"`
	Logo             string                                             `json:"logo,omitempty"`
	Remark           string                                             `json:"remark,omitempty"`
	Province         string                                             `json:"province,omitempty"`
	City             string                                             `json:"city,omitempty"`
	Area             string                                             `json:"area,omitempty"`
	ProvinceCode     string                                             `json:"provinceCode,omitempty"`
	CityCode         string                                             `json:"cityCode,omitempty"`
	AreaCode         string                                             `json:"areaCode,omitempty"`
}
type PaasUpdateMemberBasicInfoRequestExtInfo struct {
}
type PaasUpdateMemberBasicInfoRequestSourceObjectList struct {
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
	Source       string `json:"source,omitempty"`
}

type PaasUpdateMemberBasicInfoResponse struct {
	Status bool `json:"status,omitempty"`
}

func CreatePaasUpdateMemberBasicInfoResponse() *PaasUpdateMemberBasicInfoResponse {
	return &PaasUpdateMemberBasicInfoResponse{}
}

// OnPaasUpdateMemberBasicInfoServiceInvocation
// @id 1564
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=修改客户基本信息)
func (s *Service) OnPaasUpdateMemberBasicInfoServiceInvocation(
	handler spi.XinyunInvocationHandler[PaasUpdateMemberBasicInfoRequest, PaasUpdateMemberBasicInfoResponse],
) (service *Service) {
	var invocation = &PaasUpdateMemberBasicInfoService{handler}
	s.Register(
		spi.WrapInvoker[PaasUpdateMemberBasicInfoRequest, PaasUpdateMemberBasicInfoResponse](invocation),
		"PaasUpdateMemberBasicInfoService",
		"sUpdateMemberBasicInfo",
	)
	return s
}
