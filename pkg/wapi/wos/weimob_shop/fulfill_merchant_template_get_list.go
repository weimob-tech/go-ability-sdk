package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillMerchantTemplateGetList
// @id 1997
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1997?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询商家配送运费模板)
func (client *Client) FulfillMerchantTemplateGetList(request *FulfillMerchantTemplateGetListRequest) (response *FulfillMerchantTemplateGetListResponse, err error) {
	rpcResponse := CreateFulfillMerchantTemplateGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillMerchantTemplateGetListRequest struct {
	*api.BaseRequest

	BasicInfo FulfillMerchantTemplateGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsId   int64                                          `json:"goodsId,omitempty"`
}

type FulfillMerchantTemplateGetListResponse struct {
	DefaultFreightTemplate  FulfillMerchantTemplateGetListResponseDefaultFreightTemplate  `json:"defaultFreightTemplate,omitempty"`
	FreightTemplateList     []FulfillMerchantTemplateGetListResponseFreightTemplateList   `json:"freightTemplateList,omitempty"`
	SelectedFreightTemplate FulfillMerchantTemplateGetListResponseSelectedFreightTemplate `json:"selectedFreightTemplate,omitempty"`
}

func CreateFulfillMerchantTemplateGetListRequest() (request *FulfillMerchantTemplateGetListRequest) {
	request = &FulfillMerchantTemplateGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/merchant/template/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillMerchantTemplateGetListResponse() (response *api.BaseResponse[FulfillMerchantTemplateGetListResponse]) {
	response = api.CreateResponse[FulfillMerchantTemplateGetListResponse](&FulfillMerchantTemplateGetListResponse{})
	return
}

type FulfillMerchantTemplateGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type FulfillMerchantTemplateGetListResponseDefaultFreightTemplate struct {
	CalculateType int64  `json:"calculateType,omitempty"`
	IsFree        int64  `json:"isFree,omitempty"`
	TemplateId    int64  `json:"templateId,omitempty"`
	TemplateName  string `json:"templateName,omitempty"`
}

type FulfillMerchantTemplateGetListResponseFreightTemplateList struct {
	CalculateType int64  `json:"calculateType,omitempty"`
	IsFree        int64  `json:"isFree,omitempty"`
	TemplateId    int64  `json:"templateId,omitempty"`
	TemplateName  string `json:"templateName,omitempty"`
}

type FulfillMerchantTemplateGetListResponseSelectedFreightTemplate struct {
	TemplateRuleDtoList []FulfillMerchantTemplateGetListResponseTemplateRuleDtoList `json:"templateRuleDtoList,omitempty"`
	TemplateId          int64                                                       `json:"templateId,omitempty"`
	TemplateName        string                                                      `json:"templateName,omitempty"`
	CalculateType       int64                                                       `json:"calculateType,omitempty"`
	IsFree              int64                                                       `json:"isFree,omitempty"`
	IsDefault           int64                                                       `json:"isDefault,omitempty"`
}

type FulfillMerchantTemplateGetListResponseTemplateRuleDtoList struct {
	TemplateRuleItemDtoList []FulfillMerchantTemplateGetListResponseTemplateRuleItemDtoList `json:"templateRuleItemDtoList,omitempty"`
	RuleId                  int64                                                           `json:"ruleId,omitempty"`
	TemplateId              int64                                                           `json:"templateId,omitempty"`
	RuleType                int64                                                           `json:"ruleType,omitempty"`
	DeliveryMethodId        int64                                                           `json:"deliveryMethodId,omitempty"`
	DeliveryMethodName      string                                                          `json:"deliveryMethodName,omitempty"`
}

type FulfillMerchantTemplateGetListResponseTemplateRuleItemDtoList struct {
	DeliveryScopeDetail FulfillMerchantTemplateGetListResponseDeliveryScopeDetail `json:"deliveryScopeDetail,omitempty"`
	ItemId              int64                                                     `json:"itemId,omitempty"`
	TemplateId          int64                                                     `json:"templateId,omitempty"`
	RuleId              int64                                                     `json:"ruleId,omitempty"`
	BaseNum             int64                                                     `json:"baseNum,omitempty"`
	BasePrice           int64                                                     `json:"basePrice,omitempty"`
	ExtendNum           int64                                                     `json:"extendNum,omitempty"`
	ExtendPrice         int64                                                     `json:"extendPrice,omitempty"`
	ScopeName           string                                                    `json:"scopeName,omitempty"`
	IsDefault           int64                                                     `json:"isDefault,omitempty"`
}

type FulfillMerchantTemplateGetListResponseDeliveryScopeDetail struct {
}
