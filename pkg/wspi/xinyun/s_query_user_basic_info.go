package xinyun

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasQueryUserBasicInfoService struct {
	handler spi.XinyunInvocationHandler[PaasQueryUserBasicInfoRequest, PaasQueryUserBasicInfoResponse]
}

func (s PaasQueryUserBasicInfoService) NewRequest() spi.InvocationRequest[PaasQueryUserBasicInfoRequest] {
	return &spi.XinyunInvocationRequest[PaasQueryUserBasicInfoRequest]{
		Params: &PaasQueryUserBasicInfoRequest{},
	}
}

func (s PaasQueryUserBasicInfoService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasQueryUserBasicInfoRequest],
) (
	spi.InvocationResponse[PaasQueryUserBasicInfoResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.XinyunInvocationRequest[PaasQueryUserBasicInfoRequest]))
}

type PaasQueryUserBasicInfoRequest struct {
	SourceObjectList []PaasQueryUserBasicInfoRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Pid              int64                                           `json:"pid,omitempty"`
	StoreId          int64                                           `json:"storeId,omitempty"`
	Wid              int64                                           `json:"wid,omitempty"`
}
type PaasQueryUserBasicInfoRequestSourceObjectList struct {
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
	Source       int64  `json:"source,omitempty"`
}

type PaasQueryUserBasicInfoResponse struct {
	ChannelStatuses []PaasQueryUserBasicInfoResponseChannelStatuses `json:"channelStatuses,omitempty"`
	IsMember        bool                                            `json:"isMember,omitempty"`
	Status          int64                                           `json:"status,omitempty"`
	StartTime       int64                                           `json:"startTime,omitempty"`
	ExpireTime      int64                                           `json:"expireTime,omitempty"`
	AppChannel      int64                                           `json:"appChannel,omitempty"`
	GetChannel      int64                                           `json:"getChannel,omitempty"`
	GetChannelId    int64                                           `json:"getChannelId,omitempty"`
	OuterStr        string                                          `json:"outerStr,omitempty"`
	CardPublishTime int64                                           `json:"cardPublishTime,omitempty"`
	GmtCreate       int64                                           `json:"gmtCreate,omitempty"`
	Name            string                                          `json:"name,omitempty"`
	Birthday        int64                                           `json:"birthday,omitempty"`
	Province        string                                          `json:"province,omitempty"`
	City            string                                          `json:"city,omitempty"`
	Area            string                                          `json:"area,omitempty"`
	ProvinceCode    string                                          `json:"provinceCode,omitempty"`
	CityCode        string                                          `json:"cityCode,omitempty"`
	AreaCode        string                                          `json:"areaCode,omitempty"`
	IdCard          string                                          `json:"idCard,omitempty"`
	Address         string                                          `json:"address,omitempty"`
	Mail            string                                          `json:"mail,omitempty"`
	Sex             int64                                           `json:"sex,omitempty"`
	Education       string                                          `json:"education,omitempty"`
	Hobby           string                                          `json:"hobby,omitempty"`
	Income          string                                          `json:"income,omitempty"`
	Industry        string                                          `json:"industry,omitempty"`
	Remark          string                                          `json:"remark,omitempty"`
	ExtInfo         []int64                                         `json:"extInfo,omitempty"`
}
type PaasQueryUserBasicInfoResponseChannelStatuses struct {
	Channel   int64  `json:"channel,omitempty"`
	Status    string `json:"status,omitempty"`
	ThirdCode string `json:"thirdCode,omitempty"`
}

func CreatePaasQueryUserBasicInfoResponse() *PaasQueryUserBasicInfoResponse {
	return &PaasQueryUserBasicInfoResponse{}
}

// OnPaasQueryUserBasicInfoServiceInvocation
// @id 1348
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=查询客户基本信息)
func (s *Service) OnPaasQueryUserBasicInfoServiceInvocation(
	handler spi.XinyunInvocationHandler[PaasQueryUserBasicInfoRequest, PaasQueryUserBasicInfoResponse],
) (service *Service) {
	var invocation = &PaasQueryUserBasicInfoService{handler}
	s.Register(
		spi.WrapInvoker[PaasQueryUserBasicInfoRequest, PaasQueryUserBasicInfoResponse](invocation),
		"PaasQueryUserBasicInfoService",
		"sQueryUserBasicInfo",
	)
	return s
}
