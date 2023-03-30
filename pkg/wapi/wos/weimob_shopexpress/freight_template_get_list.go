package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FreightTemplateGetList
// @id 1970
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1970?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=运费模板列表)
func (client *Client) FreightTemplateGetList(request *FreightTemplateGetListRequest) (response *FreightTemplateGetListResponse, err error) {
	rpcResponse := CreateFreightTemplateGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FreightTemplateGetListRequest struct {
	*api.BaseRequest
}

type FreightTemplateGetListResponse struct {
	TemplateOutList []FreightTemplateGetListResponseTemplateOutList `json:"templateOutList,omitempty"`
}

func CreateFreightTemplateGetListRequest() (request *FreightTemplateGetListRequest) {
	request = &FreightTemplateGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "freight/template/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateFreightTemplateGetListResponse() (response *api.BaseResponse[FreightTemplateGetListResponse]) {
	response = api.CreateResponse[FreightTemplateGetListResponse](&FreightTemplateGetListResponse{})
	return
}

type FreightTemplateGetListResponseTemplateOutList struct {
	DeliveryAreaInfo  FreightTemplateGetListResponseDeliveryAreaInfo `json:"deliveryAreaInfo,omitempty"`
	RuleList          []FreightTemplateGetListResponseRuleList       `json:"ruleList,omitempty"`
	FreightTemplateNo string                                         `json:"freightTemplateNo,omitempty"`
	ProgrammeName     string                                         `json:"programmeName,omitempty"`
}

type FreightTemplateGetListResponseDeliveryAreaInfo struct {
	DeliveryArea []FreightTemplateGetListResponseDeliveryArea `json:"deliveryArea,omitempty"`
	RegionName   string                                       `json:"regionName,omitempty"`
}

type FreightTemplateGetListResponseDeliveryArea struct {
	AreaCode string `json:"areaCode,omitempty"`
	AreaName string `json:"areaName,omitempty"`
}

type FreightTemplateGetListResponseRuleList struct {
	UnifiedPrice          FreightTemplateGetListResponseUnifiedPrice       `json:"unifiedPrice,omitempty"`
	PriceComputeRule      []FreightTemplateGetListResponsePriceComputeRule `json:"priceComputeRule,omitempty"`
	FreightTemplateRuleNo string                                           `json:"freightTemplateRuleNo,omitempty"`
	FreightTemplateNo     string                                           `json:"freightTemplateNo,omitempty"`
	FreightSchemeName     string                                           `json:"freightSchemeName,omitempty"`
	PriceComputeMode      int64                                            `json:"priceComputeMode,omitempty"`
	WeightUnit            int64                                            `json:"weightUnit,omitempty"`
	MinPrice              string                                           `json:"minPrice,omitempty"`
}

type FreightTemplateGetListResponseUnifiedPrice struct {
}

type FreightTemplateGetListResponsePriceComputeRule struct {
	High    FreightTemplateGetListResponseHigh `json:"high,omitempty"`
	Low     string                             `json:"low,omitempty"`
	Freight string                             `json:"freight,omitempty"`
}

type FreightTemplateGetListResponseHigh struct {
}
