package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CityFreightTemplateAddCityTemplate
// @id 2439
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增同城模板)
func (client *Client) CityFreightTemplateAddCityTemplate(request *CityFreightTemplateAddCityTemplateRequest) (response *CityFreightTemplateAddCityTemplateResponse, err error) {
	rpcResponse := CreateCityFreightTemplateAddCityTemplateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CityFreightTemplateAddCityTemplateRequest struct {
	*api.BaseRequest

	CityTemplateItemDtoList []CityFreightTemplateAddCityTemplateRequestCityTemplateItemDtoList `json:"cityTemplateItemDtoList,omitempty"`
	TemplateName            string                                                             `json:"templateName,omitempty"`
	SupportCustomDatetime   bool                                                               `json:"supportCustomDatetime,omitempty"`
	RequiredCustomDatetime  bool                                                               `json:"requiredCustomDatetime,omitempty"`
	DeliveryDatetimeType    int                                                                `json:"deliveryDatetimeType,omitempty"`
	LeastAheadDay           int                                                                `json:"leastAheadDay,omitempty"`
	LeastAheadHour          int                                                                `json:"leastAheadHour,omitempty"`
	LeastAheadMinute        int                                                                `json:"leastAheadMinute,omitempty"`
	OptionTimeList          []string                                                           `json:"optionTimeList,omitempty"`
	MaxBehindDay            int                                                                `json:"maxBehindDay,omitempty"`
	IsRegionPrivate         bool                                                               `json:"isRegionPrivate,omitempty"`
	IsDefault               bool                                                               `json:"isDefault,omitempty"`
	Pid                     int64                                                              `json:"pid,omitempty"`
	StoreId                 int64                                                              `json:"storeId,omitempty"`
}

type CityFreightTemplateAddCityTemplateResponse struct {
	TemplateId int64 `json:"templateId,omitempty"`
	Result     bool  `json:"result,omitempty"`
}

func CreateCityFreightTemplateAddCityTemplateRequest() (request *CityFreightTemplateAddCityTemplateRequest) {
	request = &CityFreightTemplateAddCityTemplateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "cityFreightTemplate/addCityTemplate", "api")
	request.Method = api.POST
	return
}

func CreateCityFreightTemplateAddCityTemplateResponse() (response *api.BaseResponse[CityFreightTemplateAddCityTemplateResponse]) {
	response = api.CreateResponse[CityFreightTemplateAddCityTemplateResponse](&CityFreightTemplateAddCityTemplateResponse{})
	return
}

type CityFreightTemplateAddCityTemplateRequestCityTemplateItemDtoList struct {
	DeliveryScopeDetail CityFreightTemplateAddCityTemplateRequestDeliveryScopeDetail `json:"deliveryScopeDetail,omitempty"`
	BaseAmount          float64                                                      `json:"baseAmount,omitempty"`
	BaseFreight         float64                                                      `json:"baseFreight,omitempty"`
	TieredPrice         bool                                                         `json:"tieredPrice,omitempty"`
	UpperRadius         float64                                                      `json:"upperRadius,omitempty"`
	UpperWeight         float64                                                      `json:"upperWeight,omitempty"`
	ExtendRadius        float64                                                      `json:"extendRadius,omitempty"`
	ExtendWeight        float64                                                      `json:"extendWeight,omitempty"`
	ExtendRadiusFreight float64                                                      `json:"extendRadiusFreight,omitempty"`
	ExtendWeightFreight float64                                                      `json:"extendWeightFreight,omitempty"`
	DeliveryScopeType   int                                                          `json:"deliveryScopeType,omitempty"`
	RegionName          string                                                       `json:"regionName,omitempty"`
	IsDefault           bool                                                         `json:"isDefault,omitempty"`
}

type CityFreightTemplateAddCityTemplateRequestDeliveryScopeDetail struct {
	AreaList        CityFreightTemplateAddCityTemplateRequestAreaList `json:"areaList,omitempty"`
	CoordinatesList []map[string]any                                  `json:"coordinatesList,omitempty"`
	ServiceRadius   int                                               `json:"serviceRadius,omitempty"`
}

type CityFreightTemplateAddCityTemplateRequestAreaList struct {
}
