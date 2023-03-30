package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CityFreightTemplateQueryCityTemplateDetail
// @id 2440
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询同城运费模板详情)
func (client *Client) CityFreightTemplateQueryCityTemplateDetail(request *CityFreightTemplateQueryCityTemplateDetailRequest) (response *CityFreightTemplateQueryCityTemplateDetailResponse, err error) {
	rpcResponse := CreateCityFreightTemplateQueryCityTemplateDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CityFreightTemplateQueryCityTemplateDetailRequest struct {
	*api.BaseRequest

	Pid        int64 `json:"pid,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
	TemplateId int64 `json:"templateId,omitempty"`
}

type CityFreightTemplateQueryCityTemplateDetailResponse struct {
	Id                      int64    `json:"id,omitempty"`
	TemplateName            string   `json:"templateName,omitempty"`
	SupportCustomDatetime   bool     `json:"supportCustomDatetime,omitempty"`
	RequiredCustomDatetime  bool     `json:"requiredCustomDatetime,omitempty"`
	DeliveryDatetimeType    int      `json:"deliveryDatetimeType,omitempty"`
	LeastAheadDay           int      `json:"leastAheadDay,omitempty"`
	LeastAheadHour          int      `json:"leastAheadHour,omitempty"`
	LeastAheadMinute        int      `json:"leastAheadMinute,omitempty"`
	OptionTimeList          []string `json:"optionTimeList,omitempty"`
	MaxBehindDay            int      `json:"maxBehindDay,omitempty"`
	IsRegionPrivate         bool     `json:"isRegionPrivate,omitempty"`
	UpdateTime              int      `json:"updateTime,omitempty"`
	CreateTime              int      `json:"createTime,omitempty"`
	IsDefault               bool     `json:"isDefault,omitempty"`
	CityTemplateItemDtoList string   `json:"cityTemplateItemDtoList,omitempty"`
	Pid                     string   `json:"pid,omitempty"`
	StoreId                 string   `json:"storeId,omitempty"`
}

func CreateCityFreightTemplateQueryCityTemplateDetailRequest() (request *CityFreightTemplateQueryCityTemplateDetailRequest) {
	request = &CityFreightTemplateQueryCityTemplateDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "cityFreightTemplate/queryCityTemplateDetail", "api")
	request.Method = api.POST
	return
}

func CreateCityFreightTemplateQueryCityTemplateDetailResponse() (response *api.BaseResponse[CityFreightTemplateQueryCityTemplateDetailResponse]) {
	response = api.CreateResponse[CityFreightTemplateQueryCityTemplateDetailResponse](&CityFreightTemplateQueryCityTemplateDetailResponse{})
	return
}
