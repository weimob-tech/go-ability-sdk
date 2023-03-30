package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAddFreightTemplate
// @id 1413
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增运费模板)
func (client *Client) GoodsAddFreightTemplate(request *GoodsAddFreightTemplateRequest) (response *GoodsAddFreightTemplateResponse, err error) {
	rpcResponse := CreateGoodsAddFreightTemplateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAddFreightTemplateRequest struct {
	*api.BaseRequest

	TemplateRuleVoList []GoodsAddFreightTemplateRequestTemplateRuleVoList `json:"templateRuleVoList,omitempty"`
	StoreId            int64                                              `json:"storeId,omitempty"`
	CalculateType      int                                                `json:"calculateType,omitempty"`
	IsFree             int                                                `json:"isFree,omitempty"`
	TemplateName       string                                             `json:"templateName,omitempty"`
	IsDefault          int64                                              `json:"isDefault,omitempty"`
}

type GoodsAddFreightTemplateResponse struct {
}

func CreateGoodsAddFreightTemplateRequest() (request *GoodsAddFreightTemplateRequest) {
	request = &GoodsAddFreightTemplateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/addFreightTemplate", "api")
	request.Method = api.POST
	return
}

func CreateGoodsAddFreightTemplateResponse() (response *api.BaseResponse[GoodsAddFreightTemplateResponse]) {
	response = api.CreateResponse[GoodsAddFreightTemplateResponse](&GoodsAddFreightTemplateResponse{})
	return
}

type GoodsAddFreightTemplateRequestTemplateRuleVoList struct {
	TemplateRuleItemVoList []GoodsAddFreightTemplateRequestTemplateRuleItemVoList `json:"templateRuleItemVoList,omitempty"`
}

type GoodsAddFreightTemplateRequestTemplateRuleItemVoList struct {
	SupportAddressLevelInfo GoodsAddFreightTemplateRequestSupportAddressLevelInfo `json:"supportAddressLevelInfo,omitempty"`
	BaseAmount              float64                                               `json:"baseAmount,omitempty"`
	ExtendAmount            float64                                               `json:"extendAmount,omitempty"`
	UpperNum                int64                                                 `json:"upperNum,omitempty"`
	ExtendNumUnit           int64                                                 `json:"extendNumUnit,omitempty"`
	UpperWeight             float64                                               `json:"upperWeight,omitempty"`
	ExtendWeightUnit        float64                                               `json:"extendWeightUnit,omitempty"`
	UpperVolume             float64                                               `json:"upperVolume,omitempty"`
	ExtendVolumeUnit        float64                                               `json:"extendVolumeUnit,omitempty"`
	IsDefault               int                                                   `json:"isDefault,omitempty"`
	SupportAreaName         string                                                `json:"supportAreaName,omitempty"`
}

type GoodsAddFreightTemplateRequestSupportAddressLevelInfo struct {
	List GoodsAddFreightTemplateRequestlist `json:"list,omitempty"`
}

type GoodsAddFreightTemplateRequestlist struct {
}
