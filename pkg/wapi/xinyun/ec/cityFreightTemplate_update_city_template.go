package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CityFreightTemplateUpdateCityTemplate
// @id 2438
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改同城模板)
func (client *Client) CityFreightTemplateUpdateCityTemplate(request *CityFreightTemplateUpdateCityTemplateRequest) (response *CityFreightTemplateUpdateCityTemplateResponse, err error) {
	rpcResponse := CreateCityFreightTemplateUpdateCityTemplateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CityFreightTemplateUpdateCityTemplateRequest struct {
	*api.BaseRequest

	Id                      int64  `json:"id,omitempty"`
	TemplateName            string `json:"templateName,omitempty"`
	SupportCustomDatetime   bool   `json:"supportCustomDatetime,omitempty"`
	RequiredCustomDatetime  bool   `json:"requiredCustomDatetime,omitempty"`
	DeliveryDatetimeType    int64  `json:"deliveryDatetimeType,omitempty"`
	LeastAheadDay           int64  `json:"leastAheadDay,omitempty"`
	LeastAheadHour          int64  `json:"leastAheadHour,omitempty"`
	LeastAheadMinute        int64  `json:"leastAheadMinute,omitempty"`
	OptionTimeList          string `json:"optionTimeList,omitempty"`
	MaxBehindDay            int64  `json:"maxBehindDay,omitempty"`
	IsRegionPrivate         bool   `json:"isRegionPrivate,omitempty"`
	IsDefault               bool   `json:"isDefault,omitempty"`
	CityTemplateItemDtoList string `json:"cityTemplateItemDtoList,omitempty"`
	Pid                     string `json:"pid,omitempty"`
	StoreId                 string `json:"storeId,omitempty"`
}

type CityFreightTemplateUpdateCityTemplateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCityFreightTemplateUpdateCityTemplateRequest() (request *CityFreightTemplateUpdateCityTemplateRequest) {
	request = &CityFreightTemplateUpdateCityTemplateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "cityFreightTemplate/updateCityTemplate", "api")
	request.Method = api.POST
	return
}

func CreateCityFreightTemplateUpdateCityTemplateResponse() (response *api.BaseResponse[CityFreightTemplateUpdateCityTemplateResponse]) {
	response = api.CreateResponse[CityFreightTemplateUpdateCityTemplateResponse](&CityFreightTemplateUpdateCityTemplateResponse{})
	return
}
